// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package redshift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/redshift/redshiftiface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func listTags_() types.ListTagsFunc { return nil }

// []*SERVICE.Tag handling

// Tags returns redshift service tags.
func Tags(tags tftags.KeyValueTags) []*redshift.Tag {
	result := make([]*redshift.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &redshift.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from redshift service tags.
func KeyValueTags(ctx context.Context, tags []*redshift.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns redshift service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []*redshift.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets redshift service tags in Context.
func setTagsOut(ctx context.Context, tags []*redshift.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates redshift service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn redshiftiface.RedshiftAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.Redshift)
	if len(removedTags) > 0 {
		input := &redshift.DeleteTagsInput{
			ResourceName: aws.String(identifier),
			TagKeys:      aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.DeleteTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.Redshift)
	if len(updatedTags) > 0 {
		input := &redshift.CreateTagsInput{
			ResourceName: aws.String(identifier),
			Tags:         Tags(updatedTags),
		}

		_, err := conn.CreateTagsWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// updateTags_ returns a function that updates redshift service tags.
// It is called by the transparent tagging interceptor.
func updateTags_() types.UpdateTagsFunc {
	return func(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
		return updateTags(ctx, meta.(*conns.AWSClient).RedshiftConn(ctx), identifier, oldTags, newTags)
	}
}
