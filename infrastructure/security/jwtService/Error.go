package jwtService

import "errors"

var (
	ErrJWTUnsupportedSigningMethod = errors.New("unsupported signing method")
	ErrJWTMissingUserId            = errors.New("userId is required")
	ErrJWTInvalidClaims            = errors.New("invalid claims provided")
	ErrSignJWT                     = errors.New("failed to sign token")
	ErrJWTWrongClaims              = errors.New("wrong claims")
	ErrJWTMissingClaim             = errors.New("missing claim")
	ErrJWTInvalidClaimType         = errors.New("invalid claim type")
)
