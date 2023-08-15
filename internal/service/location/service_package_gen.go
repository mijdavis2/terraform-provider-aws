// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package location

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	locationservice_sdkv1 "github.com/aws/aws-sdk-go/service/locationservice"
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
			Factory:  DataSourceGeofenceCollection,
			TypeName: "aws_location_geofence_collection",
		},
		{
			Factory:  DataSourceMap,
			TypeName: "aws_location_map",
		},
		{
			Factory:  DataSourcePlaceIndex,
			TypeName: "aws_location_place_index",
		},
		{
			Factory:  DataSourceRouteCalculator,
			TypeName: "aws_location_route_calculator",
		},
		{
			Factory:  DataSourceTracker,
			TypeName: "aws_location_tracker",
		},
		{
			Factory:  DataSourceTrackerAssociation,
			TypeName: "aws_location_tracker_association",
		},
		{
			Factory:  DataSourceTrackerAssociations,
			TypeName: "aws_location_tracker_associations",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceGeofenceCollection,
			TypeName: "aws_location_geofence_collection",
			Name:     "Geofence Collection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "collection_arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceMap,
			TypeName: "aws_location_map",
			Name:     "Map",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "map_arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourcePlaceIndex,
			TypeName: "aws_location_place_index",
			Name:     "Map",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "index_arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceRouteCalculator,
			TypeName: "aws_location_route_calculator",
			Name:     "Route Calculator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "calculator_arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceTracker,
			TypeName: "aws_location_tracker",
			Name:     "Route Calculator",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "tracker_arn",
				ListTags:            listTags_(),
				UpdateTags:          updateTags_(),
			},
		},
		{
			Factory:  ResourceTrackerAssociation,
			TypeName: "aws_location_tracker_association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Location
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*locationservice_sdkv1.LocationService, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return locationservice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
