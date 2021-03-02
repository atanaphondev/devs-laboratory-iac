package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/storage"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// GCS configuration
		opts := &storage.BucketArgs{
			Name:         pulumi.String("iac-pulumi-atanaphondev-devs-laboratory-dev"),
			Location:     pulumi.String("ASIA-SOUTHEAST1"),
			StorageClass: pulumi.String("STANDARD"),
		}

		// Create a GCP resource (Storage Bucket)
		bucket, err := storage.NewBucket(ctx, "iac-pulumi-atanaphondev-devs-laboratory-dev", opts)
		if err != nil {
			return err
		}

		// Export the DNS name of the bucket
		ctx.Export("bucketName", bucket.Url)
		return nil
	})
}
