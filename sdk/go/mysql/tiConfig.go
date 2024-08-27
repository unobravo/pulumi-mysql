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

// The “TiConfig“ resource manages a TiKV or PD variables on a TiDB cluster.
//
// > **Note on TiDB:** Possible TiKV or PD variables are available [here](https://docs.pingcap.com/tidb/stable/dynamic-config)
//
// > **Note about `destroy`:** `destroy` is trying restore default values as described here.
//
//	Unfortunately not every variable support this.
//
// ## Example Usage
//
// ### PD
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/unobravo/pulumi-mysql/sdk/go/mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mysql.NewTiConfig(ctx, "logLevel", &mysql.TiConfigArgs{
//				Type:  pulumi.String("pd"),
//				Value: pulumi.String("warn"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Set variable for all PD instances
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/unobravo/pulumi-mysql/sdk/go/mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mysql.NewTiConfig(ctx, "logLevel", &mysql.TiConfigArgs{
//				Type:  pulumi.String("pd"),
//				Value: pulumi.String("warn"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ### Set variable for one PD instance only
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/unobravo/pulumi-mysql/sdk/go/mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mysql.NewTiConfig(ctx, "logLevel", &mysql.TiConfigArgs{
//				Instance: pulumi.String("127.0.0.1:2379"),
//				Type:     pulumi.String("pd"),
//				Value:    pulumi.String("warn"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## TiKV
//
// ### Set varibale for all TiKV instances
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/unobravo/pulumi-mysql/sdk/go/mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mysql.NewTiConfig(ctx, "splitQpsThreshold", &mysql.TiConfigArgs{
//				Type:  pulumi.String("tikv"),
//				Value: pulumi.String("100"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// #### Set variable for one TiKV instance only
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/unobravo/pulumi-mysql/sdk/go/mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mysql.NewTiConfig(ctx, "splitQpsThreshold", &mysql.TiConfigArgs{
//				Instance: pulumi.String("127.0.0.1:20180"),
//				Type:     pulumi.String("tikv"),
//				Value:    pulumi.String("10"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// ### Simple example
//
// #### TiKV example
//
// ```sh
// $ pulumi import mysql:index/tiConfig:TiConfig split_qps_threshold' 'tikv#split-qps-threshold'
// ```
//
// # Import value for specific instance
//
// ```sh
// $ pulumi import mysql:index/tiConfig:TiConfig split_qps_threshold' 'tikv#split-qps-threshold#127.0.0.1:20180'
// ```
//
// #### PD example
//
// ```sh
// $ pulumi import mysql:index/tiConfig:TiConfig log_level' 'pd#log.level'
// ```
type TiConfig struct {
	pulumi.CustomResourceState

	Instance pulumi.StringPtrOutput `pulumi:"instance"`
	// The name of the configuration variable.
	Name pulumi.StringOutput `pulumi:"name"`
	// The instance type to configure. Possible values are tikv or pd.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the configuration variable as string.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewTiConfig registers a new resource with the given unique name, arguments, and options.
func NewTiConfig(ctx *pulumi.Context,
	name string, args *TiConfigArgs, opts ...pulumi.ResourceOption) (*TiConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TiConfig
	err := ctx.RegisterResource("mysql:index/tiConfig:TiConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTiConfig gets an existing TiConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTiConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TiConfigState, opts ...pulumi.ResourceOption) (*TiConfig, error) {
	var resource TiConfig
	err := ctx.ReadResource("mysql:index/tiConfig:TiConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TiConfig resources.
type tiConfigState struct {
	Instance *string `pulumi:"instance"`
	// The name of the configuration variable.
	Name *string `pulumi:"name"`
	// The instance type to configure. Possible values are tikv or pd.
	Type *string `pulumi:"type"`
	// The value of the configuration variable as string.
	Value *string `pulumi:"value"`
}

type TiConfigState struct {
	Instance pulumi.StringPtrInput
	// The name of the configuration variable.
	Name pulumi.StringPtrInput
	// The instance type to configure. Possible values are tikv or pd.
	Type pulumi.StringPtrInput
	// The value of the configuration variable as string.
	Value pulumi.StringPtrInput
}

func (TiConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*tiConfigState)(nil)).Elem()
}

type tiConfigArgs struct {
	Instance *string `pulumi:"instance"`
	// The name of the configuration variable.
	Name *string `pulumi:"name"`
	// The instance type to configure. Possible values are tikv or pd.
	Type string `pulumi:"type"`
	// The value of the configuration variable as string.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a TiConfig resource.
type TiConfigArgs struct {
	Instance pulumi.StringPtrInput
	// The name of the configuration variable.
	Name pulumi.StringPtrInput
	// The instance type to configure. Possible values are tikv or pd.
	Type pulumi.StringInput
	// The value of the configuration variable as string.
	Value pulumi.StringInput
}

func (TiConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tiConfigArgs)(nil)).Elem()
}

type TiConfigInput interface {
	pulumi.Input

	ToTiConfigOutput() TiConfigOutput
	ToTiConfigOutputWithContext(ctx context.Context) TiConfigOutput
}

func (*TiConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**TiConfig)(nil)).Elem()
}

func (i *TiConfig) ToTiConfigOutput() TiConfigOutput {
	return i.ToTiConfigOutputWithContext(context.Background())
}

func (i *TiConfig) ToTiConfigOutputWithContext(ctx context.Context) TiConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiConfigOutput)
}

// TiConfigArrayInput is an input type that accepts TiConfigArray and TiConfigArrayOutput values.
// You can construct a concrete instance of `TiConfigArrayInput` via:
//
//	TiConfigArray{ TiConfigArgs{...} }
type TiConfigArrayInput interface {
	pulumi.Input

	ToTiConfigArrayOutput() TiConfigArrayOutput
	ToTiConfigArrayOutputWithContext(context.Context) TiConfigArrayOutput
}

type TiConfigArray []TiConfigInput

func (TiConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TiConfig)(nil)).Elem()
}

func (i TiConfigArray) ToTiConfigArrayOutput() TiConfigArrayOutput {
	return i.ToTiConfigArrayOutputWithContext(context.Background())
}

func (i TiConfigArray) ToTiConfigArrayOutputWithContext(ctx context.Context) TiConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiConfigArrayOutput)
}

// TiConfigMapInput is an input type that accepts TiConfigMap and TiConfigMapOutput values.
// You can construct a concrete instance of `TiConfigMapInput` via:
//
//	TiConfigMap{ "key": TiConfigArgs{...} }
type TiConfigMapInput interface {
	pulumi.Input

	ToTiConfigMapOutput() TiConfigMapOutput
	ToTiConfigMapOutputWithContext(context.Context) TiConfigMapOutput
}

type TiConfigMap map[string]TiConfigInput

func (TiConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TiConfig)(nil)).Elem()
}

func (i TiConfigMap) ToTiConfigMapOutput() TiConfigMapOutput {
	return i.ToTiConfigMapOutputWithContext(context.Background())
}

func (i TiConfigMap) ToTiConfigMapOutputWithContext(ctx context.Context) TiConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiConfigMapOutput)
}

