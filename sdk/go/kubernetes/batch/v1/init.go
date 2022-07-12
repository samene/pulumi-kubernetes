// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

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
	case "kubernetes:batch/v1:CronJob":
		r = &CronJob{}
	case "kubernetes:batch/v1:CronJobList":
		r = &CronJobList{}
	case "kubernetes:batch/v1:CronJobPatch":
		r = &CronJobPatch{}
	case "kubernetes:batch/v1:Job":
		r = &Job{}
	case "kubernetes:batch/v1:JobList":
		r = &JobList{}
	case "kubernetes:batch/v1:JobPatch":
		r = &JobPatch{}
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
		"batch/v1",
		&module{version},
	)
}
