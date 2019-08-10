package client

import (
	"flag"
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
	var kubeconfig *string
	var kubeconfigbase string

	if home := homeDir(); home != "" {
		kubeconfigbase = filepath.Join(home, ".kube", "config")
	}

	kubeconfig = flag.String(
		"kubeconfig",
		kubeconfigbase,
		"(optional) absolute path to the kubeconfig file",
	)
	flag.Parse()

	if _, err := os.Stat(*kubeconfig); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Unable to open %s\n", *kubeconfig)
		os.Exit(1)
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal: Unable to get config from %s\n", *kubeconfig)
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
