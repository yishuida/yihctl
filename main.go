package main

import (
	"context"
	"gitee.com/openeuler/go-gitee/gitee"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"regexp"
)

func main() {
	log.Println(GetTag("0.23", "devops-version"))
}

func GetTag(version, repo string) (targetVersion string) {

	client := github.NewClient(nil)

	tags, resp, err := client.Repositories.ListTags(context.Background(), "open-hand", repo, &github.ListOptions{})
	if err != nil {
		log.Panic(err)
	}
	log.Info(resp)
	reg := regexp.MustCompile("^" + version + ".\\d$")
	for _, tag := range tags {
		tagName := *tag.Name
		if reg.MatchString(tagName) {
			if targetVersion == "" {
				targetVersion = VersionOrdinal(tagName)
			}
			if targetVersion < VersionOrdinal(tagName) {
				targetVersion = VersionOrdinal(tagName)
			}
		}
	}
	return targetVersion
}

func VersionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}

func giteeGetTags(repo, version string) {
	client := gitee.NewAPIClient(nil)
	tag, resp, err := client.RepositoriesApi.GetV5ReposOwnerRepoTags(context.Background(), "open-hand", repo, &gitee.GetV5ReposOwnerRepoTagsOpts{})
	if err != nil {
		log.Panic(err)
	}
	log.Info(resp)
	log.Println(tag)
}
