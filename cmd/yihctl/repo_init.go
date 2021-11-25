package main

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"ydq.io/yihctl/pkg/config"
	gitydq "ydq.io/yihctl/pkg/git"
	"ydq.io/yihctl/pkg/util"
)

const repoInitDesc = `
Init configuration file in ./config.yaml, this default configuration fil will clone repository
github、gitlab、gitee repository.
`

func newRepoInitCmd(out io.Writer, r repoOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init [keyword]",
		Short: "initialization cfgFile and default repository",
		Long:  repoInitDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gitInitRun(out, r)
		},
	}

	return cmd
}

func gitInitRun(out io.Writer, r repoOptions) error {
	grc := readConfigFile(r.cfgFile)
	if grc == nil {
		return errors.New("config is empty")
	}
	for _, repo := range grc.Repositories {
		remote := grc.GetRemote(repo.From)

		for _, r := range repo.Repos {
			gitRepoUrl := config.GetGitRepoUrl(remote, r.Source)
			targetPath := fmt.Sprintf("%s/%s", repo.Path, strings.Split(r.Source, "/")[1])

			if isLocalRepositoryExist(targetPath) {
				cmdLogger.Warnf("git repository in %s existing", targetPath)
				break
			} else {
				util.Path(targetPath)
			}

			cmdLogger.Infof("git repository gitRepoUrl is %s, clone into %s", gitRepoUrl, targetPath)
			return gitydq.Clone(out, remote.Auth.GenerateAuth(), gitRepoUrl, targetPath)
		}
	}

	return nil
}

func isLocalRepositoryExist(path string) bool {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	if err != nil || r == nil {
		cmdLogger.Warningf("go-git open %s failed", path)
		return false
	}
	return true
}
