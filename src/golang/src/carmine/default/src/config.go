package main

// Config is a container type for configuration information.
type Config struct {
	// Debug set to true turns on debugging for the application; especially
	// useful during develepment:
	//
	// 1. Templates are compiled per request instead of per process instance
	//    allowing reloading the templates as you make changes.
	Debug bool

	// The directory that contains the templates.
	TemplatesPath string
}
