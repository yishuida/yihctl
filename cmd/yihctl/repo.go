package main // import "github.com/yishuida/yihctl/cmd/yihctl"

import (
	"github.com/spf13/cobra"
	"github.com/yishuida/yihctl/pkg/action"
	"io"
	"os"
)

const repoDesc = `
maintenance used git repository, so we can init,sync,backup then by the easy way.
default repository config in $HOME/.yih/git-repo.yaml. you can edit this file when you
want to change repositories.
`

func newRepoCmd(cfg *action.Configuration, out io.Writer) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "repo [keyword]",
		Short: "manager git repository in home dir",
		Long:  repoDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(newRepoInitCmd(os.Stdout))
	return cmd
}
