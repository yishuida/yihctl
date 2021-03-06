package config

import (
	log "github.com/sirupsen/logrus"
	myLog "github.com/yishuida/yihctl/pkg/log"
)

var (
	ConfLogger = &log.Entry{}
)

func init() {
	ConfLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "config",
	})
}
