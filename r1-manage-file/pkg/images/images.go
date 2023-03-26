package images

import "context"

type ImageInterface interface {
	Upload(ctx context.Context, file interface{}, path string) (string, error)
	Remove(ctx context.Context, path string) error
}

type Image struct {
	Bucket    string
	APIKey    string
	APISecret string
}

func NewImage(bucket, apiKey, apiSecret string) Image {
	return Image{
		Bucket:    bucket,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

func (i Image) BuildCloudinary() Cloudinary {
	return NewCloudinary(i.Bucket, i.APIKey, i.APISecret)
}

func (i Image) BuildGCS() Image {
	return NewImage(i.Bucket, i.APIKey, i.APISecret)
}
func (i Image) Upload(ctx context.Context, file interface{}, path string) (string, error) {
	panic("unimplement")
}

func (i Image) Remove(ctx context.Context, path string) error {
	panic("unimplement")
}
