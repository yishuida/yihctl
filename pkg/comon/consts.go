package comon

import "regexp"

// TODO 改成完整的 URL 校验
const gitHttpUrlRegexp = "^http(s)?"

var RegGitHttpUrl = regexp.MustCompile(gitHttpUrlRegexp)
