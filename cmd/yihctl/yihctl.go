package main // import "github.com/yishuida/yihctl/cmd/yihctl"

import (
	log "github.com/sirupsen/logrus"
	"github.com/yishuida/yihctl/pkg/cli"
	"os"
)

var (
	settings  = cli.New()
	CmdLogger = &log.Entry{}
)

func init() {
	CmdLogger.WithFields(log.Fields{
		"pkg": "cmd",
	})
}

func debug() {
	CmdLogger.Level = log.DebugLevel
}

func main() {
	cmd := newRootCmd(os.Stdout, os.Args[1:])

	if err := cmd.Execute(); err != nil {
		debug()
		os.Exit(1)
	}
}
