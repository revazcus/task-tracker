package userGatewayInterface

import (
	userObject "common/domainObject/shortUser"
	"context"
)

type UserGateway interface {
	GetUserById(ctx context.Context, Id string) (*userObject.ShortUser, error)
}
