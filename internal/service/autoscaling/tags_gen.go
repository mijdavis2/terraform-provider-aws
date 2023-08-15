// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package autoscaling

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// GetTag fetches an individual autoscaling service tag for a resource.
// Returns whether the key value and any errors. A NotFoundError is used to signal that no value was found.
// This function will optimise the handling over listTags, if possible.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func GetTag(ctx context.Context, conn autoscalingiface.AutoScalingAPI, identifier, resourceType, key string) (*tftags.TagData, error) {
	input := &autoscaling.DescribeTagsInput{
		Filters: []*autoscaling.Filter{
			{
				Name:   aws.String("auto-scaling-group"),
				Values: []*string{aws.String(identifier)},
			},
			{
				Name:   aws.String("key"),
				Values: []*string{aws.String(key)},
			},
		},
	}

	output, err := conn.DescribeTagsWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	listTags := KeyValueTags(ctx, output.Tags, identifier, resourceType)

	if !listTags.KeyExists(key) {
		return nil, tfresource.NewEmptyResultError(nil)
	}

	return listTags.KeyTagData(key), nil
}

// listTags lists autoscaling service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn autoscalingiface.AutoScalingAPI, identifier, resourceType string) (tftags.KeyValueTags, error) {
	input := &autoscaling.DescribeTagsInput{
		Filters: []*autoscaling.Filter{
			{
				Name:   aws.String("auto-scaling-group"),
				Values: []*string{aws.String(identifier)},
			},
		},
	}

	output, err := conn.DescribeTagsWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags, identifier, resourceType), nil
}

// listTags_ returns a function that lists autoscaling service tags and set them in Context.
// It is called by the transparent tagging interceptor.
func listTags_(resourceType string) types.ListTagsFunc {
	return func(ctx context.Context, meta any, identifier string) error {
		tags, err := listTags(ctx, meta.(*conns.AWSClient).AutoScalingConn(ctx), identifier, resourceType)

		if err != nil {
			return err
		}

		if inContext, ok := tftags.FromContext(ctx); ok {
			inContext.TagsOut = types.Some(tags)
		}

		return nil
	}
}

// []*SERVICE.Tag handling

// ListOfMap returns a list of autoscaling in flattened map.
//
// Compatible with setting Terraform state for strongly typed configuration blocks.
//
// This function strips tag resource identifier and type. Generally, this is
// the desired behavior so the tag schema does not require those attributes.
// Use (tftags.KeyValueTags).ListOfMap() for full tag information.
func ListOfMap(tags tftags.KeyValueTags) []any {
	var result []any

	for _, key := range tags.Keys() {
		m := map[string]any{
			"key":   key,
			"value": aws.StringValue(tags.KeyValue(key)),

			"propagate_at_launch": aws.BoolValue(tags.KeyAdditionalBoolValue(key, "PropagateAtLaunch")),
		}

		result = append(result, m)
	}

	return result
}

// Tags returns autoscaling service tags.
func Tags(tags tftags.KeyValueTags) []*autoscaling.Tag {
	var result []*autoscaling.Tag

	for _, key := range tags.Keys() {
		tag := &autoscaling.Tag{
			Key:               aws.String(key),
			Value:             tags.KeyValue(key),
			ResourceId:        tags.KeyAdditionalStringValue(key, "ResourceId"),
			ResourceType:      tags.KeyAdditionalStringValue(key, "ResourceType"),
			PropagateAtLaunch: tags.KeyAdditionalBoolValue(key, "PropagateAtLaunch"),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from autoscaling service tags.
//
// Accepts the following types:
//   - []*autoscaling.Tag
//   - []*autoscaling.TagDescription
//   - []any (Terraform TypeList configuration block compatible)
//   - *schema.Set (Terraform TypeSet configuration block compatible)
func KeyValueTags(ctx context.Context, tags any, identifier, resourceType string) tftags.KeyValueTags {
	switch tags := tags.(type) {
	case []*autoscaling.Tag:
		m := make(map[string]*tftags.TagData, len(tags))

		for _, tag := range tags {
			tagData := &tftags.TagData{
				Value: tag.Value,
			}

			tagData.AdditionalBoolFields = make(map[string]*bool)
			tagData.AdditionalBoolFields["PropagateAtLaunch"] = tag.PropagateAtLaunch
			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			m[aws.StringValue(tag.Key)] = tagData
		}

		return tftags.New(ctx, m)
	case []*autoscaling.TagDescription:
		m := make(map[string]*tftags.TagData, len(tags))

		for _, tag := range tags {
			tagData := &tftags.TagData{
				Value: tag.Value,
			}
			tagData.AdditionalBoolFields = make(map[string]*bool)
			tagData.AdditionalBoolFields["PropagateAtLaunch"] = tag.PropagateAtLaunch
			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			m[aws.StringValue(tag.Key)] = tagData
		}

		return tftags.New(ctx, m)
	case *schema.Set:
		return KeyValueTags(ctx, tags.List(), identifier, resourceType)
	case []any:
		result := make(map[string]*tftags.TagData)

		for _, tfMapRaw := range tags {
			tfMap, ok := tfMapRaw.(map[string]any)

			if !ok {
				continue
			}

			key, ok := tfMap["key"].(string)

			if !ok {
				continue
			}

			tagData := &tftags.TagData{}

			if v, ok := tfMap["value"].(string); ok {
				tagData.Value = &v
			}

			tagData.AdditionalBoolFields = make(map[string]*bool)
			if v, ok := tfMap["propagate_at_launch"].(bool); ok {
				tagData.AdditionalBoolFields["PropagateAtLaunch"] = &v
			}

			tagData.AdditionalStringFields = make(map[string]*string)
			tagData.AdditionalStringFields["ResourceId"] = &identifier
			tagData.AdditionalStringFields["ResourceType"] = &resourceType

			result[key] = tagData
		}

		return tftags.New(ctx, result)
	default:
		return tftags.New(ctx, nil)
	}
}

// getTagsIn returns autoscaling service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []*autoscaling.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets autoscaling service tags in Context.
func setTagsOut(ctx context.Context, tags any, identifier, resourceType string) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags, identifier, resourceType))
	}
}

// updateTags updates autoscaling service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn autoscalingiface.AutoScalingAPI, identifier, resourceType string, oldTagsSet, newTagsSet any) error {
	oldTags := KeyValueTags(ctx, oldTagsSet, identifier, resourceType)
	newTags := KeyValueTags(ctx, newTagsSet, identifier, resourceType)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.AutoScaling)
	if len(removedTags) > 0 {
		input := &autoscaling.DeleteTagsInput{
			Tags: Tags(removedTags),
		}

		_, err := conn.DeleteTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.AutoScaling)
	if len(updatedTags) > 0 {
		input := &autoscaling.CreateOrUpdateTagsInput{
			Tags: Tags(updatedTags),
		}

		_, err := conn.CreateOrUpdateTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// updateTags_ returns a function that updates autoscaling service tags.
// It is called by the transparent tagging interceptor.
func updateTags_(resourceType string) types.UpdateTagsFunc {
	return func(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
		return updateTags(ctx, meta.(*conns.AWSClient).AutoScalingConn(ctx), identifier, resourceType, oldTags, newTags)
	}
}
