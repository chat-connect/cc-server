package key

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateKey() (key string, err error) {
	key, err = gonanoid.New(12)
	if err != nil {
		return "", err
	}
	
	return key, nil
}
