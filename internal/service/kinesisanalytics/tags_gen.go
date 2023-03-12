// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package kinesisanalytics

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go/service/kinesisanalytics/kinesisanalyticsiface"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// ListTags lists kinesisanalytics service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(ctx context.Context, conn kinesisanalyticsiface.KinesisAnalyticsAPI, identifier string) (tftags.KeyValueTags, error) {
	input := &kinesisanalytics.ListTagsForResourceInput{
		ResourceARN: aws.String(identifier),
	}

	output, err := conn.ListTagsForResourceWithContext(ctx, input)

	if err != nil {
		return tftags.New(ctx, nil), err
	}

	return KeyValueTags(ctx, output.Tags), nil
}

func (p *servicePackage) ListTags(ctx context.Context, meta any, identifier string) (tftags.KeyValueTags, error) {
	return ListTags(ctx, meta.(*conns.AWSClient).KinesisAnalyticsConn(), identifier)
}

// []*SERVICE.Tag handling

// Tags returns kinesisanalytics service tags.
func Tags(tags tftags.KeyValueTags) []*kinesisanalytics.Tag {
	result := make([]*kinesisanalytics.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &kinesisanalytics.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from kinesisanalytics service tags.
func KeyValueTags(ctx context.Context, tags []*kinesisanalytics.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(ctx, m)
}

// UpdateTags updates kinesisanalytics service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.

func UpdateTags(ctx context.Context, conn kinesisanalyticsiface.KinesisAnalyticsAPI, identifier string, oldTagsMap, newTagsMap any) error {
	oldTags := tftags.New(ctx, oldTagsMap)
	newTags := tftags.New(ctx, newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &kinesisanalytics.UntagResourceInput{
			ResourceARN: aws.String(identifier),
			TagKeys:     aws.StringSlice(removedTags.IgnoreAWS().Keys()),
		}

		_, err := conn.UntagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &kinesisanalytics.TagResourceInput{
			ResourceARN: aws.String(identifier),
			Tags:        Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.TagResourceWithContext(ctx, input)

		if err != nil {
			return fmt.Errorf("tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}

func (p *servicePackage) UpdateTags(ctx context.Context, meta any, identifier string, oldTags, newTags any) error {
	return UpdateTags(ctx, meta.(*conns.AWSClient).KinesisAnalyticsConn(), identifier, oldTags, newTags)
}
