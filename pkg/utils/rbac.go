package utils

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	rbacv1 "k8s.io/api/rbac/v1"

	"github.com/guessi/kubectl-grep/pkg/client"
	"github.com/guessi/kubectl-grep/pkg/options"
)

// RoleList - return a list of Role(s)
func RoleList(ctx context.Context, opt *options.SearchOptions) (*rbacv1.RoleList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.RbacV1().Roles(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list Roles: %w", err)
	}
	return list, nil
}

// RoleBindingList - return a list of RoleBinding(s)
func RoleBindingList(ctx context.Context, opt *options.SearchOptions) (*rbacv1.RoleBindingList, error) {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.RbacV1().RoleBindings(ns).List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list RoleBindings: %w", err)
	}
	return list, nil
}

// ClusterRoleList - return a list of ClusterRole(s)
func ClusterRoleList(ctx context.Context, opt *options.SearchOptions) (*rbacv1.ClusterRoleList, error) {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.RbacV1().ClusterRoles().List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list ClusterRoles: %w", err)
	}
	return list, nil
}

// ClusterRoleBindingList - return a list of ClusterRoleBinding(s)
func ClusterRoleBindingList(ctx context.Context, opt *options.SearchOptions) (*rbacv1.ClusterRoleBindingList, error) {
	clientset := client.InitClient()
	_, o := setOptions(opt)
	list, err := clientset.RbacV1().ClusterRoleBindings().List(ctx, *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		})
		return nil, fmt.Errorf("failed to list ClusterRoleBindings: %w", err)
	}
	return list, nil
}
