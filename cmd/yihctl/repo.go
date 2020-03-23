package main // import "github.com/yishuida/yihctl/cmd/yihctl"

import (
	"github.com/spf13/cobra"
	"io"
)

var repoDesc = ``

type repoOptions struct {
	homeDir string
	sync    bool
}

func newRepoCmd(out io.Writer) *cobra.Command {
	r := &repoOptions{}

	cmd := &cobra.Command{
		Use:   "repo",
		Short: "",
		Long:  repoDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return r.run(out)
		},
	}
	cmd.Flags().StringVarP(&r.homeDir, "home", "d", "", "repo home path")
	cmd.Flags().BoolVarP(&r.sync, "sync", "s", false, "Sync default repo")
	return cmd
}

func (r *repoOptions) run(out io.Writer) error {
	CmdLogger.Printf("repo %s\n", r.homeDir)

	return nil
}
