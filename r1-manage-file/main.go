package main

import (
	"context"
	"nbid-files/pkg/images"
	"nbid-files/utils/config"
	"os"
)

type CloudSvc interface {
	Upload(ctx context.Context, file interface{}, path string) (string, error)
	Remove(ctx context.Context, path string) error
}

type Services struct {
	cloud CloudSvc
}

func main() {
	config.LoadConfig(".env")
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	image := images.NewImage(cloudName, apiKey, apiSecret)

	svc := Services{
		cloud: image.BuildCloudinary(),
	}

	svc.cloud.Upload(context.Background(), "", "")
}
