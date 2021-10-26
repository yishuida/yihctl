package main

import (
	"github.com/jinzhu/configor"
	"github.com/spf13/cobra"
	"github.com/yishuida/yihctl/pkg/config"
	mygit "github.com/yishuida/yihctl/pkg/git"
	"io"
	"os"
)

const repoInitDesc = `
Init config file in $HOME/.yih/git-repo.yaml, this default config will clone repository in
helm, kubernetes, choerodon, yishuida organization,
`

type repoInitOptions struct {
	gitRepo string
	// TODO add auth
}

func newRepoInitCmd(out io.Writer) *cobra.Command {
	r := &repoInitOptions{}

	cmd := &cobra.Command{
		Use:   "init [keyword]",
		Short: "initialization config and default repository",
		Long:  repoInitDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return r.run(out, args)
		},
	}

	f := cmd.Flags()
	f.StringVar(&r.gitRepo, "git-repo", "", "list manager git repository")

	return cmd
}

func (r *repoInitOptions) run(out io.Writer, args []string) error {
	gr, err := os.OpenFile(r.gitRepo, os.O_RDWR|os.O_CREATE, 0644)
	if gr == nil || err != nil {
		cmdLogger.Warn(err)
	}
	defer gr.Close()

	fileInfo, err := gr.Stat()
	if gr == nil || err != nil {
		cmdLogger.Warn(err)
	}
	if n := fileInfo.Size(); n == 0 {
		cmdLogger.Warn("git-repo.yaml is empty!")
	}

	cloneRepo(loadConfig(r.gitRepo))

	return nil
}

func loadConfig(path string) *config.GitRepo {
	gitRepo := config.GitRepo{}
	if err := configor.New(&configor.Config{}).Load(&gitRepo, path); err != nil {
		cmdLogger.Error(err)
	}

	return &gitRepo
}

func cloneRepo(gitRepos *config.GitRepo) {
	for _, remote := range gitRepos.Remotes {
		remoteMap := remote.GetRemoteUrl("")
		for url, path := range *remoteMap {
			if err := mygit.CloneOrPull(url,path); err != nil {
				cmdLogger.Errorf("git or clone repo failed: %s", err)
			}
		}
	}
}
