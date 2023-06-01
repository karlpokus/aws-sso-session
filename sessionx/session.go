package sessionx

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// New returns an aws config based on sso credentials from a profile in ~/.aws/config
func New(ctx context.Context, profile, region string) (aws.Config, error) {
	return config.LoadDefaultConfig(
		ctx,
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)
}
