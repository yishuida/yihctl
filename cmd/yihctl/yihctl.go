package main // import "github.com/yishuida/yihctl/cmd/yihctl"

import (
	"fmt"
	"github.com/yishuida/yihctl/pkg/cli"
	"log"
	"os"
)

var (
	settings = cli.New()
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func debug(format string, v ...interface{}) {
	if settings.Debug {
		format := fmt.Sprintf("[debug] %s\n", format)
		_ = log.Output(2, fmt.Sprintf(format, v...))
	}
}

func main() {
	cmd := newRootCmd(os.Stdout, os.Args[1:])

	if err := cmd.Execute(); err != nil {
		debug("%+v", err)
		os.Exit(1)
	}
}
