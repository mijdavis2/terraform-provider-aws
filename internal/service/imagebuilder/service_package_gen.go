// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	imagebuilder_sdkv1 "github.com/aws/aws-sdk-go/service/imagebuilder"
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
			Factory:  DataSourceComponent,
			TypeName: "aws_imagebuilder_component",
		},
		{
			Factory:  DataSourceComponents,
			TypeName: "aws_imagebuilder_components",
		},
		{
			Factory:  DataSourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
		},
		{
			Factory:  DataSourceContainerRecipes,
			TypeName: "aws_imagebuilder_container_recipes",
		},
		{
			Factory:  DataSourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
		},
		{
			Factory:  DataSourceDistributionConfigurations,
			TypeName: "aws_imagebuilder_distribution_configurations",
		},
		{
			Factory:  DataSourceImage,
			TypeName: "aws_imagebuilder_image",
		},
		{
			Factory:  DataSourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
		},
		{
			Factory:  DataSourceImagePipelines,
			TypeName: "aws_imagebuilder_image_pipelines",
		},
		{
			Factory:  DataSourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
		},
		{
			Factory:  DataSourceImageRecipes,
			TypeName: "aws_imagebuilder_image_recipes",
		},
		{
			Factory:  DataSourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
		},
		{
			Factory:  DataSourceInfrastructureConfigurations,
			TypeName: "aws_imagebuilder_infrastructure_configurations",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceComponent,
			TypeName: "aws_imagebuilder_component",
			Name:     "Component",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceContainerRecipe,
			TypeName: "aws_imagebuilder_container_recipe",
			Name:     "Container Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceDistributionConfiguration,
			TypeName: "aws_imagebuilder_distribution_configuration",
			Name:     "Distribution Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceImage,
			TypeName: "aws_imagebuilder_image",
			Name:     "Image",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceImagePipeline,
			TypeName: "aws_imagebuilder_image_pipeline",
			Name:     "Image Pipeline",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceImageRecipe,
			TypeName: "aws_imagebuilder_image_recipe",
			Name:     "Image Recipe",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceInfrastructureConfiguration,
			TypeName: "aws_imagebuilder_infrastructure_configuration",
			Name:     "Infrastructure Configuration",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ImageBuilder
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*imagebuilder_sdkv1.Imagebuilder, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return imagebuilder_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
