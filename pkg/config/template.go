package config

const DefaultGitRepo = `
name: git-repo
version: 0.1.0
remotes:
- domain: https://github.com
  organizations:
  - name: yishuida
    dir: project/yih
    repos:
    - name: yihctl
  - name: choerodon
    dir: workspace/c7n
    repos:
    - name: javabase
`
