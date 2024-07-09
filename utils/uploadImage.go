package utils

import (
	"DzMart/initializers"
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(ctx context.Context, img interface{}, Imgname string) {
	cld, clouderr := Credentials()
	if clouderr != nil {
		fmt.Printf("Failed to connect to cloudinary: %v\n", clouderr)
		return
	}
	folder := initializers.EnvCloudUploadFolder()
	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, img, uploader.UploadParams{
		PublicID:       Imgname,
		UniqueFilename: api.Bool(false),
		Folder:         folder,
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error")
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
}

func GetAssetInfo(ctx context.Context, PublicID string) (*admin.AssetResult, error) {
	cld, clouderr := Credentials()
	if clouderr != nil {
		return nil, fmt.Errorf("failed to connect to Cloudinary: %v", clouderr)
	}

	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{
		PublicID: "Dzmart-cloudinary/" + PublicID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get asset details: %v", err)
	}

	return resp, nil
}

func DestroyImg(ctx context.Context, PublicID string) (*uploader.DestroyResult, error) {
	cld, clouderr := Credentials()
	if clouderr != nil {
		return nil, fmt.Errorf("failed to connect to Cloudinary: %v", clouderr)
	}
	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: "Dzmart-cloudinary/" + PublicID,
	})
	if err != nil {
		// return nil,
		return nil, fmt.Errorf("failed to destroy assest: %v", err)
	}
	return resp, nil
}

// func DestroyImgs(ctx context.Context, PublicID string) (*uploader.DestroyResult, error) {
// 	cld, clouderr := Credentials()
// 	if clouderr != nil {
// 		return nil, fmt.Errorf("failed to connect to Cloudinary: %v", clouderr)
// 	}
// 	resp, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
// 		PublicID: "Dzmart-cloudinary/" + PublicID,
// 	})
// 	if err != nil {
// 		// return nil,
// 		return nil, fmt.Errorf("failed to destroy assest: %v", err)
// 	}
// 	return resp, nil
// }
