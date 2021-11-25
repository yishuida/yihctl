package util

import (
	log "github.com/sirupsen/logrus"
	myLog "ydq.io/yihctl/pkg/log"
)

var (
	utilLogger = &log.Entry{}
)

func init() {
	utilLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "config",
	})
}
