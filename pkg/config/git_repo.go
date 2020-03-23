package config

import "gopkg.in/src-d/go-git.v4/plumbing/transport"

type GitRepo struct {
	Name    string
	Version string
	Remotes []remote
}

type remote struct {
	Domain        string
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
