package client

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// utilities for kubernetes integration
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// ClientConfig
func ClientConfig() clientcmd.ClientConfig {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{})
}

// InitClient - Kubernetes Client
func InitClient() *kubernetes.Clientset {
	clientConfig := ClientConfig()
	config, err := clientConfig.ClientConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Unable to get config\n")
		os.Exit(1)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Unable to get create client\n")
		os.Exit(1)
	}

	return clientset
}
