package config

const httpUrlTpl = "%s://%s/%s/%s.git"
const sshUrlTpl = "%s@%s:%s/%s.git"

const DefaultGitRepo = `
name: git-repo
version: 0.1.0
remotes:
- domain: github.com
  scheme: https
  dir: project
  organizations:
  - name: yishuida
    dir: yih
    repos:
    - name: yihctl
    - name: ficus-virens
    - name: larch-framework
    - name: larch-parent
    - name: Dockerfile
    - name: peony
    - name: mihua
  - name: helm
    dir: openSource/helm
    repos:
    - name: helm
    - name: charts
    - name: chartmuseums
  - name: kubernetes
    dir: openSource/k8s
    repos:
    - name: kubernetes
    - name: kube-controller-manager
    - name: kube-scheduler
    - name: kube-proxy
    - name: metrics
    - name: apiextensions-apiserver
    - name: apiserver
    - name: client-go
    - name: api
    - name: kubelet
    - name: kubectl
    - name: kube-state-metrics
  - name: spring-projects
    dir: openSource/spring
    repos:
    - name: spring-boot
    #- name: spring-framework
    - name: spring-security
    - name: spring-data-jdbc
  - name: choerodon
    dir: c7n
    repos:
    - name: javabase
    - name: eureka-server
    - name: go-register-server
    - name: iam-service
    - name: manager-service
    - name: asgard-service
    - name: notify-service
    - name: api-gateway
    - name: oauth-server
    - name: file-service
    - name: choerodon-front
    - name: devops-serviceo
    - name: gitlab-service
    - name: workflow-service
    - name: agile-service
    - name: state-machine-service
    - name: issue-service
    - name: foundation-service
    - name: knowledgebase-service
    - name: test-manager-service
`
