// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access IP ranges in your firewall rules.
// 
// https://cloud.google.com/compute/docs/load-balancing/health-checks#health_check_source_ips_and_firewall_rules
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/compute_lb_ip_ranges.html.markdown.
func GetLBIPRanges(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetLBIPRangesResult, error) {
	var rv GetLBIPRangesResult
	err := ctx.Invoke("gcp:compute/getLBIPRanges:getLBIPRanges", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getLBIPRanges.
type GetLBIPRangesResult struct {
	// The IP ranges used for health checks when **HTTP(S), SSL proxy, TCP proxy, and Internal load balancing** is used
	HttpSslTcpInternals []string `pulumi:"httpSslTcpInternals"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The IP ranges used for health checks when **Network load balancing** is used
	Networks []string `pulumi:"networks"`
}

