// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	apigateway_sdkv1 "github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAPIKey,
			TypeName: "aws_api_gateway_api_key",
		},
		{
			Factory:  DataSourceAuthorizer,
			TypeName: "aws_api_gateway_authorizer",
		},
		{
			Factory:  DataSourceAuthorizers,
			TypeName: "aws_api_gateway_authorizers",
		},
		{
			Factory:  DataSourceDomainName,
			TypeName: "aws_api_gateway_domain_name",
		},
		{
			Factory:  DataSourceExport,
			TypeName: "aws_api_gateway_export",
		},
		{
			Factory:  DataSourceResource,
			TypeName: "aws_api_gateway_resource",
		},
		{
			Factory:  DataSourceRestAPI,
			TypeName: "aws_api_gateway_rest_api",
		},
		{
			Factory:  DataSourceSdk,
			TypeName: "aws_api_gateway_sdk",
		},
		{
			Factory:  DataSourceVPCLink,
			TypeName: "aws_api_gateway_vpc_link",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAccount,
			TypeName: "aws_api_gateway_account",
		},
		{
			Factory:  ResourceAPIKey,
			TypeName: "aws_api_gateway_api_key",
			Name:     "API Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceAuthorizer,
			TypeName: "aws_api_gateway_authorizer",
		},
		{
			Factory:  ResourceBasePathMapping,
			TypeName: "aws_api_gateway_base_path_mapping",
		},
		{
			Factory:  ResourceClientCertificate,
			TypeName: "aws_api_gateway_client_certificate",
			Name:     "Client Certificate",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceDeployment,
			TypeName: "aws_api_gateway_deployment",
		},
		{
			Factory:  ResourceDocumentationPart,
			TypeName: "aws_api_gateway_documentation_part",
		},
		{
			Factory:  ResourceDocumentationVersion,
			TypeName: "aws_api_gateway_documentation_version",
		},
		{
			Factory:  ResourceDomainName,
			TypeName: "aws_api_gateway_domain_name",
			Name:     "Domain Name",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceGatewayResponse,
			TypeName: "aws_api_gateway_gateway_response",
		},
		{
			Factory:  ResourceIntegration,
			TypeName: "aws_api_gateway_integration",
		},
		{
			Factory:  ResourceIntegrationResponse,
			TypeName: "aws_api_gateway_integration_response",
		},
		{
			Factory:  ResourceMethod,
			TypeName: "aws_api_gateway_method",
		},
		{
			Factory:  ResourceMethodResponse,
			TypeName: "aws_api_gateway_method_response",
		},
		{
			Factory:  ResourceMethodSettings,
			TypeName: "aws_api_gateway_method_settings",
		},
		{
			Factory:  ResourceModel,
			TypeName: "aws_api_gateway_model",
		},
		{
			Factory:  ResourceRequestValidator,
			TypeName: "aws_api_gateway_request_validator",
		},
		{
			Factory:  ResourceResource,
			TypeName: "aws_api_gateway_resource",
		},
		{
			Factory:  ResourceRestAPI,
			TypeName: "aws_api_gateway_rest_api",
			Name:     "REST API",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceRestAPIPolicy,
			TypeName: "aws_api_gateway_rest_api_policy",
		},
		{
			Factory:  ResourceStage,
			TypeName: "aws_api_gateway_stage",
			Name:     "Stage",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceUsagePlan,
			TypeName: "aws_api_gateway_usage_plan",
			Name:     "Usage Plan",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceUsagePlanKey,
			TypeName: "aws_api_gateway_usage_plan_key",
		},
		{
			Factory:  ResourceVPCLink,
			TypeName: "aws_api_gateway_vpc_link",
			Name:     "VPC Link",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.APIGateway
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*apigateway_sdkv1.APIGateway, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return apigateway_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
