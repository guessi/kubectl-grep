package client

import (
	"flag"
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

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
