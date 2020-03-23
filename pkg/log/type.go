package log

import log "github.com/sirupsen/logrus"

var StdLog = log.New()

func init() {
	formatter := new(log.TextFormatter)
	formatter.FullTimestamp = true
}
