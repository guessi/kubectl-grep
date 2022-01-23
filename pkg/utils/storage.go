package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	storagev1 "k8s.io/api/storage/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// CsiDriverList - return a list of CSIDriverList(s)
func CsiDriverList(opt *options.SearchOptions) *storagev1.CSIDriverList {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.StorageV1().CSIDrivers().List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get CSIDriver List")
	}
	return list
}

// StorageClassList - return a list of StorageClassList(s)
func StorageClassList(opt *options.SearchOptions) *storagev1.StorageClassList {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.StorageV1().StorageClasses().List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get StorageClass List")
	}
	return list
}
