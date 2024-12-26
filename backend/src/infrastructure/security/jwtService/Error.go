package jwtService

import "task-tracker/infrastructure/errors"

var (
	ErrJWTUnsupportedSigningMethod = errors.NewErrorWithLvl("SYS", "Unsupported signing method", errors.Levels.Info())
	ErrJWTMissingUserId            = errors.NewErrorWithLvl("SYS", "UserId is required", errors.Levels.Info())
	ErrJWTInvalidClaims            = errors.NewErrorWithLvl("SYS", "Invalid claims provided", errors.Levels.Info())
	ErrSignJWT                     = errors.NewErrorWithLvl("SYS", "Failed to sign token", errors.Levels.Info())
	ErrJWTWrongClaims              = errors.NewErrorWithLvl("SYS", "Wrong claims", errors.Levels.Info())
	ErrJWTMissingClaim             = errors.NewErrorWithLvl("SYS", "Missing claim", errors.Levels.Info())
	ErrJWTInvalidClaimType         = errors.NewErrorWithLvl("SYS", "Invalid claim type", errors.Levels.Info())
)
