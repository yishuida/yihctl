package config

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	myLog "github.com/yishuida/yihctl/pkg/log"
	"os"
)

var (
	ConfLogger = &log.Entry{}
	configPath = ""
)

func init() {

}

func init() {
	ConfLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "cmd",
	})
}

func HomePath() string {
	return Path("")
}
func Path(path string) string {
	configPath, err := homedir.Dir()
	if err != nil {
		ConfLogger.Error(err)
	}
	if path != "" {
		configPath += string(os.PathSeparator) + path
	}
	ensurePath(configPath)

	return configPath
}

func ensurePath(path string) {
	// when path is no exist, create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		ConfLogger.Warn(err)
		if err := os.MkdirAll(path, 644); err != nil {
			ConfLogger.Error(err)
		}
	}
}
