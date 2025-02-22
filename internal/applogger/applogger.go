package applogger

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

func Create(logLevel string) *zap.Logger {
	rawJSON := []byte(fmt.Sprintf(
		`{
	   "level": "%s",
	   "encoding": "console",
	   "outputPaths": ["stdout"],
	   "errorOutputPaths": ["stderr"],
	   "encoderConfig": {
	     "messageKey": "message",
	     "levelKey": "level",
	     "levelEncoder": "lowercase"
	   }
	  }`,
		logLevel,
	))

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
