// Configuration types and definitions.

package main

// Config is a container type for configuration information.
type Config struct {
	// Debug set to true turns on debugging for the application; especially
	// useful during develepment:
	//
	// 1. Templates are compiled per request instead of per process instance
	//    allowing reloading the templates as you make changes.
	Debug bool

	// The host interface to which the HTTP server will bind.
	Host string

	// The host port on which the HTTP server will listen.
	Port int

	// The directory that contains the templates.
	TemplatesDir string
}
