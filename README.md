# Kubectl Search

[![CircleCI](https://circleci.com/gh/guessi/kubectl-search.svg?style=svg)](https://circleci.com/gh/guessi/kubectl-search)
[![GoDoc](https://godoc.org/github.com/guessi/kubectl-search?status.svg)](https://godoc.org/github.com/guessi/kubectl-search)
[![Go Report Card](https://goreportcard.com/badge/github.com/guessi/kubectl-search)](https://goreportcard.com/report/github.com/guessi/kubectl-search)
[![GitHub release](https://img.shields.io/github/release/guessi/kubectl-search.svg)](https://github.com/guessi/kubectl-search/releases/latest)

simple kubernetes plugins for searching resources with keyword

# Requirements

- Kubernetes 1.10.0+
- Kubectl 1.13.0+

# Why we need it?

playing with Kubernetes is my daily job, and I normally search pods by `pipe`,
`grep`, `--label`, `--field-selector`, etc. while hunting abnormal pods, but
typing such long commands are quite annoyed.

Before change, we usually search pods by the following commands,

    $ kubectl get pods                                 | grep "keyword"
    $ kubectl get pods -n "my-ns"                      | grep "keyword"
    $ kubectl get pods --all-namespaces                | grep "keyword"
    $ kubectl get pods -l "key=value"                  | grep "keyword"
    $ kubectl get pods -l "key=value" -n "my-ns"       | grep "keyword"
    $ kubectl get pods -l "key=value" --all-namespaces | grep "keyword"
    $ kubectl get pods --field-selector "key=value"    | grep "keyword"

With this plugin installed, you can search pod with `kubectl search` easily

    $ kubectl search pods "keyword"
    $ kubectl search pods "keyword" -n "my-ns"
    $ kubectl search pods "keyword" --all-namespaces
    $ kubectl search pods "keyword" -l "key=value"
    $ kubectl search pods "keyword" -l "key=value" -n "my-ns"
    $ kubectl search pods "keyword" -l "key=value" --all-namespaces
    $ kubectl search pods "keyword" --field-selector "key=value"

# Basic Usage

    $ curl -L https://github.com/guessi/kubectl-search/releases/download/v1.0.0/kubectl-search-`uname -s`-`uname -m` -o /usr/local/bin/kubectl-search
    $ chmod +x /usr/local/bin/kubectl-search
    $ kubectl search pods "keyword"

# How to get developer build?

    $ go get -u github.com/guessi/kubectl-search

    $ cd ${GOPATH}/src/github.com/guessi/kubectl-search

    $ glide install --strip-vendor

    $ make clean all

# FAQ

What kind of resource does current release supported?

    now it support Pods, Deployments, and HPA

Is it possible to search StatefulSets, DaemonSets, etc?

    sure, but not yet supported, I need more time to write code, and tests

How do I check the tool's version?

    $ kubectl search version

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
