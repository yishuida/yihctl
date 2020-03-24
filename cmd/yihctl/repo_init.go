package main

import (
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yishuida/yihctl/pkg/config"
	"gopkg.in/src-d/go-git.v4"
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
	f.StringVar(&r.gitRepo, "git-repo", config.Path(".yih")+string(os.PathSeparator)+"git-repo.yaml", "list manager git repository")

	return cmd
}

func (r *repoInitOptions) run(out io.Writer, args []string) error {
	gr, err := os.OpenFile(r.gitRepo, os.O_RDWR|os.O_CREATE, 0644)
	if gr == nil || err != nil {
		CmdLogger.Warn(err)
	}
	defer gr.Close()

	fileInfo, err := gr.Stat()
	if gr == nil || err != nil {
		CmdLogger.Warn(err)
	}
	if n := fileInfo.Size(); n == 0 {
		CmdLogger.Info("init default git-repo.yaml")
		_, _ = gr.WriteString(config.DefaultGitRepo)
	}

	cloneRepo(loadConfig(r.gitRepo))

	return nil
}

func loadConfig(path string) *config.GitRepo {
	gitRepo := config.GitRepo{}
	if err := configor.New(&configor.Config{Debug: true}).Load(&gitRepo, path); err != nil {
		CmdLogger.Error(err)
	}

	return &gitRepo
}

// TODO move to utils package
func cloneRepo(gitRepos *config.GitRepo) {
	for _, remote := range gitRepos.Remotes {
		for _, org := range remote.Organizations {
			for _, repo := range org.Repos {
				url := remote.Domain + string(os.PathSeparator) + org.Name + string(os.PathSeparator) + repo.Name + ".git"
				destPath := config.HomePath() + string(os.PathSeparator) + org.Dir + string(os.PathSeparator) + repo.Name

				if _, err := os.Stat(destPath); os.IsNotExist(err) {
					config.Path(destPath)
					CmdLogger.WithFields(log.Fields{
						"repo": url,
						"path": destPath,
					}).Info("cloning repository")
					_, _ = git.PlainClone(destPath, false, &git.CloneOptions{
						URL:      url,
						Progress: os.Stdout,
					})
				} else {
					CmdLogger.WithFields(log.Fields{
						"repo": url,
						"path": destPath,
					}).Info("repository is existing!")
				}
			}
		}
	}
}
