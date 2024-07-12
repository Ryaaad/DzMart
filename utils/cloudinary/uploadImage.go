package cloudinary

import (
	"DzMart/initializers"
	"context"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(ctx context.Context, img interface{}, Imgname string) (*uploader.UploadResult, error) {
	cld, clouderr := Credentials()
	if clouderr != nil {
		// fmt.Printf("Failed to connect to cloudinary: %v\n", clouderr)
		return nil, clouderr
	}
	folder := initializers.EnvCloudUploadFolder()

	resp, err := cld.Upload.Upload(ctx, img, uploader.UploadParams{
		PublicID:       Imgname,
		UniqueFilename: api.Bool(false),
		Folder:         folder,
		Overwrite:      api.Bool(true)})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
