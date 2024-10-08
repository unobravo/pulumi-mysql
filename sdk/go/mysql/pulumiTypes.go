// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unobravo/pulumi-mysql/sdk/go/mysql/internal"
)

var _ = internal.GetEnvOrDefault

type ProviderAwsConfig struct {
	AccessKey *string `pulumi:"accessKey"`
	Profile   *string `pulumi:"profile"`
	Region    *string `pulumi:"region"`
	SecretKey *string `pulumi:"secretKey"`
}

// ProviderAwsConfigInput is an input type that accepts ProviderAwsConfigArgs and ProviderAwsConfigOutput values.
// You can construct a concrete instance of `ProviderAwsConfigInput` via:
//
//	ProviderAwsConfigArgs{...}
type ProviderAwsConfigInput interface {
	pulumi.Input

	ToProviderAwsConfigOutput() ProviderAwsConfigOutput
	ToProviderAwsConfigOutputWithContext(context.Context) ProviderAwsConfigOutput
}

type ProviderAwsConfigArgs struct {
	AccessKey pulumi.StringPtrInput `pulumi:"accessKey"`
	Profile   pulumi.StringPtrInput `pulumi:"profile"`
	Region    pulumi.StringPtrInput `pulumi:"region"`
	SecretKey pulumi.StringPtrInput `pulumi:"secretKey"`
}

func (ProviderAwsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAwsConfig)(nil)).Elem()
}

func (i ProviderAwsConfigArgs) ToProviderAwsConfigOutput() ProviderAwsConfigOutput {
	return i.ToProviderAwsConfigOutputWithContext(context.Background())
}

func (i ProviderAwsConfigArgs) ToProviderAwsConfigOutputWithContext(ctx context.Context) ProviderAwsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsConfigOutput)
}

// ProviderAwsConfigArrayInput is an input type that accepts ProviderAwsConfigArray and ProviderAwsConfigArrayOutput values.
// You can construct a concrete instance of `ProviderAwsConfigArrayInput` via:
//
//	ProviderAwsConfigArray{ ProviderAwsConfigArgs{...} }
type ProviderAwsConfigArrayInput interface {
	pulumi.Input

	ToProviderAwsConfigArrayOutput() ProviderAwsConfigArrayOutput
	ToProviderAwsConfigArrayOutputWithContext(context.Context) ProviderAwsConfigArrayOutput
}

type ProviderAwsConfigArray []ProviderAwsConfigInput

func (ProviderAwsConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderAwsConfig)(nil)).Elem()
}

func (i ProviderAwsConfigArray) ToProviderAwsConfigArrayOutput() ProviderAwsConfigArrayOutput {
	return i.ToProviderAwsConfigArrayOutputWithContext(context.Background())
}

func (i ProviderAwsConfigArray) ToProviderAwsConfigArrayOutputWithContext(ctx context.Context) ProviderAwsConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAwsConfigArrayOutput)
}

type ProviderAwsConfigOutput struct{ *pulumi.OutputState }

func (ProviderAwsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAwsConfig)(nil)).Elem()
}

func (o ProviderAwsConfigOutput) ToProviderAwsConfigOutput() ProviderAwsConfigOutput {
	return o
}

func (o ProviderAwsConfigOutput) ToProviderAwsConfigOutputWithContext(ctx context.Context) ProviderAwsConfigOutput {
	return o
}

func (o ProviderAwsConfigOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAwsConfig) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o ProviderAwsConfigOutput) Profile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAwsConfig) *string { return v.Profile }).(pulumi.StringPtrOutput)
}

func (o ProviderAwsConfigOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAwsConfig) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o ProviderAwsConfigOutput) SecretKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAwsConfig) *string { return v.SecretKey }).(pulumi.StringPtrOutput)
}

type ProviderAwsConfigArrayOutput struct{ *pulumi.OutputState }

func (ProviderAwsConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderAwsConfig)(nil)).Elem()
}

func (o ProviderAwsConfigArrayOutput) ToProviderAwsConfigArrayOutput() ProviderAwsConfigArrayOutput {
	return o
}

func (o ProviderAwsConfigArrayOutput) ToProviderAwsConfigArrayOutputWithContext(ctx context.Context) ProviderAwsConfigArrayOutput {
	return o
}

func (o ProviderAwsConfigArrayOutput) Index(i pulumi.IntInput) ProviderAwsConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderAwsConfig {
		return vs[0].([]ProviderAwsConfig)[vs[1].(int)]
	}).(ProviderAwsConfigOutput)
}

type ProviderAzureConfig struct {
	ClientId     *string `pulumi:"clientId"`
	ClientSecret *string `pulumi:"clientSecret"`
	Environment  *string `pulumi:"environment"`
	TenantId     *string `pulumi:"tenantId"`
}

