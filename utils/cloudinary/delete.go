package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

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
