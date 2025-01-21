package usernamePrimitive

import (
	"regexp"
)

// FORMAT доступны только буквы, цифры, точки и подчёркивания (не может начинаться/закачиваться точкой или подчёркиванием)
const FORMAT = "^[a-zA-Z0-9]+$"

type Username string

func UsernameFrom(username string) (Username, error) {
	if username == "" {
		return "", ErrUsernameIsEmpty
	}

	length := len(username)
	if length < 3 || length > 24 {
		return "", ErrUsernameLength
	}

	validFormat := regexp.MustCompile(FORMAT)
	if !validFormat.MatchString(username) {
		return "", ErrUsernameIsInvalid
	}

	return Username(username), nil

}

func (u Username) Verify(username string) bool {
	return string(u) == username
}

func (u Username) String() string {
	return string(u)
}
