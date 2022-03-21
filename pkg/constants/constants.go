package constants

const (
	// apps/v1
	DaemonsetHeader      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE"
	DaemonsetRowTemplate = "%s\t%s\t%d\t%d\t%d\t%d\t%s\t%s"

	DaemonsetHeaderWide      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tNODE SELECTOR\tAGE\tCONTAINERS\tIMAGES\tSELECTOR"
	DaemonsetRowTemplateWide = "%s\t%s\t%d\t%d\t%d\t%d\t%s\t%s\t%s\t%s\t%s"

	DeploymentHeader      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tAGE"
	DeploymentRowTemplate = "%s\t%s\t%d\t%d\t%d\t%d\t%s"

	DeploymentHeaderWide      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tUP-TO-DATE\tAVAILABLE\tAGE\tCONTAINERS\tIMAGES"
	DeploymentRowTemplateWide = "%s\t%s\t%d\t%d\t%d\t%d\t%s\t%s\t%s"

	ReplicasetHeader      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tAGE"
	ReplicasetRowTemplate = "%s\t%s\t%d\t%d\t%s"

	ReplicasetHeaderWide      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tAGE\tCONTAINERS\tIMAGES\tSELECTOR"
	ReplicasetRowTemplateWide = "%s\t%s\t%d\t%d\t%s\t%s\t%s\t%s"

	StatefulsetHeader      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tAGE"
	StatefulsetRowTemplate = "%s\t%s\t%d\t%d\t%s"

	StatefulsetHeaderWide      = "NAMESPACE\tNAME\tDESIRED\tCURRENT\tAGE\tCONTAINERS\tIMAGES"
	StatefulsetRowTemplateWide = "%s\t%s\t%d\t%d\t%s\t%s\t%s"

	// autoscaling/v1
	HpaHeader      = "NAMESPACE\tNAME\tREFERENCE\tTARGETS\tMINPODS\tMAXPODS\tREPLICAS\tAGE"
	HpaRowTemplate = "%s\t%s\t%s/%s\t%d%%/%d%%\t%d\t%d\t%d\t%s"

	// batch/v1
	JobsHeader      = "NAMESPACE\tNAME\tCOMPLETIONS\tDURATION\tAGE"
	JobsRowTemplate = "%s\t%s\t%d/%d\t%s\t%s"

	// networking.k8s.io/v1
	IngressHeader      = "NAMESPACE\tNAME\tCLASS\tHOSTS\tADDRESS\tPORTS\tAGE"
	IngressRowTemplate = "%s\t%s\t%s\t%s\t%s\t%s\t%s"

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

	PodHeader      = "NAMESPACE\tNAME\tREADY\tSTATUS\tRESTART\tAGE"
	PodRowTemplate = "%s\t%s\t%d/%d\t%s\t%d\t%s"

	PodHeaderWide      = "NAMESPACE\tNAME\tREADY\tSTATUS\tRESTART\tAGE\tIP\tNODENAME"
	PodRowTemplateWide = "%s\t%s\t%d/%d\t%s\t%d\t%s\t%s\t%s"

	SecretHeader      = "NAMESPACE\tNAME\tTYPE\tDATA\tAGE"
	SecretRowTemplate = "%s\t%s\t%s\t%d\t%s"

	ServicesHeaderWide      = "NAMESPACE\tNAME\tTYPE\tCLUSTER-IP\tEXTERNAL-IP\tPORT(S)\tAGE\tSELECTOR"
	ServicesRowTemplateWide = "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s"

	ServicesHeader      = "NAMESPACE\tNAME\tTYPE\tCLUSTER-IP\tEXTERNAL-IP\tPORT(S)\tAGE"
	ServicesRowTemplate = "%s\t%s\t%s\t%s\t%s\t%s\t%s"
)
