# v1.20.0 / 2025-01-25

* Build with Kubernetes Client SDK v1.32.1
* Build with go 1.23
* Bump dependencies

# v1.19.2 / 2024-12-16

* Build with Kubernetes Client SDK v1.31.3
* Bump dependencies

# v1.19.1 / 2024-11-12

* Implement --invert-match/-v support
* Bump dependencies

# v1.19.0 / 2024-11-03

* Build with Kubernetes Client SDK v1.31.2
* Bump dependencies

# v1.18.3 / 2024-09-30

* Build with Kubernetes Client SDK v1.30.5
* Bump dependencies

# v1.18.2 / 2024-08-15

* Build with Kubernetes Client SDK v1.30.4
* Bump dependencies

# v1.18.1 / 2024-08-02

* Build with Kubernetes Client SDK v1.30.3
* Bump dependencies

# v1.18.0 / 2024-06-15

* Build with go 1.22
* Build with Kubernetes Client SDK v1.30.2
* Bump github.com/spf13/cobra v1.8.1

# v1.17.1 / 2024-04-26

* Build with Kubernetes Client SDK v1.29.4
* Bump dependencies

# v1.17.0 / 2024-03-24

* Build with Kubernetes Client SDK v1.29.3
* Bump dependencies

# v1.16.1 / 2024-02-16

* Build with Kubernetes Client SDK v1.28.7
* Bump actions/cache@v4

# v1.16.0 / 2024-01-07

* Build with Kubernetes Client SDK v1.28.5
* Bump actions/steup-go@v5
* Bump actions/stale@v9

# v1.15.3 / 2023-11-19

* Build with Kubernetes Client SDK v1.27.8
* Bump github.com/spf13/cobra v1.8.0

# v1.15.2 / 2023-11-04

* Build with Kubernetes Client SDK v1.27.7
* Bump actions/checkout@v4

# v1.15.1 / 2023-09-29

* Build with Kubernetes Client SDK v1.27.6

# v1.15.0 / 2023-08-27

* Upstream k8s 1.27 built with go 1.20
* Build with Kubernetes Client SDK v1.27.5
* More detail for staticcheck
* Bump github.com/sirupsen/logrus v1.9.3
* Bump github.com/spf13/cobra v1.7.0

# v1.14.1 / 2023-08-18

* Build with Kubernetes Client SDK v1.26.7

# v1.14.0 / 2023-07-23

* Build with Kubernetes Client SDK v1.26.6

# v1.13.2 / 2023-06-15

* Build with Kubernetes Client SDK v1.25.11

# v1.13.1 / 2023-05-26

* Build with Kubernetes Client SDK v1.25.10

# v1.13.0 / 2023-04-19

* Release with krew-release-bot v0.0.46
* Build with Kubernetes Client SDK v1.25.9
* Keep output aligned with kubectl and bug fixes
* Trigger actions/stale once per week
* Bump actions/stale@v8

# v1.12.3 / 2023-03-26

* Base: build with Kubernetes Client SDK v1.24.12
* Misc: release with krew-release-bot v0.0.44
* Misc: bump actions/setup-go@v4

# v1.12.2 / 2023-03-12

* Base: build with Kubernetes Client SDK v1.24.11
* Address vulnerabilities fixes

# v1.12.1 / 2023-02-18

* Address vulnerabilities
  * https://pkg.go.dev/vuln/GO-2022-0619
  * https://pkg.go.dev/vuln/GO-2023-1571
  * https://pkg.go.dev/vuln/GO-2023-1570
  * https://groups.google.com/g/golang-announce/c/V0aBFqaFs_E

# v1.12.0 / 2023-01-25

* CI: upgrade actions/stale@v7
* Base: build with Kubernetes Client SDK v1.24.10 and Golang 1.19
  * https://github.com/kubernetes/kubernetes/pull/113956
  * https://github.com/kubernetes/kubernetes/pull/115012

# v1.11.0 / 2023-01-03

* Release with GitHub Actions

# v1.10.4 / 2022-12-11

* Base: bump Kubernetes Client SDK v1.24.9
  * Mitigated CVEs: CVE-2022-41717, CVE-2022-27664

# v1.10.3 / 2022-11-14

