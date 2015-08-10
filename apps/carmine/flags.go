import "flag"

func init() {
	config = Config{}
	flag.StringVar(&config.Host, "host", "", "the host on which to listen")
	flag.IntVar(&config.Port, "port", 8080, "the port on which to listen")
	flag.BoolVar(&config.Debug, "debug", false, "turns on debugging")
	flag.StringVar(
		&config.TemplatesDir,
		"templatesDir",
		"templates", "the directory that contains the templates")
	flag.Parse()
}
