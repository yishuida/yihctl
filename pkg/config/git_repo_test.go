package config

import (
	"gopkg.in/yaml.v2"
	"testing"
)

const DefaultGitRepo = `name: git-repo
version: 0.1.0

remotes:
  github:
    name: origin
    domain: github.com
    scheme: https
    type: github
  github-ssh:
    name: origin
    domain: github.com
    scheme: ssh
    type: github
  gitee:
    name: origin
    domain: gitee.com
    scheme: https
    type: gitee
  gitlab:
    name: origin
    domain: gitlab.com
    scheme: https
    type: gitlab
  vista:
    name: vista
    domain: code.ydq.io
    scheme: ssh
    type: gitea
repositories:
  - name: "github-ssh-sync-gitea"
    description: "同步 vista 在 github 中的私有仓库到 Vista's Gitea"
    from: github-ssh
    to: vista
    path: project/vista
    repos:
    - source: yidaqiang/yihctl
      target: vista/yihctl
    - source: yidaqiang/script-tool
      target: vista/script-tool
    - source: yishuida/larch-framework
      target: vista/larch-framework
    - source: yishuida/peony
      target: vista/peony
    - source: yishuida/mihua
      target: vista/mihua
    - source: ficus-virens/ficus-virens
      target: ficus-virens/ficus-virens
  - name: "gitee-sync-gitea"
    descriptio: "同步 gitee 的汉得开源代码到 Vista's Gitea"
    from: gitee
    to: vista
    path: project/hand/choerodon
    repos:
      - source: open-hand/agile-service
        target: choerodon/agile-service
      - source: open-hand/devops-service
        target: choerodon/devops-service
  - name: "gitlab-sync-gitea"
    description: "同步 Gitlab 源代码到 Vista's Gitea"
    from: gitlab
    to: vista
    path: project/open-souce/gitlab
    repos:
      - source: gitlab-org/gitlab-foss
        target: software/gitlab-foss
  - name: "github-lang-sync-gitea"
    description: "同步 github 编程语言相关的源代码到 Vista's Gitea"
    from: github
    to: vista
    path: project/open-source/programming-language
    repos:
      - source: rust-lang/rust
        target: programming-language/rust
      - source: golang/go
        target: programming-language/go
  - name: "github-container-sync-gitea"
    description: "同步 github 容器化相关的源代码到 Vista's Gitea"
    from: github
    to: vista
    path: project/open-source/container
    repos:
      - source: kubernetes/kubernetes
        target: container/kubernetes
      - source: k3s-io/k3s
        target: container/k3s
      - source: containerd/containerd
        target: container/containerd
`

func TestBuildGitRepoConfig(t *testing.T) {
	var grc GitRepoConfig
	err := yaml.Unmarshal([]byte(DefaultGitRepo), &grc)
	if err != nil {
		t.Error(err)
	}
	if grc.Name == "" {
		t.Error("GitRepoConfig is empty")
	}
}
