module github.com/guessi/kubectl-grep

go 1.19

require (
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.6.1
	k8s.io/api v0.25.8
	k8s.io/apimachinery v0.25.8
	k8s.io/client-go v0.25.8
)

require (
	cloud.google.com/go v0.97.0 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.27 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.20 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.8.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.14 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/imdario/mergo v0.3.6 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/term v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	cloud.google.com/go => cloud.google.com/go v0.97.0
	cloud.google.com/go/bigquery => cloud.google.com/go/bigquery v1.8.0
	cloud.google.com/go/storage => cloud.google.com/go/storage v1.10.0
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.2.0+incompatible
	github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.11.27
	github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.9.20
	github.com/Azure/go-autorest/autorest/date => github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/mocks => github.com/Azure/go-autorest/autorest/mocks v0.4.2
	github.com/Azure/go-autorest/autorest/to => github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/go-autorest/autorest/validation => github.com/Azure/go-autorest/autorest/validation v0.1.0
	github.com/Azure/go-autorest/logger => github.com/Azure/go-autorest/logger v0.2.1
	github.com/Azure/go-autorest/tracing => github.com/Azure/go-autorest/tracing v0.6.0
	github.com/PuerkitoBio/purell => github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc => github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
	github.com/davecgh/go-spew => github.com/davecgh/go-spew v1.1.1
	github.com/form3tech-oss/jwt-go => github.com/form3tech-oss/jwt-go v3.2.3+incompatible
	github.com/go-logr/logr => github.com/go-logr/logr v1.2.3
	github.com/go-openapi/jsonpointer => github.com/go-openapi/jsonpointer v0.19.5
	github.com/go-openapi/jsonreference => github.com/go-openapi/jsonreference v0.19.5
	github.com/go-openapi/swag => github.com/go-openapi/swag v0.19.14
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2
	github.com/golang-jwt/jwt/v4 => github.com/golang-jwt/jwt/v4 v4.2.0
	github.com/golang/protobuf => github.com/golang/protobuf v1.5.2
	github.com/google/gnostic => github.com/google/gnostic v0.5.7-v3refs
	github.com/google/go-cmp => github.com/google/go-cmp v0.5.8
	github.com/google/gofuzz => github.com/google/gofuzz v1.1.0
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.6
	github.com/inconshreveable/mousetrap => github.com/inconshreveable/mousetrap v1.0.0
	github.com/josharian/intern => github.com/josharian/intern v1.0.0
	github.com/json-iterator/go => github.com/json-iterator/go v1.1.12
	github.com/mailru/easyjson => github.com/mailru/easyjson v0.7.6
	github.com/modern-go/concurrent => github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 => github.com/modern-go/reflect2 v1.0.2
	github.com/munnerz/goautoneg => github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
	github.com/spf13/pflag => github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify => github.com/stretchr/testify v1.8.0
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	golang.org/x/mod => golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4
	golang.org/x/net => golang.org/x/net v0.7.0
	golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	golang.org/x/sync => golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	golang.org/x/sys => golang.org/x/sys v0.5.0
	golang.org/x/term => golang.org/x/term v0.5.0
	golang.org/x/text => golang.org/x/text v0.7.0
	golang.org/x/time => golang.org/x/time v0.0.0-20220210224613-90d013bbcef8
	golang.org/x/tools => golang.org/x/tools v0.1.12
	golang.org/x/xerrors => golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/api => google.golang.org/api v0.60.0
	google.golang.org/appengine => google.golang.org/appengine v1.6.7
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21
	google.golang.org/grpc => google.golang.org/grpc v1.47.0
	google.golang.org/protobuf => google.golang.org/protobuf v1.28.0
	gopkg.in/check.v1 => gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
	gopkg.in/inf.v0 => gopkg.in/inf.v0 v0.9.1
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.1
	honnef.co/go/tools => honnef.co/go/tools v0.0.1-2020.1.4
	k8s.io/api => k8s.io/api v0.25.8
	k8s.io/apimachinery => k8s.io/apimachinery v0.25.8
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.25.8
	k8s.io/client-go => k8s.io/client-go v0.25.8
	k8s.io/code-generator => k8s.io/code-generator v0.25.8
	k8s.io/component-base => k8s.io/component-base v0.25.8
	k8s.io/component-helpers => k8s.io/component-helpers v0.25.8
	k8s.io/klog/v2 => k8s.io/klog/v2 v2.70.1
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1
	k8s.io/metrics => k8s.io/metrics v0.25.8
	k8s.io/utils => k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed
	sigs.k8s.io/json => sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2
	sigs.k8s.io/structured-merge-diff/v4 => sigs.k8s.io/structured-merge-diff/v4 v4.2.3
	sigs.k8s.io/yaml => sigs.k8s.io/yaml v1.2.0
)
