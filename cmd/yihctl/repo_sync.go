package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"strings"
	"ydq.io/yihctl/pkg/config"
	gitydq "ydq.io/yihctl/pkg/git"
)

func newRepoSyncCmd(out io.Writer, r repoOptions) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "sync [keyword]",
		Short: "sync git repository",
		Long:  repoInitDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return gitSyncRun(out, r)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&r.cfgFile, "config", "v", "config.yaml", "manager git repository")

	return cmd
}

func gitSyncRun(out io.Writer, r repoOptions) error {
	grc := readConfigFile(r.cfgFile)
	if grc == nil {
		return errors.New("config is empty")
	}

	for _, repo := range grc.Repositories {
		from := grc.GetRemote(repo.From)
		to := grc.GetRemote(repo.To)

		for _, r := range repo.Repos {
			sourceGitRepoUrl := config.GetGitRepoUrl(from, r.Source)
			targetGitRepoUrl := config.GetGitRepoUrl(to, r.Target)
			targetPath := fmt.Sprintf("%s/%s", repo.Path, strings.Split(r.Source, "/")[1])

			if !isLocalRepositoryExist(targetPath) {
				cmdLogger.Warnf("git repository in %s is not existing, skip up.", targetPath)
				break
			}
			err := gitydq.AddRemote(targetPath, to.Name, targetGitRepoUrl)
			if err != nil {
				return err
			}

			cmdLogger.Infof("Starting pull git repo %s, path %s\n", sourceGitRepoUrl, targetPath)
			err = gitydq.Pull(out, from.Auth.GenerateAuth(), targetPath, from.Name)
			if err != nil {
				return err
			}

			cmdLogger.Infof("Starting push git repo :url %s. path %s\n", targetGitRepoUrl, targetPath)
			err = gitydq.Push(out, to.Auth.GenerateAuth(), targetPath, to.Name)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
