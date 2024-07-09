package utils

import (
	"DzMart/initializers"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
)

func Credentials() (*cloudinary.Cloudinary, error) {
	CloudName := initializers.EnvCloudName()
	API_SECRET := initializers.EnvCloudAPISecret()
	API_KEY := initializers.EnvCloudAPIKey()

	cld, err := cloudinary.NewFromParams(CloudName, API_KEY, API_SECRET)
	if err != nil {
		fmt.Printf("Failed to initialize Cloudinary: %v\n", err)
		return nil, err
	}
	return cld, nil
}
