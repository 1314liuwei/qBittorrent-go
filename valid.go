package qBittorent

import "errors"

func IsValidHash(hash string) error {
	if len(hash) != 32 {
		return errors.New("this is not a valid hash value")
	}

	return nil
}
