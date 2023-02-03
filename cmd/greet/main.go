package main

import (
	"example.com/m/cmd/greet/app"
	"example.com/m/pkg/logger"
	"github.com/BurntSushi/toml"
	"github.com/alexflint/go-arg"
)

// main Entry point of the application
func main() {
	// main Entry point of the application
	var args app.Args
	var conf app.Conf
	var defaultConf = app.GetDefaultConf()

	args = loadArgsValid()
	conf = loadConfValid(args.Config, defaultConf, "config.toml")

	app.InitLogger(defaultConf.Log, args.Verbose)
	logger.Debug("log level: ", logger.GetLevel())

	app.Run(args, conf)
}

func loadArgsValid() app.Args {
	var args app.Args
	arg.MustParse(&args)
	return args
}

func loadConfValid(path string, defaultConf app.Conf, defaultConfPath string) app.Conf {
	if path == "" {
		path = defaultConfPath
	}
	_, err := toml.DecodeFile(path, &defaultConf)
	if err != nil {
		logger.Warn("failed to load config file: ", err, " using default config")
	}
	logger.WithField("conf", &defaultConf).Debug("configuration loaded")
	return defaultConf
}
