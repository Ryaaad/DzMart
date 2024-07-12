package cloudinary

import (
	"DzMart/initializers"
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func RenameImage(ctx context.Context, OldID string, NewID string) (*uploader.RenameResult, error) {
	cld, clouderr := Credentials()
	if clouderr != nil {
		return nil, fmt.Errorf("failed to connect to cloudinary: %v", clouderr)
	}
	folder := initializers.EnvCloudUploadFolder()
	resp, err := cld.Upload.Rename(ctx, uploader.RenameParams{
		FromPublicID: folder + "/" + OldID,
		ToPublicID:   folder + "/" + NewID})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
