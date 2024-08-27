// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unobravo/pulumi-mysql/sdk/go/mysql/internal"
)

type UserPassword struct {
	pulumi.CustomResourceState

	// The source host of the user. Defaults to `localhost`.
	Host              pulumi.StringPtrOutput `pulumi:"host"`
	PlaintextPassword pulumi.StringPtrOutput `pulumi:"plaintextPassword"`
	RetainOldPassword pulumi.BoolPtrOutput   `pulumi:"retainOldPassword"`
	// The IAM user to associate with this access key.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewUserPassword registers a new resource with the given unique name, arguments, and options.
func NewUserPassword(ctx *pulumi.Context,
	name string, args *UserPasswordArgs, opts ...pulumi.ResourceOption) (*UserPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.User == nil {
		return nil, errors.New("invalid value for required argument 'User'")
	}
	if args.PlaintextPassword != nil {
		args.PlaintextPassword = pulumi.ToSecret(args.PlaintextPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"plaintextPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserPassword
	err := ctx.RegisterResource("mysql:index/userPassword:UserPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPassword gets an existing UserPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPasswordState, opts ...pulumi.ResourceOption) (*UserPassword, error) {
	var resource UserPassword
	err := ctx.ReadResource("mysql:index/userPassword:UserPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPassword resources.
type userPasswordState struct {
	// The source host of the user. Defaults to `localhost`.
	Host              *string `pulumi:"host"`
	PlaintextPassword *string `pulumi:"plaintextPassword"`
	RetainOldPassword *bool   `pulumi:"retainOldPassword"`
	// The IAM user to associate with this access key.
	User *string `pulumi:"user"`
}

type UserPasswordState struct {
	// The source host of the user. Defaults to `localhost`.
	Host              pulumi.StringPtrInput
	PlaintextPassword pulumi.StringPtrInput
	RetainOldPassword pulumi.BoolPtrInput
	// The IAM user to associate with this access key.
	User pulumi.StringPtrInput
}

func (UserPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPasswordState)(nil)).Elem()
}

type userPasswordArgs struct {
	// The source host of the user. Defaults to `localhost`.
	Host              *string `pulumi:"host"`
	PlaintextPassword *string `pulumi:"plaintextPassword"`
	RetainOldPassword *bool   `pulumi:"retainOldPassword"`
	// The IAM user to associate with this access key.
	User string `pulumi:"user"`
}

// The set of arguments for constructing a UserPassword resource.
type UserPasswordArgs struct {
	// The source host of the user. Defaults to `localhost`.
	Host              pulumi.StringPtrInput
	PlaintextPassword pulumi.StringPtrInput
	RetainOldPassword pulumi.BoolPtrInput
	// The IAM user to associate with this access key.
	User pulumi.StringInput
}

func (UserPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPasswordArgs)(nil)).Elem()
}

type UserPasswordInput interface {
	pulumi.Input

	ToUserPasswordOutput() UserPasswordOutput
	ToUserPasswordOutputWithContext(ctx context.Context) UserPasswordOutput
}

func (*UserPassword) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPassword)(nil)).Elem()
}

func (i *UserPassword) ToUserPasswordOutput() UserPasswordOutput {
	return i.ToUserPasswordOutputWithContext(context.Background())
}

func (i *UserPassword) ToUserPasswordOutputWithContext(ctx context.Context) UserPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPasswordOutput)
}

// UserPasswordArrayInput is an input type that accepts UserPasswordArray and UserPasswordArrayOutput values.
// You can construct a concrete instance of `UserPasswordArrayInput` via:
//
//	UserPasswordArray{ UserPasswordArgs{...} }
type UserPasswordArrayInput interface {
	pulumi.Input

	ToUserPasswordArrayOutput() UserPasswordArrayOutput
	ToUserPasswordArrayOutputWithContext(context.Context) UserPasswordArrayOutput
}

type UserPasswordArray []UserPasswordInput