// ProviderAzureConfigInput is an input type that accepts ProviderAzureConfigArgs and ProviderAzureConfigOutput values.
// You can construct a concrete instance of `ProviderAzureConfigInput` via:
//
//	ProviderAzureConfigArgs{...}
type ProviderAzureConfigInput interface {
	pulumi.Input

	ToProviderAzureConfigOutput() ProviderAzureConfigOutput
	ToProviderAzureConfigOutputWithContext(context.Context) ProviderAzureConfigOutput
}

type ProviderAzureConfigArgs struct {
	ClientId     pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret pulumi.StringPtrInput `pulumi:"clientSecret"`
	Environment  pulumi.StringPtrInput `pulumi:"environment"`
	TenantId     pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (ProviderAzureConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAzureConfig)(nil)).Elem()
}

func (i ProviderAzureConfigArgs) ToProviderAzureConfigOutput() ProviderAzureConfigOutput {
	return i.ToProviderAzureConfigOutputWithContext(context.Background())
}

func (i ProviderAzureConfigArgs) ToProviderAzureConfigOutputWithContext(ctx context.Context) ProviderAzureConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAzureConfigOutput)
}

// ProviderAzureConfigArrayInput is an input type that accepts ProviderAzureConfigArray and ProviderAzureConfigArrayOutput values.
// You can construct a concrete instance of `ProviderAzureConfigArrayInput` via:
//
//	ProviderAzureConfigArray{ ProviderAzureConfigArgs{...} }
type ProviderAzureConfigArrayInput interface {
	pulumi.Input

	ToProviderAzureConfigArrayOutput() ProviderAzureConfigArrayOutput
	ToProviderAzureConfigArrayOutputWithContext(context.Context) ProviderAzureConfigArrayOutput
}

type ProviderAzureConfigArray []ProviderAzureConfigInput

func (ProviderAzureConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderAzureConfig)(nil)).Elem()
}

func (i ProviderAzureConfigArray) ToProviderAzureConfigArrayOutput() ProviderAzureConfigArrayOutput {
	return i.ToProviderAzureConfigArrayOutputWithContext(context.Background())
}

func (i ProviderAzureConfigArray) ToProviderAzureConfigArrayOutputWithContext(ctx context.Context) ProviderAzureConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderAzureConfigArrayOutput)
}

type ProviderAzureConfigOutput struct{ *pulumi.OutputState }

func (ProviderAzureConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderAzureConfig)(nil)).Elem()
}

func (o ProviderAzureConfigOutput) ToProviderAzureConfigOutput() ProviderAzureConfigOutput {
	return o
}

func (o ProviderAzureConfigOutput) ToProviderAzureConfigOutputWithContext(ctx context.Context) ProviderAzureConfigOutput {
	return o
}

func (o ProviderAzureConfigOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAzureConfig) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ProviderAzureConfigOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAzureConfig) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ProviderAzureConfigOutput) Environment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAzureConfig) *string { return v.Environment }).(pulumi.StringPtrOutput)
}

func (o ProviderAzureConfigOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderAzureConfig) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ProviderAzureConfigArrayOutput struct{ *pulumi.OutputState }

func (ProviderAzureConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderAzureConfig)(nil)).Elem()
}

func (o ProviderAzureConfigArrayOutput) ToProviderAzureConfigArrayOutput() ProviderAzureConfigArrayOutput {
	return o
}

func (o ProviderAzureConfigArrayOutput) ToProviderAzureConfigArrayOutputWithContext(ctx context.Context) ProviderAzureConfigArrayOutput {
	return o
}

func (o ProviderAzureConfigArrayOutput) Index(i pulumi.IntInput) ProviderAzureConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderAzureConfig {
		return vs[0].([]ProviderAzureConfig)[vs[1].(int)]
	}).(ProviderAzureConfigOutput)
}

type ProviderCustomTl struct {
	CaCert     string  `pulumi:"caCert"`
	ClientCert string  `pulumi:"clientCert"`
	ClientKey  string  `pulumi:"clientKey"`
	ConfigKey  *string `pulumi:"configKey"`
}

// ProviderCustomTlInput is an input type that accepts ProviderCustomTlArgs and ProviderCustomTlOutput values.
// You can construct a concrete instance of `ProviderCustomTlInput` via:
//
//	ProviderCustomTlArgs{...}
type ProviderCustomTlInput interface {
	pulumi.Input

	ToProviderCustomTlOutput() ProviderCustomTlOutput
	ToProviderCustomTlOutputWithContext(context.Context) ProviderCustomTlOutput
}

type ProviderCustomTlArgs struct {
	CaCert     pulumi.StringInput    `pulumi:"caCert"`
	ClientCert pulumi.StringInput    `pulumi:"clientCert"`
	ClientKey  pulumi.StringInput    `pulumi:"clientKey"`
	ConfigKey  pulumi.StringPtrInput `pulumi:"configKey"`
}