* Base: bump Kubernetes Client SDK v1.24.8
  * Mitigated CVEs: [CVE-2022-3162](https://discuss.kubernetes.io/t/security-advisory-cve-2022-3162-unauthorized-read-of-custom-resources/21902), [CVE-2022-3294](https://discuss.kubernetes.io/t/security-advisory-cve-2022-3294-node-address-isnt-always-verified-when-proxying/21903)
* Upgrade dependencies:
  - spf13/cobra v1.6.1

# v1.10.2 / 2022-10-16

* Bump Kubernetes Client SDK v1.24.7

# v1.10.1 / 2022-09-17

* Build with golang 1.18.x (CI)
* Replace deprecated golint with Staticcheck
* Download ghr binary directly

# v1.10.0 / 2022-09-17

* Bump Kubernetes Client SDK v1.24.5
* Build with golang 1.18.x
* Include more info for "version" command
* Introduce stable bot

**Known issue:** it is actually build with golang 1.17, please use v1.10.1 instead.

# v1.9.0 / 2022-09-11

* Added support for the following resources
  - ClusterRoleBindings
  - ClusterRoles
  - RoleBindings
  - Roles
  - ServiceAccounts

# v1.8.0 / 2022-09-06

* Initial support for arm-based Linux
* Initial support for arm-based macOS (M1/M2-series)
* Upgrade dependencies:
  - spf13/cobra v1.5.0
  - sirupsen/logrus v1.9.0

# v1.7.2 / 2022-08-29

* Bump Kubernetes Client SDK: kubernetes-1.23.10
* Bump krew-release-bot v0.0.43

# v1.7.1 / 2022-05-28

* CVE-2022-28948

# v1.7.0 / 2022-05-28

* Bump Kubernetes Client SDK: kubernetes-1.23.7
* Release with krew-release-bot v0.0.42
* Keep output aligned when no resource found

# v1.6.0 / 2022-03-22

* Build with github.com/spf13/cobra v1.4.0
* Build with next-gen convenience image: cimg/go:1.16
* Added support for the following resources
  - CronJobs
  - ReplicaSets

# v1.5.1 / 2022-03-19

* Bump Kubernetes Client SDK: kubernetes-1.22.8
* Cleanup go.mod

# v1.5.0 / 2022-01-23

* Bump Kubernetes Client SDK: kubernetes-1.22.6
* Cleanup go.mod
* Added support for the following resources
  - CSIDrivers
  - StorageClasses

# v1.4.4 / 2022-01-05

* Upgrade krew-release-bot to v0.0.40
* Bump Kubernetes Client SDK: kubernetes-1.21.8
* Upgrade dependencies:
  - spf13/cobra v1.3.0
* Cleanup CircleCI configuration

# v1.4.3 / 2021-11-21

* Bump Kubernetes Client SDK: kubernetes-1.21.7

# v1.4.2 / 2021-10-17

* Fix incorrect node role display

# v1.4.1 / 2021-10-16

* Added support for the following resources
  - Services, thanks to @wshihadeh
* Bump Kubernetes Client SDK: kubernetes-1.21.5

# v1.4.0 / 2021-09-04

* Bump Kubernetes Client SDK: kubernetes-1.21.4
* Build with go 1.16

# v1.3.3 / 2021-06-25

* Bump Kubernetes Client SDK: kubernetes-1.20.8
* Bump golang.org/x/crypto for CVE-2020-29652

# v1.3.2 / 2021-03-26

* NO CHANGE, to fix incorrect sha256sum on krew-index

# v1.3.1 / 2021-03-24

* Introduce krew-release-bot for release automation

# v1.3.0 / 2021-03-15

* Drop support for Kubernetes 1.16 or earlier
* Bump Kubernetes Client SDK: kubernetes-1.20.4
* Upgrade dependencies:
  - spf13/cobra v1.1.3
  - sirupsen/logrus v1.8.1
* Build with go 1.15

# v1.2.7 / 2020-10-25

* Added support for the following resources
  - Jobs
  - Ingresses
* Bump Kubernetes Client SDK: kubernetes-1.18.10
* Upgrade dependencies:
  - spf13/cobra v1.1.1
  - sirupsen/logrus v1.7.0

# v1.2.6 / 2020-08-08

* Upgrade to Go 1.14
* Bump Kubernetes Client SDK: kubernetes-1.18.6
* Upgrade dependencies:
  - spf13/cobra v1.0.0
  - sirupsen/logrus v1.6.0

# v1.2.5 / 2020-05-11

* Bump Kubernetes Client SDK: kubernetes-1.18.2
* Fix logic error for client init process

# v1.2.4 / 2020-03-25

* Bump Kubernetes Client SDK: kubernetes-1.16.8

# v1.2.3 / 2020-01-13

* Bump Kubernetes Client SDK: kubernetes-1.16.4
* Added support for ConfigMaps search, with command aliases:
  - configmaps, configmap, cm
* Added support for Secrets search, with command aliases:
  - secrets, secret

# v1.2.2 / 2019-12-09

* Release with LICENSE, resolved [#16](https://github.com/guessi/kubectl-grep/issues/16)

# v1.2.1 / 2019-11-15

* Bump Kubernetes Client SDK: kubernetes-1.15.6
* Fix KUBECONFIG not working issue, @tsaarni thanks ([#14](https://github.com/guessi/kubectl-grep/pull/14))
* Fix namespace not respect the setting from KUBECONFIG, reported by @fredleger ([#13](https://github.com/guessi/kubectl-grep/issues/13), [#15](https://github.com/guessi/kubectl-grep/pull/15))

# v1.2.0 / 2019-09-15

* Bump Kubernetes Client SDK: kubernetes-1.15.3
* Upgrade to Go 1.13.0
* Added support for StatefulSets search
* Added support for `-A` as the shortcut of `--all-namespaces`
* Added support for install via `kubectl krew install grep`
* Fixed `--help` not work for the root command
* Support for command aliases:
  - daemonsets, daemonset, ds
  - deployments, deployment, deploy
  - hpas, hpa
  - nodes, node, no
  - pods, pod, po
  - statefulsets, stateful, sts

# v1.1.0 / 2019-07-19

* Bump Kubernetes Client SDK: kubernetes-1.15.1
* Cleanup go.mod / go.sum
* *BREAKING CHANGE*: Renamed as `kubectl-grep`

# v1.0.5 / 2019-05-12

* Exit if `.kube/config` not found
* Code refactoring

# v1.0.4 / 2019-04-20

* Added Support for Node Search
* Bump Kubernetes Client SDK: kubernetes-1.14.1

# v1.0.3 / 2019-04-07

* Security Fixes for CVE-2019-1002101 and CVE-2019-9946

# v1.0.2 / 2019-03-22

* Supported NodeName Display for Pods
* Supported Multi-Selector Display for DaemonSets
* Don't Panic if Cluster Unreachable

# v1.0.1 / 2019-03-17

* Added Support for DaemonSets
* Added Support for "-o wide" Option
* Added Support for GO Modules

# v1.0.0 / 2019-03-09

* Initial Release
