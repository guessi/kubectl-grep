package constants

const (
	// common
	UNKNOWN = "<unknown>"

	// apps/v1
	DaemonsetHeader      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tREADY\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE"
	DaemonsetRowTemplate = "%s\t%s\t%d\t%d\t%d\t%d\t%d\t%s\t%s"

	DaemonsetHeaderWide      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tREADY\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE\tCONTAINERS\tIMAGES\tSELECTOR"
	DaemonsetRowTemplateWide = "%s\t%s\t%d\t%d\t%d\t%d\t%d\t%s\t%s\t%s\t%s\t%s"

	DeploymentHeader      = "NAMESPACE\tNAME\tREADY\tUP-TO-DATE\tAVAILABLE\tAGE"
	DeploymentRowTemplate = "%s\t%s\t%d/%d\t%d\t%d\t%s"

	DeploymentHeaderWide      = "NAMESPACE\tNAME\tREADY\tUP-TO-DATE\tAVAILABLE\tAGE\tCONTAINERS\tIMAGES\tSELECTOR"
	DeploymentRowTemplateWide = "%s\t%s\t%d/%d\t%d\t%d\t%s\t%s\t%s\t%s"

	ReplicasetHeader      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tREADY\tAGE"
	ReplicasetRowTemplate = "%s\t%s\t%d\t%d\t%d\t%s"

	ReplicasetHeaderWide      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tREADY\tAGE\tCONTAINERS\tIMAGES\tSELECTOR"
	ReplicasetRowTemplateWide = "%s\t%s\t%d\t%d\t%d\t%s\t%s\t%s\t%s"

	StatefulsetHeader      = "NAMESPACE\tNAME\tREADY\tAGE"
	StatefulsetRowTemplate = "%s\t%s\t%d/%d\t%s"

	StatefulsetHeaderWide      = "NAMESPACE\tNAME\tREADY\tAGE\tCONTAINERS\tIMAGES"
	StatefulsetRowTemplateWide = "%s\t%s\t%d/%d\t%s\t%s\t%s"

	// autoscaling/v1
	HpaHeader      = "NAMESPACE\tNAME\tREFERENCE\tTARGETS\tMINPODS\tMAXPODS\tREPLICAS\tAGE"
	HpaRowTemplate = "%s\t%s\t%s/%s\t%s\t%d\t%d\t%d\t%s"

	// batch/v1
	CronJobsHeader      = "NAMESPACE\tNAME\tSCHEDULE\tSUSPEND\tACTIVE\tLAST SCHEDULE\tAGE"
	CronJobsRowTemplate = "%s\t%s\t%s\t%s\t%d\t%s\t%s"

	JobsHeader      = "NAMESPACE\tNAME\tCOMPLETIONS\tDURATION\tAGE"
	JobsRowTemplate = "%s\t%s\t%d/%d\t%s\t%s"

	// networking.k8s.io/v1
	IngressHeader      = "NAMESPACE\tNAME\tCLASS\tHOSTS\tADDRESS\tPORTS\tAGE"
	IngressRowTemplate = "%s\t%s\t%s\t%s\t%s\t%s\t%s"

	// rbac.authorization.k8s.io/v1
	RolesHeader      = "NAMESPACE\tNAME\tCREATED AT"
	RolesRowTemplate = "%s\t%s\t%s"

	RoleBindingsHeader      = "NAMESPACE\tNAME\tROLE\tAGE"
	RoleBindingsRowTemplate = "%s\t%s\t%s\t%s"

	ClusterRolesHeader      = "NAME\tCREATED AT"
	ClusterRolesRowTemplate = "%s\t%s"

	ClusterRoleBindingsHeader      = "NAME\tROLE\tAGE"
	ClusterRoleBindingsRowTemplate = "%s\t%s\t%s"

	// storage.k8s.io/v1
	CsiDriversHeader      = "NAME\tATTACHREQUIRED\tPODINFOONMOUNT\tSTORAGECAPACITY\tTOKENREQUESTS\tREQUIRESREPUBLISH\tMODES\tAGE"
	CsiDriversRowTemplate = "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s"

	StorageClassesRowTemplate = "%s %s\t%s\t%s\t%s\t%s\t%s"
	StorageClassesHeader      = "NAME\tPROVISIONER\tRECLAIMPOLICY\tVOLUMEBINDINGMODE\tALLOWVOLUMEEXPANSION\tAGE"

	// v1
	ConfigMapHeader      = "NAMESPACE\tNAME\tDATA\tAGE"
	ConfigMapRowTemplate = "%s\t%s\t%d\t%s"

	NodeHeader      = "NAME\tSTATUS\tROLES\tAGE\tVERSION"
	NodeRowTemplate = "%s\t%s\t%s\t%s\t%s"

	NodeHeaderWide      = "NAME\tSTATUS\tROLES\tAGE\tVERSION\tINTERNAL-IP\tEXTERNAL-IP\tOS-IMAGE\tKERNEL-VERSION\tCONTAINER-RUNTIME"
	NodeRowTemplateWide = "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s"

	PodHeader      = "NAMESPACE\tNAME\tREADY\tSTATUS\tRESTARTS\tAGE"
	PodRowTemplate = "%s\t%s\t%d/%d\t%s\t%d\t%s"

	PodHeaderWide      = "NAMESPACE\tNAME\tREADY\tSTATUS\tRESTARTS\tAGE\tIP\tNODENAME"
	PodRowTemplateWide = "%s\t%s\t%d/%d\t%s\t%d\t%s\t%s\t%s"

	SecretHeader      = "NAMESPACE\tNAME\tTYPE\tDATA\tAGE"
	SecretRowTemplate = "%s\t%s\t%s\t%d\t%s"

	ServiceAccountsHeader      = "NAMESPACE\tNAME\tSECRETS\tAGE"
	ServiceAccountsRowTemplate = "%s\t%s\t%d\t%s"

	ServicesHeaderWide      = "NAMESPACE\tNAME\tTYPE\tCLUSTER-IP\tEXTERNAL-IP\tPORT(S)\tAGE\tSELECTOR"
	ServicesRowTemplateWide = "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s"

	ServicesHeader      = "NAMESPACE\tNAME\tTYPE\tCLUSTER-IP\tEXTERNAL-IP\tPORT(S)\tAGE"
	ServicesRowTemplate = "%s\t%s\t%s\t%s\t%s\t%s\t%s"
)
