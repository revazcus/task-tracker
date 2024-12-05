package jwtServiceInterface

import "context"

type JWTService interface {
	Verify(tokenStr string) bool
	FillCtxWithParams(ctx context.Context, tokenStr string) (context.Context, error)
	CreateUserToken(userID int, claims map[string]string) (string, error)
}
