package cmd // github.com/yishuida/yihctl/cmd
import (
	"github.com/spf13/cobra"
	"os"
)

var cmdLongDescribe = `A lnger description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`

var rootCmd = &cobra.Command{
	Use:   "yihctl",
	Short: "A sample command tool line",
	Long:  cmdLongDescribe,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
