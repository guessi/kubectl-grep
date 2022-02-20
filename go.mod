module github.com/guessi/kubectl-grep

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	k8s.io/api v0.22.7
	k8s.io/apimachinery v0.22.7
	k8s.io/client-go v0.22.7
)

replace (
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
)
