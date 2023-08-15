// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package appsync

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	appsync_sdkv1 "github.com/aws/aws-sdk-go/service/appsync"
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
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAPICache,
			TypeName: "aws_appsync_api_cache",
		},
		{
			Factory:  ResourceAPIKey,
			TypeName: "aws_appsync_api_key",
		},
		{
			Factory:  ResourceDataSource,
			TypeName: "aws_appsync_datasource",
		},
		{
			Factory:  ResourceDomainName,
			TypeName: "aws_appsync_domain_name",
		},
		{
			Factory:  ResourceDomainNameAPIAssociation,
			TypeName: "aws_appsync_domain_name_api_association",
		},
		{
			Factory:  ResourceFunction,
			TypeName: "aws_appsync_function",
		},
		{
			Factory:  ResourceGraphQLAPI,
			TypeName: "aws_appsync_graphql_api",
			Name:     "GraphQL API",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceResolver,
			TypeName: "aws_appsync_resolver",
		},
		{
			Factory:  ResourceType,
			TypeName: "aws_appsync_type",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AppSync
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*appsync_sdkv1.AppSync, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return appsync_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