type TiConfigOutput struct{ *pulumi.OutputState }

func (TiConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TiConfig)(nil)).Elem()
}

func (o TiConfigOutput) ToTiConfigOutput() TiConfigOutput {
	return o
}

func (o TiConfigOutput) ToTiConfigOutputWithContext(ctx context.Context) TiConfigOutput {
	return o
}

func (o TiConfigOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiConfig) pulumi.StringPtrOutput { return v.Instance }).(pulumi.StringPtrOutput)
}

// The name of the configuration variable.
func (o TiConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TiConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The instance type to configure. Possible values are tikv or pd.
func (o TiConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TiConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The value of the configuration variable as string.
func (o TiConfigOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *TiConfig) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type TiConfigArrayOutput struct{ *pulumi.OutputState }

func (TiConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TiConfig)(nil)).Elem()
}

func (o TiConfigArrayOutput) ToTiConfigArrayOutput() TiConfigArrayOutput {
	return o
}

func (o TiConfigArrayOutput) ToTiConfigArrayOutputWithContext(ctx context.Context) TiConfigArrayOutput {
	return o
}

func (o TiConfigArrayOutput) Index(i pulumi.IntInput) TiConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TiConfig {
		return vs[0].([]*TiConfig)[vs[1].(int)]
	}).(TiConfigOutput)
}

type TiConfigMapOutput struct{ *pulumi.OutputState }

func (TiConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TiConfig)(nil)).Elem()
}

func (o TiConfigMapOutput) ToTiConfigMapOutput() TiConfigMapOutput {
	return o
}

func (o TiConfigMapOutput) ToTiConfigMapOutputWithContext(ctx context.Context) TiConfigMapOutput {
	return o
}

func (o TiConfigMapOutput) MapIndex(k pulumi.StringInput) TiConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TiConfig {
		return vs[0].(map[string]*TiConfig)[vs[1].(string)]
	}).(TiConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TiConfigInput)(nil)).Elem(), &TiConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiConfigArrayInput)(nil)).Elem(), TiConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiConfigMapInput)(nil)).Elem(), TiConfigMap{})
	pulumi.RegisterOutputType(TiConfigOutput{})
	pulumi.RegisterOutputType(TiConfigArrayOutput{})
	pulumi.RegisterOutputType(TiConfigMapOutput{})
}
