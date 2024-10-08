// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unobravo/pulumi-mysql/sdk/go/mysql/internal"
)

// The “getDatabases“ gets databases on a MySQL
// server.
//
// ## Example Usage
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
//			_, err := mysql.GetDatabases(ctx, &mysql.GetDatabasesArgs{
//				Pattern: pulumi.StringRef("test_%"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetDatabases(ctx *pulumi.Context, args *GetDatabasesArgs, opts ...pulumi.InvokeOption) (*GetDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabasesResult
	err := ctx.Invoke("mysql:index/getDatabases:getDatabases", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesArgs struct {
	// Patterns for searching databases.
	Pattern *string `pulumi:"pattern"`
}

// A collection of values returned by getDatabases.
type GetDatabasesResult struct {
	// The list of the database names.
	Databases []string `pulumi:"databases"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Pattern *string `pulumi:"pattern"`
}

func GetDatabasesOutput(ctx *pulumi.Context, args GetDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetDatabasesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabasesResult, error) {
			args := v.(GetDatabasesArgs)
			r, err := GetDatabases(ctx, &args, opts...)
			var s GetDatabasesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabasesResultOutput)
}

// A collection of arguments for invoking getDatabases.
type GetDatabasesOutputArgs struct {
	// Patterns for searching databases.
	Pattern pulumi.StringPtrInput `pulumi:"pattern"`
}

func (GetDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getDatabases.
type GetDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabasesResult)(nil)).Elem()
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutput() GetDatabasesResultOutput {
	return o
}

func (o GetDatabasesResultOutput) ToGetDatabasesResultOutputWithContext(ctx context.Context) GetDatabasesResultOutput {
	return o
}

// The list of the database names.
func (o GetDatabasesResultOutput) Databases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabasesResult) []string { return v.Databases }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatabasesResultOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabasesResult) *string { return v.Pattern }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabasesResultOutput{})
}
