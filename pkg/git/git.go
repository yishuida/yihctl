package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	log "github.com/sirupsen/logrus"
	"github.com/yishuida/yihctl/pkg/comon"
	myLog "github.com/yishuida/yihctl/pkg/log"
	"github.com/yishuida/yihctl/pkg/util"
	"os"
)

var (
	gitLogger = &log.Entry{}
)

func init() {
	gitLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "git",
	})
}



func CloneOrPull(url, path string) error {
	if util.EmptyDir(path) {
		return Clone(url, path)
	}
	return Pull(path)
}

func Clone(url, path string) error {
	// 当目标路径不为空时，clone 代码
	if util.EmptyDir(path) {
		gitLogger.Infof("Staring clone repo from %s to %s", url, path)
		var err error
		if comon.RegGitHttpUrl.MatchString(url) {
			_, err = git.PlainClone(path, false, &git.CloneOptions{
				URL: url,
				Progress: os.Stdout,
			})
		} else {
			publicKeys, err := ssh.NewPublicKeysFromFile("git", "/Users/vista/.ssh/id_rsa", "")
			if err != nil {
				gitLogger.Warn(err)
			}
			_, err = git.PlainClone(path, false, &git.CloneOptions{
				Auth:     publicKeys,
				URL:      url,
				Progress: os.Stdout,
			})
		}

		if err != nil {
			return err
		}
	} else {
		gitLogger.Infof("Skipup clone repo, %s is existing", path)
	}
	return nil
}

func Pull(path string) error {
	gitLogger.Infof("Staring pull repo in path %s", path)

	r, err := git.PlainOpen(path)
	if err != nil || r == nil {
		gitLogger.Warningf("go-git open %s failed", path)
		return err
	}
	w, err := r.Worktree()
	if err != nil || w == nil {
		gitLogger.Warningf("go-git worktree %s failed", path)
		return err
	}
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		gitLogger.Warningf("go-git pull %s failed", path)
		return err

	}

	ref, err := r.Head()
	if err != nil {
		gitLogger.Warningf("go-git head %s failed", path)
		return err
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		gitLogger.Warningf("go-git head %s failed", path)
		return err
	}

	gitLogger.Infof("now %s commit is %s", path, commit)

	return nil
}