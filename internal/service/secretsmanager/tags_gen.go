// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package secretsmanager

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func listTags_() types.ListTagsFunc { return nil }

// []*SERVICE.Tag handling

// Tags returns secretsmanager service tags.
func Tags(tags tftags.KeyValueTags) []*secretsmanager.Tag {
	result := make([]*secretsmanager.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &secretsmanager.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from secretsmanager service tags.
func KeyValueTags(ctx context.Context, tags []*secretsmanager.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// getTagsIn returns secretsmanager service tags from Context.
// nil is returned if there are no input tags.
func getTagsIn(ctx context.Context) []*secretsmanager.Tag {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// setTagsOut sets secretsmanager service tags in Context.
func setTagsOut(ctx context.Context, tags []*secretsmanager.Tag) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}

// updateTags updates secretsmanager service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func updateTags(ctx context.Context, conn secretsmanageriface.SecretsManagerAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	removedTags := oldTags.Removed(newTags)
	removedTags = removedTags.IgnoreSystem(names.SecretsManager)
	if len(removedTags) > 0 {
		input := &secretsmanager.UntagResourceInput{
			SecretId: aws.String(identifier),
			TagKeys:  aws.StringSlice(removedTags.Keys()),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	updatedTags := oldTags.Updated(newTags)
	updatedTags = updatedTags.IgnoreSystem(names.SecretsManager)
	if len(updatedTags) > 0 {
		input := &secretsmanager.TagResourceInput{
			SecretId: aws.String(identifier),
			Tags:     Tags(updatedTags),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

// updateTags_ returns a function that updates secretsmanager service tags.
// It is called by the transparent tagging interceptor.
func updateTags_() types.UpdateTagsFunc {
	return func(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
		return updateTags(ctx, meta.(*conns.AWSClient).SecretsManagerConn(ctx), identifier, oldTags, newTags)
	}
}
