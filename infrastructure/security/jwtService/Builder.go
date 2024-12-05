package jwtService

import "task-tracker/infrastructure/errors"

type JwtServiceBuilder struct {
	service *JWTService
}

func NewBuilder() *JwtServiceBuilder {
	return &JwtServiceBuilder{
		service: &JWTService{},
	}
}

func (b *JwtServiceBuilder) Secret(secret string) *JwtServiceBuilder {
	b.service.secret = secret
	return b
}

func (b *JwtServiceBuilder) ValidClaims(validClaims map[string]bool) *JwtServiceBuilder {
	b.service.validClaims = validClaims
	return b
}

func (b *JwtServiceBuilder) Build() (*JWTService, error) {
	err := b.checkRequiredFields()
	if err != nil {
		return nil, err
	}

	b.fillDefaultFields()

	return b.service, nil
}

func (b *JwtServiceBuilder) checkRequiredFields() error {
	if b.service.secret == "" {
		return errors.ErrSecretIsRequired
	}
	return nil
}

func (b *JwtServiceBuilder) fillDefaultFields() {
	if b.service.validClaims == nil {
		b.service.validClaims = map[string]bool{
			RoleTokenKey:       true,
			ScopeTokenKey:      true,
			expirationTokenKey: true,
		}
	}
}
