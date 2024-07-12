package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2/api/admin"
)

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
