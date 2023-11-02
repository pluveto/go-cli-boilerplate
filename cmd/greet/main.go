package main

import (
	"os"
	"path/filepath"

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

	app.InitLogger(conf.Log, args.Verbose)
	logger.Debug("log level: ", logger.GetLevel())

	app.Run(args, conf)
}

func loadArgsValid() app.Args {
	var args app.Args
	arg.MustParse(&args)
	return args
}

func getAppDir() string {
	dir, err := os.Executable()
	if err != nil {
		logger.Fatal(err)
	}
	return filepath.Dir(dir)
}

// loadConfValid loads the configuration from the given path.
// If the path is empty, the defaultConfPath is used.
// If the path is relative, the app executable dir is prepended.
//
// Example
func loadConfValid(configFileName string, defaultConf app.Conf, defaultConfPath string) app.Conf {
	if configFileName == "" {
		configFileName = defaultConfPath
	}
	path := findSuitablePath(configFileName)

	_, err := toml.DecodeFile(path, &defaultConf)
	if err != nil {
		logger.Info("failed to load config file: ", err, " using default config")
	}
	logger.WithField("conf", &defaultConf).Debug("configuration loaded")
	return defaultConf
}

func findSuitablePath(configFileName string) string {
	path := ""
	// app executable dir + config.toml has the highest priority
	preferredPath := filepath.Join(getAppDir(), configFileName)
	if _, err := os.Stat(preferredPath); err == nil {
		path = preferredPath
	}
	if len(path) == 0 {
		preferredPathUserConfig := filepath.Join(os.Getenv("HOME"), ".config", app.Name, configFileName)
		if _, err := os.Stat(preferredPathUserConfig); err == nil {
			path = preferredPathUserConfig
		}
	}
	return path
}
