package passwordPrimitive

import (
	"golang.org/x/crypto/bcrypt"
)

// Password доменный примитив, инкапсулирующий хеш пароля
type Password string

func PasswordFrom(password string) (*Password, error) {
	if len(password) < 8 {
		return nil, ErrPasswordLength
	}

	// Хеширование пароля с использованием bcrypt (соль в комплекте)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	pass := Password(hash)

	return &pass, nil
}

func (p *Password) ChangePassword(oldPassword, newPassword string) (*Password, error) {
	if !p.Verify(oldPassword) {
		return nil, ErrPasswordIsWrong
	}

	updatedPassword, err := PasswordFrom(newPassword)
	if err != nil {
		return nil, err
	}

	return updatedPassword, nil
}

// Verify проверяет совпадения хешей
func (p *Password) Verify(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(*p), []byte(password)) == nil
}
