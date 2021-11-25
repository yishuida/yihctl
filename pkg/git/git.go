package git

import (
	"crypto/tls"
	"github.com/go-git/go-git/v5"
	examples "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/client"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"

	"io"
	"net/http"
	"time"
	myLog "ydq.io/yihctl/pkg/log"
	"ydq.io/yihctl/pkg/util"
)

var (
	gitLogger = &log.Entry{}
)

func init() {
	gitLogger = myLog.StdLog.WithFields(log.Fields{
		"pkg": "git",
	})
}

func Clone(out io.Writer, auth transport.AuthMethod, url, path string) error {
	// 当目标路径为空时，在当前路径 clone 代码
	if util.EmptyDir(path) {
		// path, _ = os.Getwd()
		gitLogger.Warningf("git clone path is empty, new clone into %s", path)
	}

	gitLogger.Infof("Staring clone repository from %s into %s", url, path)

	option := &git.CloneOptions{
		Auth:     auth,
		URL:      url,
		Progress: out,
	}

	_, err := git.PlainClone(path, false, option)
	examples.CheckIfError(err)

	return nil
}

func Pull(out io.Writer, auth transport.AuthMethod, path, remoteName string) error {
	gitLogger.Infof("Staring pull repo in path %s", path)

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	if err != nil || r == nil {
		gitLogger.Warningf("go-git open %s failed", path)
		return err
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil || w == nil {
		gitLogger.Warningf("go-git worktree %s failed", path)
		return err
	}

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{
		RemoteName: remoteName,
		Progress:   out,
		Auth:       auth,
	})
	if err != nil && err.Error() != "already up-to-date" {
		gitLogger.Warningf("go-git pull %s failed", path)
		return err
	}

	// Print the latest commit that was just pulled
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

func Push(out io.Writer, auth transport.AuthMethod, path, remoteName string) error {

	// Create a custom http(s) client with your config
	customClient := &http.Client{
		// accept any certificate (might be useful for testing)
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},

		// 15 second timeout
		Timeout: 15 * time.Second,

		// don't follow redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// Override http(s) default protocol to use our custom client
	client.InstallProtocol("https", githttp.NewClient(customClient))

	gitLogger.Infof("Staring push repo in path %s", path)

	r, err := git.PlainOpen(path)
	if err != nil || r == nil {
		gitLogger.Warningf("go-git open %s failed", path)
		return err
	}

	// push using default options
	err = r.Push(&git.PushOptions{
		RemoteName: remoteName,
		Progress:   out,
		Auth:       auth,
	})
	if err != nil {
		gitLogger.Warningf("go-git push %s failed", path)
		return err
	}
	return nil
}

func AddRemote(path, remoteName, url string) error {
	gitLogger.Infof("Staring add repo remote in path %s", path)

	r, err := git.PlainOpen(path)
	if err != nil || r == nil {
		gitLogger.Warningf("go-git open %s failed", path)
		return err
	}

	remoteList, err := r.Remotes()
	if err != nil {
		gitLogger.Error(err)
	}

	for _, remote := range remoteList {
		rc := remote.Config()
		if rc.Name == remoteName && rc.URLs[0] == url {
			return nil
		}
	}
	err = r.DeleteRemote(remoteName)
	if err != nil {
		gitLogger.Warningf("go-git add remote %s failed", path)
		gitLogger.Error(err)
	}
	// Add a new remote
	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{url},
	})
	if err != nil {
		gitLogger.Warningf("go-git add remote %s failed", path)
		return err
	}
	return nil
}