func (ProviderCustomTlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderCustomTl)(nil)).Elem()
}

func (i ProviderCustomTlArgs) ToProviderCustomTlOutput() ProviderCustomTlOutput {
	return i.ToProviderCustomTlOutputWithContext(context.Background())
}

func (i ProviderCustomTlArgs) ToProviderCustomTlOutputWithContext(ctx context.Context) ProviderCustomTlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderCustomTlOutput)
}

// ProviderCustomTlArrayInput is an input type that accepts ProviderCustomTlArray and ProviderCustomTlArrayOutput values.
// You can construct a concrete instance of `ProviderCustomTlArrayInput` via:
//
//	ProviderCustomTlArray{ ProviderCustomTlArgs{...} }
type ProviderCustomTlArrayInput interface {
	pulumi.Input

	ToProviderCustomTlArrayOutput() ProviderCustomTlArrayOutput
	ToProviderCustomTlArrayOutputWithContext(context.Context) ProviderCustomTlArrayOutput
}

type ProviderCustomTlArray []ProviderCustomTlInput

func (ProviderCustomTlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderCustomTl)(nil)).Elem()
}

func (i ProviderCustomTlArray) ToProviderCustomTlArrayOutput() ProviderCustomTlArrayOutput {
	return i.ToProviderCustomTlArrayOutputWithContext(context.Background())
}

func (i ProviderCustomTlArray) ToProviderCustomTlArrayOutputWithContext(ctx context.Context) ProviderCustomTlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderCustomTlArrayOutput)
}

type ProviderCustomTlOutput struct{ *pulumi.OutputState }

func (ProviderCustomTlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderCustomTl)(nil)).Elem()
}

func (o ProviderCustomTlOutput) ToProviderCustomTlOutput() ProviderCustomTlOutput {
	return o
}

func (o ProviderCustomTlOutput) ToProviderCustomTlOutputWithContext(ctx context.Context) ProviderCustomTlOutput {
	return o
}

func (o ProviderCustomTlOutput) CaCert() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderCustomTl) string { return v.CaCert }).(pulumi.StringOutput)
}

func (o ProviderCustomTlOutput) ClientCert() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderCustomTl) string { return v.ClientCert }).(pulumi.StringOutput)
}

func (o ProviderCustomTlOutput) ClientKey() pulumi.StringOutput {
	return o.ApplyT(func(v ProviderCustomTl) string { return v.ClientKey }).(pulumi.StringOutput)
}

func (o ProviderCustomTlOutput) ConfigKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProviderCustomTl) *string { return v.ConfigKey }).(pulumi.StringPtrOutput)
}

type ProviderCustomTlArrayOutput struct{ *pulumi.OutputState }

func (ProviderCustomTlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProviderCustomTl)(nil)).Elem()
}

func (o ProviderCustomTlArrayOutput) ToProviderCustomTlArrayOutput() ProviderCustomTlArrayOutput {
	return o
}

func (o ProviderCustomTlArrayOutput) ToProviderCustomTlArrayOutputWithContext(ctx context.Context) ProviderCustomTlArrayOutput {
	return o
}

func (o ProviderCustomTlArrayOutput) Index(i pulumi.IntInput) ProviderCustomTlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProviderCustomTl {
		return vs[0].([]ProviderCustomTl)[vs[1].(int)]
	}).(ProviderCustomTlOutput)
}

type UserAadIdentity struct {
	Identity string  `pulumi:"identity"`
	Type     *string `pulumi:"type"`
}

// UserAadIdentityInput is an input type that accepts UserAadIdentityArgs and UserAadIdentityOutput values.
// You can construct a concrete instance of `UserAadIdentityInput` via:
//
//	UserAadIdentityArgs{...}
type UserAadIdentityInput interface {
	pulumi.Input

	ToUserAadIdentityOutput() UserAadIdentityOutput
	ToUserAadIdentityOutputWithContext(context.Context) UserAadIdentityOutput
}

type UserAadIdentityArgs struct {
	Identity pulumi.StringInput    `pulumi:"identity"`
	Type     pulumi.StringPtrInput `pulumi:"type"`
}

func (UserAadIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAadIdentity)(nil)).Elem()
}

func (i UserAadIdentityArgs) ToUserAadIdentityOutput() UserAadIdentityOutput {
	return i.ToUserAadIdentityOutputWithContext(context.Background())
}

func (i UserAadIdentityArgs) ToUserAadIdentityOutputWithContext(ctx context.Context) UserAadIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAadIdentityOutput)
}

func (i UserAadIdentityArgs) ToUserAadIdentityPtrOutput() UserAadIdentityPtrOutput {
	return i.ToUserAadIdentityPtrOutputWithContext(context.Background())
}

