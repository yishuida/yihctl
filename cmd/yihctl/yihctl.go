package main // import "ydq.io/yihctl/cmd/yihctl"

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	myLog "ydq.io/yihctl/pkg/log"
)

var (
	cmdLogger = &log.Entry{}
)

func init() {
	cmdLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "cmd",
	})
}

func main() {
	cmd := newRootCmd(os.Stdout, os.Args[1:])

	cobra.OnInitialize(func() {
		//actionConfig.Init()
	})
	if err := cmd.Execute(); err != nil {
		cmdLogger.Error(err)
		os.Exit(1)
	}
}
