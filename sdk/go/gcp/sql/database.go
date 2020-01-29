// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/sql_database.html.markdown.
type Database struct {
	pulumi.CustomResourceState

	// The charset value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
	// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values. Postgres
	// databases only support a value of 'UTF8' at creation time.
	Charset pulumi.StringOutput `pulumi:"charset"`
	// The collation value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
	// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values. Postgres
	// databases only support a value of 'en_US.UTF8' at creation time.
	Collation pulumi.StringOutput `pulumi:"collation"`
	// The name of the Cloud SQL instance. This does not include the project ID.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	var resource Database
	err := ctx.RegisterResource("gcp:sql/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("gcp:sql/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// The charset value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
	// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values. Postgres
	// databases only support a value of 'UTF8' at creation time.
	Charset *string `pulumi:"charset"`
	// The collation value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
	// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values. Postgres
	// databases only support a value of 'en_US.UTF8' at creation time.
	Collation *string `pulumi:"collation"`
	// The name of the Cloud SQL instance. This does not include the project ID.
	Instance *string `pulumi:"instance"`
	// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	SelfLink *string `pulumi:"selfLink"`
}

type DatabaseState struct {
	// The charset value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
	// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values. Postgres
	// databases only support a value of 'UTF8' at creation time.
	Charset pulumi.StringPtrInput
	// The collation value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
	// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values. Postgres
	// databases only support a value of 'en_US.UTF8' at creation time.
	Collation pulumi.StringPtrInput
	// The name of the Cloud SQL instance. This does not include the project ID.
	Instance pulumi.StringPtrInput
	// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// The charset value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
	// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values. Postgres
	// databases only support a value of 'UTF8' at creation time.
	Charset *string `pulumi:"charset"`
	// The collation value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
	// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values. Postgres
	// databases only support a value of 'en_US.UTF8' at creation time.
	Collation *string `pulumi:"collation"`
	// The name of the Cloud SQL instance. This does not include the project ID.
	Instance string `pulumi:"instance"`
	// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// The charset value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set
	// Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values. Postgres
	// databases only support a value of 'UTF8' at creation time.
	Charset pulumi.StringPtrInput
	// The collation value. See MySQL's [Supported Character Sets and
	// Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation
	// Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values. Postgres
	// databases only support a value of 'en_US.UTF8' at creation time.
	Collation pulumi.StringPtrInput
	// The name of the Cloud SQL instance. This does not include the project ID.
	Instance pulumi.StringInput
	// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

