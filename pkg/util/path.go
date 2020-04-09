package util

import (
	"github.com/mitchellh/go-homedir"
	"os"
)

// 返回 home 目录
func HomePath() string {
	home, err := homedir.Dir()
	if err != nil {
		utilLogger.Error(err)
	}
	return home
}

func Path(paths ...string) string {
	var resultPath = ""

	for _, path := range paths {
		if path != "" {
			resultPath += path + string(os.PathSeparator)
		}
	}
	ensurePath(resultPath)

	return resultPath
}

func ensurePath(path string) {
	// when path is no exist, create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		utilLogger.Warn(err)
		if err := os.MkdirAll(path, 764); err != nil {
			utilLogger.Error(err)
		}
	}
}
