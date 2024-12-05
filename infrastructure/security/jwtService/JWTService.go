package jwtService

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	jwtTokenTimeToLive = 60 * 60 * 24 * 2 // 2 дня
)

type JWTService struct {
	secret      string
	validClaims map[string]bool
}

func (j *JWTService) Verify(tokenStr string) bool {
	_, err := j.parse(tokenStr)
	return err == nil
}

func (j *JWTService) FillCtxWithParams(ctx context.Context, tokenStr string) (context.Context, error) {
	token, err := j.parse(tokenStr)
	if err != nil {
		return ctx, nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx, ErrJWTWrongClaims
	}

	// TODO подумать об извлечении и инициализации параметров для контекста в цикле
	userId, err := j.extractClaimAsString(claims, userIdTokenKey)
	if err != nil {
		return ctx, err
	}

	role, err := j.extractClaimAsString(claims, RoleTokenKey)
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, UserIdKey, userId)
	ctx = context.WithValue(ctx, UserRoleKey, role)

	return ctx, nil
}

func (j *JWTService) CreateUserToken(userID int, claims map[string]string) (string, error) {
	if userID <= 0 {
		return "", ErrJWTMissingUserId
	}

	// Время жизни токена от текущего момента + 2 дня
	expirationTime := time.Now().Add(time.Duration(jwtTokenTimeToLive) * time.Second).Unix()

	tokenClaims := jwt.MapClaims{
		userIdTokenKey:     userID,
		expirationTokenKey: expirationTime,
	}

	for key, value := range claims {
		if !j.validClaims[key] {
			return "", ErrJWTInvalidClaims
		}
		tokenClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)

	signedToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", ErrSignJWT
	}

	return signedToken, nil
}

// Парсит токен, проверяя на соответствие HMAC метода подписи
func (j *JWTService) parse(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи токена
		switch token.Method.(type) {
		case *jwt.SigningMethodHMAC:
			return []byte(j.secret), nil
		default:
			return nil, ErrJWTUnsupportedSigningMethod
		}
	})
}

func (j *JWTService) extractClaimAsString(claims jwt.MapClaims, key string) (string, error) {
	value, ok := claims[key]
	if !ok {
		return "", ErrJWTMissingClaim
	}

	valueStr, ok := value.(string)
	if !ok {
		return "", ErrJWTInvalidClaimType
	}

	return valueStr, nil
}
