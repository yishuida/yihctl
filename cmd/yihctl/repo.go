package main // import "github.com/yishuida/yihctl/cmd/yihctl"

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/cmd/helm/require"
	"io"
	"io/ioutil"
	"ydq.io/yihctl/pkg/config"
)

const repoDesc = `
maintenance used git repository, so we can init,sync,backup then by the easy way.
default repository configuration file is ./config.yaml. you can edit this file when you
want to change repositories.
`

type repoOptions struct {
	cfgFile string
	// TODO add auth
}

func newRepoCmd(out io.Writer) *cobra.Command {
	r := repoOptions{}
	cmd := &cobra.Command{
		Use:   "repo init|sync [ARGS]",
		Short: "manager git repository in home dir",
		Long:  repoDesc,
		Args:  require.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	f := cmd.Flags()
	f.StringVarP(&r.cfgFile, "config", "c", "config.yaml", "manager git repository")

	cmd.AddCommand(newRepoInitCmd(out, r))
	cmd.AddCommand(newRepoSyncCmd(out))

	return cmd
}

func readConfigFile(path string) *config.GitRepoConfig {
	fileBuf, err := ioutil.ReadFile(path)
	if err != nil {
		cmdLogger.Error(err)
		return nil
	}
	var grc config.GitRepoConfig
	err = yaml.Unmarshal(fileBuf, &grc)
	if err != nil {
		cmdLogger.Error(err)
		return nil
	}
	return &grc
}
