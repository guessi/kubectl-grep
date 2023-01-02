# Kubectl Grep

[![GitHub Actions](https://github.com/guessi/kubectl-grep/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/guessi/kubectl-grep/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/guessi/kubectl-grep?status.svg)](https://godoc.org/github.com/guessi/kubectl-grep)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/kubectl-grep)](https://goreportcard.com/report/github.com/guessi/kubectl-grep)
[![GitHub release](https://img.shields.io/github/release/guessi/kubectl-grep.svg)](https://github.com/guessi/kubectl-grep/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/kubectl-grep)](https://github.com/guessi/kubectl-grep/blob/master/go.mod)

Filter Kubernetes resources by matching their names

# Requirements

- Kubernetes 1.23.0+
- Kubectl 1.23.0+
- Krew 0.4.3+

# Compatibility

please refer to [Kubernetes version policy](https://kubernetes.io/docs/setup/release/version-skew-policy/#kubectl) and [CHANGELOG](CHANGELOG.md) for supported version matrix.

Example: `kubectl-grep` build with Kubernetes-1.24.x should be compatable with Kubernetes cluster version 1.23, 1.24, 1.25.

# Why we need it?

playing with Kubernetes is my daily job, and I normally search pods by `pipe`,
`grep`, `--label`, `--field-selector`, etc. while hunting abnormal pods, but
typing such long commands is quite annoying.

Before change, we usually filter pods by the following commands,

    $ kubectl get pods -n star-lab | grep "flash"

With this plugin installed, you can filter pod with `kubectl grep` easily

    $ kubectl grep pods -n star-lab flash

# Installation

Installation via [krew](https://krew.sigs.k8s.io/docs/user-guide/setup/install/)

    $ kubectl krew version # make sure you are running 0.4.3+
    $ kubectl krew install grep
    $ kubectl krew upgrade

Manual Installation

    $ export KUBECTL_GREP_VERSION=$(curl -s https://api.github.com/repos/guessi/kubectl-grep/releases/latest | jq -r .tag_name)
    $ curl -L -O https://github.com/guessi/kubectl-grep/releases/download/${KUBECTL_GREP_VERSION}/kubectl-grep-$(uname -s)-$(uname -m).tar.gz
    $ tar zxvf kubectl-grep-$(uname -s)-$(uname -m).tar.gz
    $ mv kubectl-grep /usr/local/bin
    $ chmod +x /usr/local/bin/kubectl-grep

# How to get developer build?

    $ go get -u github.com/guessi/kubectl-grep

    $ cd ${GOPATH}/src/github.com/guessi/kubectl-grep

    $ make all

# FAQ

How do I check the tool's version?

* `kubectl grep version`

Can I use version X `kubectl` with version Y `kubectl-grep`?

* Sure, no problem

What kind of resource(s) `kubectl-grep` support?

* Please refer to [Resource Types](RESOURCE_TYPES.md)

# Reference

- [Kubectl Plugins](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/)

# License

[Apache-2.0](LICENSE)
