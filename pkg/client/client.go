package client

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// utilities for kubernetes integration
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// InitClient - Kubernetes Client
func InitClient() *kubernetes.Clientset {
	// determine which kubeconfig to use
	var kubeConfig string

	if home := homeDir(); home != "" {
		kubeConfig = filepath.Join(home, ".kube", "config")
	}

	if _, err := os.Stat(kubeConfig); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Unable to open %s\n", kubeConfig)
		os.Exit(1)
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Unable to get config from %s\n", kubeConfig)
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

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
