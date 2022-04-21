// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kubernetes:apps/v1beta1:ControllerRevision":
		r = &ControllerRevision{}
	case "kubernetes:apps/v1beta1:ControllerRevisionList":
		r = &ControllerRevisionList{}
	case "kubernetes:apps/v1beta1:Deployment":
		r = &Deployment{}
	case "kubernetes:apps/v1beta1:DeploymentList":
		r = &DeploymentList{}
	case "kubernetes:apps/v1beta1:StatefulSet":
		r = &StatefulSet{}
	case "kubernetes:apps/v1beta1:StatefulSetList":
		r = &StatefulSetList{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := kubernetes.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"kubernetes",
		"apps/v1beta1",
		&module{version},
	)
}
