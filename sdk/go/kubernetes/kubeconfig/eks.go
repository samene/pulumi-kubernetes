// Code generated by pulumigen DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubeconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Generate a kubeconfig for cluster authentication that does not use the default AWS credential provider chain, and instead is scoped to the supported options in `KubeconfigOptions`.
//
// The kubeconfig generated is automatically stringified for ease of use with the pulumi/kubernetes provider.
//
// See for more details:
// - https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-role.html
// - https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html
func Eks(ctx *pulumi.Context, args *EksArgs, opts ...pulumi.InvokeOption) (*EksResult, error) {
	var rv EksResult
	err := ctx.Invoke("kubernetes:kubeconfig:eks", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type EksArgs struct {
	// AWS credential profile name to always use instead of the default AWS credential provider chain.
	//
	// The profile is passed to kubeconfig as an authentication environment setting.
	ProfileName *string `pulumi:"profileName"`
	// Role ARN to assume instead of the default AWS credential provider chain.
	//
	// The role is passed to kubeconfig as an authentication exec argument.
	RoleArn *string `pulumi:"roleArn"`
}

type EksResult struct {
	Result string `pulumi:"result"`
}

func EksOutput(ctx *pulumi.Context, args EksOutputArgs, opts ...pulumi.InvokeOption) EksResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (EksResult, error) {
			args := v.(EksArgs)
			r, err := Eks(ctx, &args, opts...)
			var s EksResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(EksResultOutput)
}

type EksOutputArgs struct {
	// AWS credential profile name to always use instead of the default AWS credential provider chain.
	//
	// The profile is passed to kubeconfig as an authentication environment setting.
	ProfileName pulumi.StringPtrInput `pulumi:"profileName"`
	// Role ARN to assume instead of the default AWS credential provider chain.
	//
	// The role is passed to kubeconfig as an authentication exec argument.
	RoleArn pulumi.StringPtrInput `pulumi:"roleArn"`
}

func (EksOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EksArgs)(nil)).Elem()
}

type EksResultOutput struct{ *pulumi.OutputState }

func (EksResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EksResult)(nil)).Elem()
}

func (o EksResultOutput) ToEksResultOutput() EksResultOutput {
	return o
}

func (o EksResultOutput) ToEksResultOutputWithContext(ctx context.Context) EksResultOutput {
	return o
}

func (o EksResultOutput) Result() pulumi.StringOutput {
	return o.ApplyT(func(v EksResult) string { return v.Result }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EksResultOutput{})
}
