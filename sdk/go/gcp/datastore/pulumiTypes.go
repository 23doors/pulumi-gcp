// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package datastore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type DataStoreIndexProperty struct {
	Direction string `pulumi:"direction"`
	Name string `pulumi:"name"`
}

type DataStoreIndexPropertyInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput
	ToDataStoreIndexPropertyOutputWithContext(context.Context) DataStoreIndexPropertyOutput
}

type DataStoreIndexPropertyArgs struct {
	Direction pulumi.StringInput `pulumi:"direction"`
	Name pulumi.StringInput `pulumi:"name"`
}

func (DataStoreIndexPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexProperty)(nil)).Elem()
}

func (i DataStoreIndexPropertyArgs) ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput {
	return i.ToDataStoreIndexPropertyOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArgs) ToDataStoreIndexPropertyOutputWithContext(ctx context.Context) DataStoreIndexPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyOutput)
}

type DataStoreIndexPropertyArrayInput interface {
	pulumi.Input

	ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput
	ToDataStoreIndexPropertyArrayOutputWithContext(context.Context) DataStoreIndexPropertyArrayOutput
}

type DataStoreIndexPropertyArray []DataStoreIndexPropertyInput

func (DataStoreIndexPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexProperty)(nil)).Elem()
}

func (i DataStoreIndexPropertyArray) ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput {
	return i.ToDataStoreIndexPropertyArrayOutputWithContext(context.Background())
}

func (i DataStoreIndexPropertyArray) ToDataStoreIndexPropertyArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreIndexPropertyArrayOutput)
}

type DataStoreIndexPropertyOutput struct { *pulumi.OutputState }

func (DataStoreIndexPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreIndexProperty)(nil)).Elem()
}

func (o DataStoreIndexPropertyOutput) ToDataStoreIndexPropertyOutput() DataStoreIndexPropertyOutput {
	return o
}

func (o DataStoreIndexPropertyOutput) ToDataStoreIndexPropertyOutputWithContext(ctx context.Context) DataStoreIndexPropertyOutput {
	return o
}

func (o DataStoreIndexPropertyOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func (v DataStoreIndexProperty) string { return v.Direction }).(pulumi.StringOutput)
}

func (o DataStoreIndexPropertyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v DataStoreIndexProperty) string { return v.Name }).(pulumi.StringOutput)
}

type DataStoreIndexPropertyArrayOutput struct { *pulumi.OutputState}

func (DataStoreIndexPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreIndexProperty)(nil)).Elem()
}

func (o DataStoreIndexPropertyArrayOutput) ToDataStoreIndexPropertyArrayOutput() DataStoreIndexPropertyArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArrayOutput) ToDataStoreIndexPropertyArrayOutputWithContext(ctx context.Context) DataStoreIndexPropertyArrayOutput {
	return o
}

func (o DataStoreIndexPropertyArrayOutput) Index(i pulumi.IntInput) DataStoreIndexPropertyOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) DataStoreIndexProperty {
		return vs[0].([]DataStoreIndexProperty)[vs[1].(int)]
	}).(DataStoreIndexPropertyOutput)
}

func init() {
	pulumi.RegisterOutputType(DataStoreIndexPropertyOutput{})
	pulumi.RegisterOutputType(DataStoreIndexPropertyArrayOutput{})
}