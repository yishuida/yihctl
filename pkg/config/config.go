package config

import (
	log "github.com/sirupsen/logrus"
	myLog "ydq.io/yihctl/pkg/log"
)

var (
	ConfLogger = &log.Entry{}
)

func init() {
	ConfLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "config",
	})
}
