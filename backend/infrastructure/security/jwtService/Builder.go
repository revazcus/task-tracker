package jwtService

import "infrastructure/errors"

type Builder struct {
	service *JWTService
}

func NewBuilder() *Builder {
	return &Builder{
		service: &JWTService{},
	}
}

func (b *Builder) Secret(secret string) *Builder {
	b.service.secret = secret
	return b
}

func (b *Builder) ValidClaims(validClaims map[string]bool) *Builder {
	b.service.validClaims = validClaims
	return b
}

func (b *Builder) Build() (*JWTService, error) {
	err := b.checkRequiredFields()
	if err != nil {
		return nil, err
	}

	b.fillDefaultFields()

	return b.service, nil
}

func (b *Builder) checkRequiredFields() error {
	if b.service.secret == "" {
		return errors.NewError("SYS", "JwtServiceBuilder: Secret is required")
	}
	return nil
}

// Перечисляем валидные ключи
func (b *Builder) fillDefaultFields() {
	if b.service.validClaims == nil {
		b.service.validClaims = map[string]bool{
			RoleTokenKey:       true,
			ScopeTokenKey:      true,
			expirationTokenKey: true,
		}
	}
}
