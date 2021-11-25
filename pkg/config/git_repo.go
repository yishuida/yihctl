package config

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	ssh2 "golang.org/x/crypto/ssh"
	"os"
)

type GitRepoConfig struct {
	Name         string
	Version      string
	Remotes      map[string]Remote
	Repositories []Repository
}

type Remote struct {
	Name   string
	Domain string
	Scheme string
	Type   string
	Auth   Auth `yaml:"auth"`
}

type Auth struct {
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	PrivateKeyFile string `yaml:"privateKeyFile"`
}

type Repository struct {
	Name        string
	Description string
	From        string
	To          string
	Path        string
	Repos       []struct {
		Source string
		Target string
	}
}

func GetGitRepoUrl(remote Remote, repo string) string {

	return renderUrl(remote, repo)
}

func (g *GitRepoConfig) GetRemote(name string) Remote {
	return g.Remotes[name]
}

func (a *Auth) GenerateAuth() transport.AuthMethod {
	if a.Username != "" && a.Password != "" {
		return &http.BasicAuth{
			Username: a.Username,
			Password: a.Password,
		}
	}
	if a.PrivateKeyFile != "" {

		_, err := os.Stat(a.PrivateKeyFile)
		if err != nil {
			ConfLogger.Warningf("read file %s failed %s\n", a.PrivateKeyFile, err.Error())
			return nil
		}

		// Clone the given repository to the given directory
		publicKeys, err := ssh.NewPublicKeysFromFile("git", a.PrivateKeyFile, "")
		if err != nil {
			ConfLogger.Warningf("generate publickeys failed: %s\n", err.Error())
			return nil
		}
		publicKeys.HostKeyCallback = ssh2.InsecureIgnoreHostKey()

		return publicKeys
	}
	return nil
}

func renderUrl(remote Remote, repo string) string {
	tpl := getTpl(remote.Scheme)
	return fmt.Sprintf(tpl, remote.Domain, repo)
}

func getTpl(scheme string) string {
	switch scheme {
	case "http":
		return httpUrlTpl
	// case "https": return httpsUrlTpl
	case "ssh":
		return sshUrlTpl
	default:
		return httpsUrlTpl
	}
}