func (UserPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPassword)(nil)).Elem()
}

func (i UserPasswordArray) ToUserPasswordArrayOutput() UserPasswordArrayOutput {
	return i.ToUserPasswordArrayOutputWithContext(context.Background())
}

func (i UserPasswordArray) ToUserPasswordArrayOutputWithContext(ctx context.Context) UserPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPasswordArrayOutput)
}

// UserPasswordMapInput is an input type that accepts UserPasswordMap and UserPasswordMapOutput values.
// You can construct a concrete instance of `UserPasswordMapInput` via:
//
//	UserPasswordMap{ "key": UserPasswordArgs{...} }
type UserPasswordMapInput interface {
	pulumi.Input

	ToUserPasswordMapOutput() UserPasswordMapOutput
	ToUserPasswordMapOutputWithContext(context.Context) UserPasswordMapOutput
}

type UserPasswordMap map[string]UserPasswordInput

func (UserPasswordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPassword)(nil)).Elem()
}

func (i UserPasswordMap) ToUserPasswordMapOutput() UserPasswordMapOutput {
	return i.ToUserPasswordMapOutputWithContext(context.Background())
}

func (i UserPasswordMap) ToUserPasswordMapOutputWithContext(ctx context.Context) UserPasswordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPasswordMapOutput)
}

type UserPasswordOutput struct{ *pulumi.OutputState }

func (UserPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPassword)(nil)).Elem()
}

func (o UserPasswordOutput) ToUserPasswordOutput() UserPasswordOutput {
	return o
}

func (o UserPasswordOutput) ToUserPasswordOutputWithContext(ctx context.Context) UserPasswordOutput {
	return o
}

// The source host of the user. Defaults to `localhost`.
func (o UserPasswordOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPassword) pulumi.StringPtrOutput { return v.Host }).(pulumi.StringPtrOutput)
}

func (o UserPasswordOutput) PlaintextPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserPassword) pulumi.StringPtrOutput { return v.PlaintextPassword }).(pulumi.StringPtrOutput)
}

func (o UserPasswordOutput) RetainOldPassword() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserPassword) pulumi.BoolPtrOutput { return v.RetainOldPassword }).(pulumi.BoolPtrOutput)
}

// The IAM user to associate with this access key.
func (o UserPasswordOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPassword) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type UserPasswordArrayOutput struct{ *pulumi.OutputState }

func (UserPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPassword)(nil)).Elem()
}

func (o UserPasswordArrayOutput) ToUserPasswordArrayOutput() UserPasswordArrayOutput {
	return o
}

func (o UserPasswordArrayOutput) ToUserPasswordArrayOutputWithContext(ctx context.Context) UserPasswordArrayOutput {
	return o
}

func (o UserPasswordArrayOutput) Index(i pulumi.IntInput) UserPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPassword {
		return vs[0].([]*UserPassword)[vs[1].(int)]
	}).(UserPasswordOutput)
}

type UserPasswordMapOutput struct{ *pulumi.OutputState }

func (UserPasswordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPassword)(nil)).Elem()
}

func (o UserPasswordMapOutput) ToUserPasswordMapOutput() UserPasswordMapOutput {
	return o
}

func (o UserPasswordMapOutput) ToUserPasswordMapOutputWithContext(ctx context.Context) UserPasswordMapOutput {
	return o
}

func (o UserPasswordMapOutput) MapIndex(k pulumi.StringInput) UserPasswordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPassword {
		return vs[0].(map[string]*UserPassword)[vs[1].(string)]
	}).(UserPasswordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPasswordInput)(nil)).Elem(), &UserPassword{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPasswordArrayInput)(nil)).Elem(), UserPasswordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPasswordMapInput)(nil)).Elem(), UserPasswordMap{})
	pulumi.RegisterOutputType(UserPasswordOutput{})
	pulumi.RegisterOutputType(UserPasswordArrayOutput{})
	pulumi.RegisterOutputType(UserPasswordMapOutput{})
}
