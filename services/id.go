package services

import (
	"errors"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateID(prefix string) (*string, error) {
	if len(prefix) != 3 {
		return nil, errors.New("prefix needs to be 4 characters long")
	}
	id, _ := gonanoid.New(12)

	str := fmt.Sprintf("%v_%v", prefix, id)
	return &str, nil
}
