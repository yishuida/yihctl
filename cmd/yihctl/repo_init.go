package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yishuida/yihctl/pkg/config"
	"github.com/yishuida/yihctl/pkg/util"
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

// TODO move to util package
func cloneRepo(gitRepos *config.GitRepo) {
	for _, remote := range gitRepos.Remotes {
		remoteMap := remote.GetRemoteUrl("")
		for url, path := range *remoteMap {
			if util.EmptyDir(path) {
				cmdLogger.WithFields(log.Fields{
					"url":  url,
					"path": path,
				}).Info("cloning repository")

				var err error
				if remote.Scheme == "git" {
					publicKeys, err := ssh.NewPublicKeysFromFile("git", "/Users/vista/.ssh/id_rsa", "")
					if err != nil {
						cmdLogger.Warn(err)
					}
					_, err = git.PlainClone(path, false, &git.CloneOptions{
						Auth:     publicKeys,
						URL:      url,
						Progress: os.Stdout,
					})
					CheckIfError(err)

				} else {
					_, err = git.PlainClone(path, false, &git.CloneOptions{
						URL:      url,
						Progress: os.Stdout,
					})
				}

				CheckIfError(err)
			} else {
				cmdLogger.WithFields(log.Fields{"url": url, "path": path}).Info("Spik up git clone")
			}
		}
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	//os.Exit(1)
}
