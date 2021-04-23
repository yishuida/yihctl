package util

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

// 返回 home 目录
func HomePath() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
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
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			utilLogger.Error(err)
		}
	}
}

func FileIsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		utilLogger.WithFields(log.Fields{"error": err, "file": filePath}).Warn("file isn't exist.")
		return false
	}
	return true
}

func EmptyDir(dir string) bool {
	s, err := ioutil.ReadDir(dir)
	if err != nil {
		utilLogger.Error(err)
	}
	return len(s) == 0
}
