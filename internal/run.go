package internal

import (
	"os"

	"github.com/no-src/log"
	"github.com/no-src/rest-client/internal/flag"
)

func Run() {
	flag.InitFlags()

	defer log.Close()
	if len(flag.LogConf) > 0 {
		logger, err := log.CreateLoggerFromConfig(flag.LogConf)
		if err != nil {
			log.Error(err, "init logger with config failed, fallback to use the default logger")
		} else {
			log.InitDefaultLogger(logger)
		}
	}

	var h handler
	if flag.IsRun {
		h = sendHandler(flag.RequestId)
	} else {
		h = showHandler(flag.RequestId)
	}

	if err := run(flag.ConfigFile, flag.HttpFile, h); err != nil {
		log.Error(err, "run rest client error")
	}
}

func run(configFile, httpFile string, h handler) error {
	if err := initConfig(configFile); err != nil {
		return err
	}
	data, err := os.ReadFile(httpFile)
	if err != nil {
		return err
	}
	requests, err := parseHttp(string(data))
	if err != nil {
		return err
	}

	for index, request := range requests {
		err = h(index+1, request)
		if err != nil {
			return err
		}
	}
	return nil
}
