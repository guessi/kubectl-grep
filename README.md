# Kubectl Grep

[![CircleCI](https://img.shields.io/circleci/build/github/guessi/kubectl-grep?token=39cb1435f15bbd1fd965997db3484aed48d458b0)](https://circleci.com/gh/guessi/kubectl-grep)
[![GoDoc](https://godoc.org/github.com/guessi/kubectl-grep?status.svg)](https://godoc.org/github.com/guessi/kubectl-grep)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/kubectl-grep)](https://goreportcard.com/report/github.com/guessi/kubectl-grep)
[![GitHub release](https://img.shields.io/github/release/guessi/kubectl-grep.svg)](https://github.com/guessi/kubectl-grep/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/kubectl-grep)](https://github.com/guessi/kubectl-grep/blob/master/go.mod)

[![GitHub watchers](https://img.shields.io/github/watchers/guessi/kubectl-grep?style=social)](https://github.com/guessi/kubectl-grep/watchers)
[![GitHub Repo stars](https://img.shields.io/github/stars/guessi/kubectl-grep?style=social)](https://github.com/guessi/kubectl-grep/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/guessi/kubectl-grep?style=social)](https://github.com/guessi/kubectl-grep/network/members)

Filter Kubernetes resources by matching their names

# Requirements

- Kubernetes 1.10.0+
- Kubectl 1.13.0+
- Krew 0.4.0+

# Compatibility

please refer to [Kubernetes version policy](https://kubernetes.io/docs/setup/release/version-skew-policy/#kubectl) and [CHANGELOG](CHANGELOG.md) for supported version matrix.

Example: `kubectl-grep` build with Kubernetes-1.13.x should be compatable with Kubernetes cluster version 1.12, 1.13, 1.14.

# Why we need it?

playing with Kubernetes is my daily job, and I normally search pods by `pipe`,
`grep`, `--label`, `--field-selector`, etc. while hunting abnormal pods, but
typing such long commands are quite annoyed.

Before change, we usually filter pods by the following commands,

    $ kubectl get pods                           | grep "keyword"
    $ kubectl get pods -o wide                   | grep "keyword"
    $ kubectl get pods -n "my-ns"                | grep "keyword"
    $ kubectl get pods -A                        | grep "keyword"
    $ kubectl get pods -l "key=value"            | grep "keyword"
    $ kubectl get pods -l "key=value" -n "my-ns" | grep "keyword"
    $ kubectl get pods -l "key=value" -A         | grep "keyword"

With this plugin installed, you can filter pod with `kubectl grep` easily

    $ kubectl grep pods "keyword"
    $ kubectl grep pods "keyword" -o wide
    $ kubectl grep pods "keyword" -n "my-ns"
    $ kubectl grep pods "keyword" -A
    $ kubectl grep pods "keyword" -l "key=value"
    $ kubectl grep pods "keyword" -l "key=value" -n "my-ns"
    $ kubectl grep pods "keyword" -l "key=value" -A

# Supported Resources

- [X] ConfigMaps
- [X] DaemonSets
- [X] Deployments
- [X] HPAs
- [X] Ingresses
- [X] Jobs
- [X] Nodes
- [X] Pods
- [X] Secrets
- [X] StatefulSets

# Installation

Installation via [krew](https://krew.sigs.k8s.io/docs/user-guide/setup/install/)

    $ kubectl krew version # make sure you are running 0.4.0+
    $ kubectl krew install grep
    $ kubectl krew upgrade

Manual Installation

    $ export KUBECTL_GREP_VERSION=$(curl -s https://api.github.com/repos/guessi/kubectl-grep/releases/latest | jq -r .tag_name)
    $ curl -L -O https://github.com/guessi/kubectl-grep/releases/download/${KUBECTL_GREP_VERSION}/kubectl-grep-$(uname -s)-$(uname -m).tar.gz
    $ tar zxvf kubectl-grep-$(uname -s)-$(uname -m).tar.gz
    $ mv kubectl-grep /usr/local/bin
    $ chmod +x /usr/local/bin/kubectl-grep

# Examples

List all pods in default namespace,

    $ kubectl grep pods

List all pods in all namespaces,

    $ kubectl grep pods -A

List all pods with specific keyword, under specific namespace,

    $ kubectl grep pods -n star-lab flash

# How to get developer build?

    $ go get -u github.com/guessi/kubectl-grep

    $ cd ${GOPATH}/src/github.com/guessi/kubectl-grep

    $ make all

# FAQ

How do I check the tool's version?

    $ kubectl grep version

Any plan to support Kubernetes version before 1.10.0?

    sorry, it only support Kubernetes 1.10.0+

I'm now running Kubernetes 1.10.0, do I need to upgrade my cluster?

    nope, the only requirement is to upgrade your `kubectl` to 1.13.0+

Can I run Kubernetes 1.12.0 with kubectl 1.13.0?

    sure, no problem

# Reference

- [Kubectl Plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/)

# License

[Apache-2.0](LICENSE)
