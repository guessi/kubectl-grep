module github.com/guessi/kubectl-grep

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	k8s.io/api v0.22.8
	k8s.io/apimachinery v0.22.8
	k8s.io/client-go v0.22.8
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.54.0
	github.com/beorn7/perks => github.com/beorn7/perks v1.0.1
	github.com/golang/groupcache => github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/golang/mock => github.com/golang/mock v1.4.4
	github.com/golang/protobuf => github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp => github.com/google/go-cmp v0.5.5
	github.com/json-iterator/go => github.com/json-iterator/go v1.1.11
	github.com/mailru/easyjson => github.com/mailru/easyjson v0.7.6
	github.com/mattn/go-colorable => github.com/mattn/go-colorable v0.0.9
	github.com/mattn/go-isatty => github.com/mattn/go-isatty v0.0.3
	github.com/modern-go/reflect2 => github.com/modern-go/reflect2 v1.0.1
	github.com/onsi/ginkgo => github.com/onsi/ginkgo v1.14.0
	github.com/prometheus/client_model => github.com/prometheus/client_model v0.2.0
	github.com/sirupsen/logrus => github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag => github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify => github.com/stretchr/testify v1.7.0
	go.opencensus.io => go.opencensus.io v0.22.3
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/exp => golang.org/x/exp v0.0.0-20210220032938-85be41e4509f
	golang.org/x/image => golang.org/x/image v0.0.0-20190802002840-cff245a6509b
	golang.org/x/lint => golang.org/x/lint v0.0.0-20210508222113-6edffad5e616
	golang.org/x/mobile => golang.org/x/mobile v0.0.0-20201217150744-e6ae53a27f4f
	golang.org/x/mod => golang.org/x/mod v0.4.2
	golang.org/x/net => golang.org/x/net v0.0.0-20211209124913-491a49abca63
	golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync => golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys => golang.org/x/sys v0.0.0-20210616094352-59db8d763f22
	golang.org/x/term => golang.org/x/term v0.0.0-20210220032956-6a3ed077a48d
	golang.org/x/text => golang.org/x/text v0.3.6
	golang.org/x/time => golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	golang.org/x/tools => golang.org/x/tools v0.1.2
	golang.org/x/xerrors => golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/api => google.golang.org/api v0.20.0
	google.golang.org/appengine => google.golang.org/appengine v1.6.5
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20210602131652-f16073e35f0c
	google.golang.org/grpc => google.golang.org/grpc v1.38.0
	google.golang.org/protobuf => google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.9.0
	sigs.k8s.io/structured-merge-diff/v4 => sigs.k8s.io/structured-merge-diff/v4 v4.2.1
)
