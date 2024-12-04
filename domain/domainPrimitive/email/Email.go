package emailPrimitive

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type Email string

func EmailFrom(email string) (Email, error) {
	if email == "" {
		return "", ErrEmailIsEmpty
	}

	// Обрезаем пробелы и приводим к нижнему регистру
	trimmedEmail := strings.TrimSpace(email)
	trimmedEmail = strings.ToLower(trimmedEmail)

	validate := validator.New()

	// Проверяем email через библиотеку
	err := validate.Var(trimmedEmail, "required,email")
	if err != nil {
		return "", ErrEmailIsInvalid
	}

	return Email(trimmedEmail), nil
}
