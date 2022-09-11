package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// RoleList - return a list of Role(s)
func RoleList(opt *options.SearchOptions) *rbacv1.RoleList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.RbacV1().Roles(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Role List")
	}
	return list
}

// RoleBindingList - return a list of RoleBinding(s)
func RoleBindingList(opt *options.SearchOptions) *rbacv1.RoleBindingList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.RbacV1().RoleBindings(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get RoleBinding List")
	}
	return list
}

// ClusterRoleList - return a list of ClusterRole(s)
func ClusterRoleList(opt *options.SearchOptions) *rbacv1.ClusterRoleList {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get ClusterRole List")
	}
	return list
}

// ClusterRoleBindingList - return a list of ClusterRoleBinding(s)
func ClusterRoleBindingList(opt *options.SearchOptions) *rbacv1.ClusterRoleBindingList {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get ClusterRoleBinding List")
	}
	return list
}