func (i UserAadIdentityArgs) ToUserAadIdentityPtrOutputWithContext(ctx context.Context) UserAadIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAadIdentityOutput).ToUserAadIdentityPtrOutputWithContext(ctx)
}

// UserAadIdentityPtrInput is an input type that accepts UserAadIdentityArgs, UserAadIdentityPtr and UserAadIdentityPtrOutput values.
// You can construct a concrete instance of `UserAadIdentityPtrInput` via:
//
//	        UserAadIdentityArgs{...}
//
//	or:
//
//	        nil
type UserAadIdentityPtrInput interface {
	pulumi.Input

	ToUserAadIdentityPtrOutput() UserAadIdentityPtrOutput
	ToUserAadIdentityPtrOutputWithContext(context.Context) UserAadIdentityPtrOutput
}

type userAadIdentityPtrType UserAadIdentityArgs

func UserAadIdentityPtr(v *UserAadIdentityArgs) UserAadIdentityPtrInput {
	return (*userAadIdentityPtrType)(v)
}

func (*userAadIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAadIdentity)(nil)).Elem()
}

func (i *userAadIdentityPtrType) ToUserAadIdentityPtrOutput() UserAadIdentityPtrOutput {
	return i.ToUserAadIdentityPtrOutputWithContext(context.Background())
}

func (i *userAadIdentityPtrType) ToUserAadIdentityPtrOutputWithContext(ctx context.Context) UserAadIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAadIdentityPtrOutput)
}

type UserAadIdentityOutput struct{ *pulumi.OutputState }

func (UserAadIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAadIdentity)(nil)).Elem()
}

func (o UserAadIdentityOutput) ToUserAadIdentityOutput() UserAadIdentityOutput {
	return o
}

func (o UserAadIdentityOutput) ToUserAadIdentityOutputWithContext(ctx context.Context) UserAadIdentityOutput {
	return o
}

func (o UserAadIdentityOutput) ToUserAadIdentityPtrOutput() UserAadIdentityPtrOutput {
	return o.ToUserAadIdentityPtrOutputWithContext(context.Background())
}

func (o UserAadIdentityOutput) ToUserAadIdentityPtrOutputWithContext(ctx context.Context) UserAadIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserAadIdentity) *UserAadIdentity {
		return &v
	}).(UserAadIdentityPtrOutput)
}

func (o UserAadIdentityOutput) Identity() pulumi.StringOutput {
	return o.ApplyT(func(v UserAadIdentity) string { return v.Identity }).(pulumi.StringOutput)
}

func (o UserAadIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserAadIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type UserAadIdentityPtrOutput struct{ *pulumi.OutputState }

func (UserAadIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserAadIdentity)(nil)).Elem()
}

func (o UserAadIdentityPtrOutput) ToUserAadIdentityPtrOutput() UserAadIdentityPtrOutput {
	return o
}

func (o UserAadIdentityPtrOutput) ToUserAadIdentityPtrOutputWithContext(ctx context.Context) UserAadIdentityPtrOutput {
	return o
}

func (o UserAadIdentityPtrOutput) Elem() UserAadIdentityOutput {
	return o.ApplyT(func(v *UserAadIdentity) UserAadIdentity {
		if v != nil {
			return *v
		}
		var ret UserAadIdentity
		return ret
	}).(UserAadIdentityOutput)
}

func (o UserAadIdentityPtrOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAadIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Identity
	}).(pulumi.StringPtrOutput)
}

func (o UserAadIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserAadIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsConfigInput)(nil)).Elem(), ProviderAwsConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAwsConfigArrayInput)(nil)).Elem(), ProviderAwsConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAzureConfigInput)(nil)).Elem(), ProviderAzureConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderAzureConfigArrayInput)(nil)).Elem(), ProviderAzureConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderCustomTlInput)(nil)).Elem(), ProviderCustomTlArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderCustomTlArrayInput)(nil)).Elem(), ProviderCustomTlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAadIdentityInput)(nil)).Elem(), UserAadIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserAadIdentityPtrInput)(nil)).Elem(), UserAadIdentityArgs{})
	pulumi.RegisterOutputType(ProviderAwsConfigOutput{})
	pulumi.RegisterOutputType(ProviderAwsConfigArrayOutput{})
	pulumi.RegisterOutputType(ProviderAzureConfigOutput{})
	pulumi.RegisterOutputType(ProviderAzureConfigArrayOutput{})
	pulumi.RegisterOutputType(ProviderCustomTlOutput{})
	pulumi.RegisterOutputType(ProviderCustomTlArrayOutput{})
	pulumi.RegisterOutputType(UserAadIdentityOutput{})
	pulumi.RegisterOutputType(UserAadIdentityPtrOutput{})
}
