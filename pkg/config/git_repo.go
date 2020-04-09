package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yishuida/yihctl/pkg/util"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	"strings"
)

type GitRepo struct {
	Name    string
	Version string
	Remotes []Remote
}

type Remote struct {
	Domain        string
	Scheme        string
	Dir           string
	Organizations []Organization
	Auth          transport.AuthMethod
}

type Organization struct {
	Name  string
	Dir   string
	Repos []Repo
}

type Repo struct {
	Name   string
	Branch string `default:"origin"`
}

func (r *Remote) GetRemoteUrl(gitPath string) *map[string]string {
	var remoteMap = map[string]string{}

	gitHome := util.HomePath()
	if gitPath != "" {
		gitHome = gitPath
	}

	for _, org := range r.Organizations {
		if org.Repos != nil {
			for _, repo := range org.Repos {
				url := fmt.Sprintf(getTpl(r.Scheme), r.Scheme, r.Domain, org.Name, repo.Name)

				destPath := util.Path(gitHome, r.Dir, org.Dir, repo.Name)
				ConfLogger.WithFields(log.Fields{
					"repo": url,
					"path": destPath,
				}).Info("Add repository to clone list")

				remoteMap[url] = destPath
			}
		}
	}

	return &remoteMap
}

func getTpl(scheme string) string {
	if strings.ToLower(scheme) == "https" || strings.ToLower(scheme) == "http" {
		return httpUrlTpl
	} else if strings.ToLower(scheme) == "git" {
		return sshUrlTpl
	}
	return ""
}
