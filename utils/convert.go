package utils

import "strconv"

func ToUint(id string) (uint, error) {
	ID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}
	uintID := uint(ID)
	return uintID, nil
}
