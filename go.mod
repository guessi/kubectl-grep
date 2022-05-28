module github.com/guessi/kubectl-grep

go 1.17

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.4.0
	k8s.io/api v0.23.7
	k8s.io/apimachinery v0.23.7
	k8s.io/client-go v0.23.7
)

require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.18 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.13 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/form3tech-oss/jwt-go v3.2.3+incompatible // indirect
	github.com/go-logr/logr v1.2.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.5 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/imdario/mergo v0.3.5 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20211209124913-491a49abca63 // indirect
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f // indirect
	golang.org/x/sys v0.0.0-20210831042530-f4d43177bf5e // indirect
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65 // indirect
	k8s.io/utils v0.0.0-20211116205334-6203023598ed // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.81.0
	cloud.google.com/go/bigquery => cloud.google.com/go/bigquery v1.8.0
	cloud.google.com/go/firestore => cloud.google.com/go/firestore v1.1.0
	cloud.google.com/go/storage => cloud.google.com/go/storage v1.10.0
	github.com/beorn7/perks => github.com/beorn7/perks v1.0.1
	github.com/davecgh/go-spew => github.com/davecgh/go-spew v1.1.1
	github.com/form3tech-oss/jwt-go => github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/fsnotify/fsnotify => github.com/fsnotify/fsnotify v1.4.9
	github.com/golang/groupcache => github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/mock => github.com/golang/mock v1.5.0
	github.com/golang/protobuf => github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp => github.com/google/go-cmp v0.5.5
	github.com/google/gofuzz => github.com/google/gofuzz v1.1.0
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.5.5
	github.com/json-iterator/go => github.com/json-iterator/go v1.1.12
	github.com/mailru/easyjson => github.com/mailru/easyjson v0.7.6
	github.com/mattn/go-colorable => github.com/mattn/go-colorable v0.0.9
	github.com/mattn/go-isatty => github.com/mattn/go-isatty v0.0.3
	github.com/modern-go/concurrent => github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 => github.com/modern-go/reflect2 v1.0.2
	github.com/onsi/ginkgo => github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega => github.com/onsi/gomega v1.10.1
	github.com/prometheus/client_model => github.com/prometheus/client_model v0.2.0
	github.com/sirupsen/logrus => github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag => github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify => github.com/stretchr/testify v1.7.0
	go.opencensus.io => go.opencensus.io v0.23.0
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/exp => golang.org/x/exp v0.0.0-20210220032938-85be41e4509f
	golang.org/x/image => golang.org/x/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/lint => golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/mobile => golang.org/x/mobile v0.0.0-20201217150744-e6ae53a27f4f
	golang.org/x/mod => golang.org/x/mod v0.4.2
	golang.org/x/net => golang.org/x/net v0.0.0-20211209124913-491a49abca63
	golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sync => golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys => golang.org/x/sys v0.0.0-20210831042530-f4d43177bf5e
	golang.org/x/term => golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	golang.org/x/text => golang.org/x/text v0.3.7
	golang.org/x/time => golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	golang.org/x/tools => golang.org/x/tools v0.1.6-0.20210820212750-d4cc65f0b2ff
	golang.org/x/xerrors => golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/api => google.golang.org/api v0.46.0
	google.golang.org/appengine => google.golang.org/appengine v1.6.7
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20210831024726-fe130286e0e2
	google.golang.org/grpc => google.golang.org/grpc v1.40.0
	google.golang.org/protobuf => google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.1
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.30.0
	k8s.io/utils => k8s.io/utils v0.0.0-20211116205334-6203023598ed
	sigs.k8s.io/json => sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6
	sigs.k8s.io/structured-merge-diff/v4 => sigs.k8s.io/structured-merge-diff/v4 v4.2.1
)
