// TODO(yesudeep): Move this to a middleware package.
package main

import "net/http"

// An HTTP error handler interface that allows middleware-style error-handling.
type HTTPErrorHandler interface {
	// ServeHTTP does exactly what http.Handler.ServeHTTP does, except, it also
	// returns an error for our error handlers to handle it.
	ServeHTTP(w http.ResponseWriter, r *http.Request) error
}

// Internal type used by the NewHTTPError function.
type httpError struct {
	status  int
	message string
}

func (e httpError) Status() int   { return e.status }
func (e httpError) Error() string { return e.message }

// HTTPError encapsulates an HTTP error.
type HTTPError interface {
	Status() int
	error
}

// NewHTTPError generates a new HTTPError.
func NewHTTPError(status int, message string) HTTPError {
	return httpError{status, message}
}

// handleError internalizes functionality common to ErrorHandler and
// ErrorHandlerFunc.
func handleError(err error, w http.ResponseWriter, r *http.Request) {
	status := http.StatusInternalServerError // Default HTTP error status.
	if herr, ok := err.(HTTPError); ok {
		status = herr.Status()
	}
	http.Error(w, err.Error(), status)
}

// ErrorHandlerFunc allows converting HTTP handler functions to those that
// handle HTTP errors painlessly.
//
// In order to use this type, simply cast your HTTP handler function to this
// type.
type ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP implements the http.Handler interface for ErrorHandlerFunc.
//
// The default error status code is `http.StatusInternalServerError` (status
// code 500) with the error message included as the description.
func (h ErrorHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		handleError(err, w, r)
	}
}

// ErrorHandler allows composing handlers middleware-style to handle HTTP
// errors painlessly.
type ErrorHandler struct {
	Handler HTTPErrorHandler
}

// ServeHTTP implements the http.Handler interface for ErrorHandler.
//
// The default error status code is `http.StatusInternalServerError` (status
// code 500) with the error message included as the description.
func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h.Handler.ServeHTTP(w, r); err != nil {
		handleError(err, w, r)
	}
}
