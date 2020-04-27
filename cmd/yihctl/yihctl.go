package main // import "github.com/yishuida/yihctl/cmd/yihctl"

import (
	log "github.com/sirupsen/logrus"
	"github.com/yishuida/yihctl/pkg/cli"
	myLog "github.com/yishuida/yihctl/pkg/log"
	"os"
)

var (
	envSetting = cli.New()
	cmdLogger  = &log.Entry{}
)

func init() {
	cmdLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "cmd",
	})
}

func debug() {
	cmdLogger.Level = log.DebugLevel
}

func main() {
	cmdLogger.Info("Starting...")
	cmd := newRootCmd(os.Stdout, os.Args[1:])

	if err := cmd.Execute(); err != nil {
		debug()
		os.Exit(1)
	}
}
