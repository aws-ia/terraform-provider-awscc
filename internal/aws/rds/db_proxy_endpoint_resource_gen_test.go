// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRDSDBProxyEndpoint_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RDS::DBProxyEndpoint", "awscc_rds_db_proxy_endpoint", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}