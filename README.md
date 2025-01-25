# Kubectl Grep

[![GitHub Actions](https://github.com/guessi/kubectl-grep/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/guessi/kubectl-grep/actions/workflows/go.yml)
[![GoDoc](https://godoc.org/github.com/guessi/kubectl-grep?status.svg)](https://godoc.org/github.com/guessi/kubectl-grep)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/kubectl-grep)](https://goreportcard.com/report/github.com/guessi/kubectl-grep)
[![GitHub release](https://img.shields.io/github/release/guessi/kubectl-grep.svg)](https://github.com/guessi/kubectl-grep/releases/latest)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guessi/kubectl-grep)](https://github.com/guessi/kubectl-grep/blob/main/go.mod)

Filter Kubernetes resources by matching their names

## 🤔 Why we need this? what it is trying to resolve?

Playing with Kubernetes in our daily job, we normally search pods by `pipe`, `grep`, `--label`, `--field-selector`, etc. while hunting abnormal pods, but typing such long commands is quite annoying. With plugin installed, it could be easily be done by `kubectl grep`.

## 🔢 Prerequisites

* An existing Kubernetes cluster.
* Have [kubectl](https://kubernetes.io/docs/tasks/tools/) installed.
* Have [krew](https://krew.sigs.k8s.io/docs/user-guide/setup/install/) installed.

## 🚀 Quick start

```bash
# Have krew plugin installed
kubectl krew install grep
```

```bash
# Before change, we usually filter pods by the following commands,
kubectl get pods -n star-lab | grep "flash"
```

```bash
# With this plugin installed, you can filter pod easily
kubectl grep pods -n star-lab flash
```

## :accessibility: FAQ

How do I check the version installed?

* `kubectl grep version`

How do I know the version installed is compatible with my cluster version?

* Ideally, it should be compatible with [supported Kubernetes versions](https://kubernetes.io/releases/).

What kind of resource(s) `kubectl-grep` support?

* Please refer to [Resource Types](RESOURCE_TYPES.md)

## 👷 Install

### Recommended way

```bash
kubectl krew install grep && kubectl krew update && kubectl krew upgrade grep
```

### Manual Installation

<details><!-- markdownlint-disable-line -->
<summary>Click to expand!</summary><!-- markdownlint-disable-line -->

```bash
curl -fsSL -O https://github.com/guessi/kubectl-grep/releases/latest/download/kubectl-grep-$(uname -s)-$(uname -m).tar.gz
tar zxvf kubectl-grep-$(uname -s)-$(uname -m).tar.gz
mv kubectl-grep /usr/local/bin
```

</details>

### Developer build

<details><!-- markdownlint-disable-line -->
<summary>Click to expand!</summary><!-- markdownlint-disable-line -->
  
```bash
go get -u github.com/guessi/kubectl-grep
cd ${GOPATH}/src/github.com/guessi/kubectl-grep
make all
```

</details>

## ⚖️ License

[Apache-2.0](LICENSE)
