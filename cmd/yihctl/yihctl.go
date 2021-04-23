package main // import "github.com/yishuida/yihctl/cmd/yihctl"

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yishuida/yihctl/pkg/action"
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

func main() {
	actionConfig := new(action.Configuration)

	cmd := newRootCmd(actionConfig, os.Stdout, os.Args[1:])

	cobra.OnInitialize(func() {
		actionConfig.Init()
	})
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
