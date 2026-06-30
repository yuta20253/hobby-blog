package user

import (
	"errors"
	"net/mail"
	"strings"
)

type Email string

func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return "", errors.New("メールアドレスは必須です")
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return "", errors.New("メールアドレスの形式が正しくありません")
	}

	return Email(value), nil
}

func (e Email) String() string {
	return string(e)
}
