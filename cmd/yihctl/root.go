package main // import "ydq.io/yihctl/cmd/yihctl"
import (
	"github.com/spf13/cobra"
	"io"
)

const rootDesc = `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`

func newRootCmd(out io.Writer, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yihctl",
		Short: "A sample command tool line",
		Long:  rootDesc,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	flags := cmd.PersistentFlags()

	//settings.AddFlags(flags)

	// We can safely ignore any errors that flags.Parse encounters since
	// those errors will be caught later during the call to cmd.Execution.
	// This call is required to gather configuration information prior to
	// execution.
	flags.ParseErrorsWhitelist.UnknownFlags = true
	_ = flags.Parse(args)

	cmd.AddCommand(
		newRepoCmd(out),
		newVersionCmd(out),
	)

	return cmd
}
