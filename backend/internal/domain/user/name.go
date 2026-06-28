package user

import (
	"errors"
	"strings"
)

type Name string

func NewName(value string) (Name, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return "", errors.New("名前は必須です")
	}

	if len(value) > 100 {
		return "", errors.New("名前は100文字以内です")
	}

	return Name(value), nil
}

func (n Name) String() string {
	return string(n)
}
