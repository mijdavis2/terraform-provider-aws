// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package swf

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	awstypes "github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listTags lists swf service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn *swf.Client, identifier string) (tftags.KeyValueTags, error) {
	input := &swf.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResource(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

// listTags_ returns a function that lists swf service tags and set them in Context.
// It is called by the transparent tagging interceptor.
func listTags_() types.ListTagsFunc {
	return func(ctx context.Context, meta any, identifier string) error {
		tags, err := listTags(ctx, meta.(*conns.AWSClient).SWFClient(ctx), identifier)

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

// Tags returns swf service tags.
func Tags(tags tftags.KeyValueTags) []awstypes.ResourceTag {
	result := make([]awstypes.ResourceTag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := awstypes.ResourceTag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from swf service tags.
func KeyValueTags(ctx context.Context, tags []awstypes.ResourceTag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.ToString(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns swf service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []awstypes.ResourceTag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets swf service tags in Context.
func setTagsOut(ctx context.Context, tags []awstypes.ResourceTag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates swf service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn *swf.Client, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.SWF)
	if len(removedTags) > 0 {
		input := &swf.UntagResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     removedTags.Keys(),
		}

		_, err := conn.UntagResource(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.SWF)
	if len(updatedTags) > 0 {
		input := &swf.TagResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.TagResource(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// updateTags_ returns a function that updates swf service tags.
// It is called by the transparent tagging interceptor.
func updateTags_() types.UpdateTagsFunc {
	return func(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
		return updateTags(ctx, meta.(*conns.AWSClient).SWFClient(ctx), identifier, oldTags, newTags)
	}
}
