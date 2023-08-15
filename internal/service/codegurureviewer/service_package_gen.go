// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package codegurureviewer

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	codegurureviewer_sdkv1 "github.com/aws/aws-sdk-go/service/codegurureviewer"
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
			Factory:  ResourceRepositoryAssociation,
			TypeName: "aws_codegurureviewer_repository_association",
			Name:     "Repository Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.CodeGuruReviewer
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*codegurureviewer_sdkv1.CodeGuruReviewer, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return codegurureviewer_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
