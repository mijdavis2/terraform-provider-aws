// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package dms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// listTags lists dms service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func listTags(ctx context.Context, conn databasemigrationserviceiface.DatabaseMigrationServiceAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &databasemigrationservice.ListTagsForResourceInput{
		ResourceArn: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.TagList), nil
}

// listTags_ returns a function that lists dms service tags and set them in Context.
// It is called by the transparent tagging interceptor.
func listTags_() types.ListTagsFunc {
	return func(ctx context.Context, meta any, identifier string) error {
		tags, err := listTags(ctx, meta.(*conns.AWSClient).DMSConn(ctx), identifier)

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

// Tags returns dms service tags.
func Tags(tags tftags.KeyValueTags) []*databasemigrationservice.Tag {
	result := make([]*databasemigrationservice.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &databasemigrationservice.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from databasemigrationservice service tags.
func KeyValueTags(ctx context.Context, tags []*databasemigrationservice.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns dms service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []*databasemigrationservice.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets dms service tags in Context.
func setTagsOut(ctx context.Context, tags []*databasemigrationservice.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates dms service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn databasemigrationserviceiface.DatabaseMigrationServiceAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.DMS)
	if len(removedTags) > 0 {
		input := &databasemigrationservice.RemoveTagsFromResourceInput{
			ResourceArn: aws.String(identifier),
			TagKeys:     aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.RemoveTagsFromResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.DMS)
	if len(updatedTags) > 0 {
		input := &databasemigrationservice.AddTagsToResourceInput{
			ResourceArn: aws.String(identifier),
			Tags:        Tags(updatedTags),
		}

		_, err := conn.AddTagsToResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// updateTags_ returns a function that updates dms service tags.
// It is called by the transparent tagging interceptor.
func updateTags_() types.UpdateTagsFunc {
	return func(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
		return updateTags(ctx, meta.(*conns.AWSClient).DMSConn(ctx), identifier, oldTags, newTags)
	}
}
