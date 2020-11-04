module github.com/yishuida/yihctl

go 1.14

require (
	gitee.com/openeuler/go-gitee v0.0.0-20200918065743-cef3fb7bc147
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/jinzhu/configor v1.1.1
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rubenv/sql-migrate v0.0.0-20200402132117-435005d389bc // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	gopkg.in/src-d/go-git.v4 v4.13.1
	k8s.io/apiextensions-apiserver v0.16.5 // indirect
	k8s.io/client-go v0.16.5
	k8s.io/helm v2.16.5+incompatible
	k8s.io/kubernetes v1.16.5 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.16.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.5
	k8s.io/apiserver => k8s.io/apiserver v0.16.5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.5
	k8s.io/client-go => k8s.io/client-go v0.16.5
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.16.5
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.16.5
	k8s.io/code-generator => k8s.io/code-generator v0.16.5
	k8s.io/component-base => k8s.io/component-base v0.16.5
	k8s.io/cri-api => k8s.io/cri-api v0.16.5
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.16.5
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.16.5
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.16.5
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.5
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.16.5
	k8s.io/kubectl => k8s.io/kubectl v0.16.5
	k8s.io/kubelet => k8s.io/kubelet v0.16.5
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.16.5
	k8s.io/metrics => k8s.io/metrics v0.16.5
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.16.5
)
