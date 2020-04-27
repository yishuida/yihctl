package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/helm"
	helm_env "k8s.io/helm/pkg/helm/environment"
	"k8s.io/helm/pkg/helm/portforwarder"
	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/tlsutil"
	"os"
)

var (
	tillerTunnel *kube.Tunnel
	settings     helm_env.EnvSettings
)

func newHelmCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "helm",
		Short:   "asdf",
		PreRunE: func(_ *cobra.Command, _ []string) error { return setupConnection() },
		PersistentPreRun: func(*cobra.Command, []string) {
			if settings.TLSCaCertFile == helm_env.DefaultTLSCaCert || settings.TLSCaCertFile == "" {
				settings.TLSCaCertFile = settings.Home.TLSCaCert()
			} else {
				settings.TLSCaCertFile = os.ExpandEnv(settings.TLSCaCertFile)
			}
			if settings.TLSCertFile == helm_env.DefaultTLSCert || settings.TLSCertFile == "" {
				settings.TLSCertFile = settings.Home.TLSCert()
			} else {
				settings.TLSCertFile = os.ExpandEnv(settings.TLSCertFile)
			}
			if settings.TLSKeyFile == helm_env.DefaultTLSKeyFile || settings.TLSKeyFile == "" {
				settings.TLSKeyFile = settings.Home.TLSKey()
			} else {
				settings.TLSKeyFile = os.ExpandEnv(settings.TLSKeyFile)
			}
		},
		PersistentPostRun: func(*cobra.Command, []string) {
			teardown()
		},
		Run: func(cmd *cobra.Command, args []string) {
			hc := newClient()
			resp, _ := hc.ListReleases()
			for _, r := range resp.Releases {
				fmt.Print(r.GetName())
			}
		},
	}
	flags := cmd.PersistentFlags()
	settings.AddFlags(flags)

	flags.Parse(args)

	// set defaults from environment
	settings.Init(flags)
	return cmd
}

func setupConnection() error {
	if settings.TillerHost == "" {
		config, client, err := getKubeClient(settings.KubeContext, settings.KubeConfig)
		if err != nil {
			return err
		}

		tillerTunnel, err = portforwarder.New(settings.TillerNamespace, client, config)
		if err != nil {
			return err
		}

		settings.TillerHost = fmt.Sprintf("127.0.0.1:%d", tillerTunnel.Local)
		// debug("Created tunnel using local port: '%d'\n", tillerTunnel.Local)
	}

	// Set up the gRPC config.
	// debug("SERVER: %q\n", settings.TillerHost)

	// Plugin support.
	return nil
}

func teardown() {
	if tillerTunnel != nil {
		tillerTunnel.Close()
	}
}

// configForContext creates a Kubernetes REST client configuration for a given kubeconfig context.
func configForContext(context string, kubeconfig string) (*rest.Config, error) {
	config, err := kube.GetConfig(context, kubeconfig).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("could not get Kubernetes config for context %q: %s", context, err)
	}
	return config, nil
}

// getKubeClient creates a Kubernetes config and client for a given kubeconfig context.
func getKubeClient(context string, kubeconfig string) (*rest.Config, kubernetes.Interface, error) {
	config, err := configForContext(context, kubeconfig)
	if err != nil {
		return nil, nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get Kubernetes client: %s", err)
	}
	return config, client, nil
}

// ensureHelmClient returns a new helm client impl. if h is not nil.
func ensureHelmClient(h helm.Interface) helm.Interface {
	if h != nil {
		return h
	}
	return newClient()
}

func newClient() helm.Interface {
	options := []helm.Option{helm.Host(settings.TillerHost), helm.ConnectTimeout(settings.TillerConnectionTimeout)}

	if settings.TLSVerify || settings.TLSEnable {
		// debug("Host=%q, Key=%q, Cert=%q, CA=%q\n", settings.TLSServerName, settings.TLSKeyFile, settings.TLSCertFile, settings.TLSCaCertFile)
		tlsopts := tlsutil.Options{
			ServerName:         settings.TLSServerName,
			KeyFile:            settings.TLSKeyFile,
			CertFile:           settings.TLSCertFile,
			InsecureSkipVerify: true,
		}
		if settings.TLSVerify {
			tlsopts.CaCertFile = settings.TLSCaCertFile
			tlsopts.InsecureSkipVerify = false
		}
		tlscfg, err := tlsutil.ClientConfig(tlsopts)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		options = append(options, helm.WithTLS(tlscfg))
	}
	return helm.NewClient(options...)
}
