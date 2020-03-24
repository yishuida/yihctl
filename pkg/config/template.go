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
    - name: ficus-virens
    - name: larch-framework
    - name: larch-parent
    - name: Dockerfile
    - name: peony
    - name: mihua
  - name: helm
    dir: project/openSource/helm
    repos:
    - name: helm
    - name: charts
    - name: chartmuseums
  - name: kubernetes
    dir: project/openSource/k8s
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
    dir: project/openSource/spring
    repos:
    - name: spring-boot
    - name: spring-framework
    - name: spring-security
    - name: spring-data-jdbc
  - name: choerodon
    dir: workspace/c7n
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
    - name: devops-service
    - name: gitlab-service
    - name: workflow-service
    - name: agile-service
    - name: state-machine-service
    - name: issue-service
    - name: foundation-service
    - name: knowledgebase-service
    - name: test-manager-service
`
