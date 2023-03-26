package images

import (
	"bytes"
	"context"
	"io"
	"nbid-files/utils/config"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var path = "./../../.env"
var cloud Cloudinary

func init() {
	config.LoadConfig(path)
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloud = NewCloudinary(cloudName, apiKey, apiSecret)
}

func TestBuildCloudinary(t *testing.T) {
	cloudName := os.Getenv("CLOUDINARY_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cloudinary := NewCloudinary(cloudName, apiKey, apiSecret)

	require.Nil(t, cloudinary.IsError)
}

func TestUploadByBuffer(t *testing.T) {
	imagePath := "img.jpg"
	file, err := os.Open(imagePath)
	require.Nil(t, err)

	defer file.Close()

	buffer := bytes.NewBuffer(nil)

	io.Copy(buffer, file)

	url, err := cloud.Upload(context.Background(), buffer, "christo@pcr.ac.id/products")
	require.Nil(t, err)
	require.NotEmpty(t, url)
}
func TestUploadByURL(t *testing.T) {
	imagePath := "https://media.suara.com/pictures/653x366/2022/01/03/24746-upin-dan-ipin.jpg"

	url, err := cloud.Upload(context.Background(), imagePath, "christo@pcr.ac.id/products")
	require.Nil(t, err)
	require.NotEmpty(t, url)
}

func TestRemove(t *testing.T) {
	path := "NBID-Training/IniDariPublicId"
	err := cloud.Remove(context.Background(), path)
	require.Nil(t, err)
}
