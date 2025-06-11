package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	storagev1 "k8s.io/api/storage/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// CsiDriverList - return a list of CSIDriver(s)
func CsiDriverList(ctx context.Context, opt *options.SearchOptions) (*storagev1.CSIDriverList, error) {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.StorageV1().CSIDrivers().List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get CSIDriver List")
		return nil, fmt.Errorf("failed to list CSIDrivers: %w", err)
	}
	return list, nil
}

// StorageClassList - return a list of StorageClass(es)
func StorageClassList(ctx context.Context, opt *options.SearchOptions) (*storagev1.StorageClassList, error) {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.StorageV1().StorageClasses().List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get StorageClass List")
		return nil, fmt.Errorf("failed to list StorageClasses: %w", err)
	}
	return list, nil
}
