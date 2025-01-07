package userGatewayInterface

import (
	"context"
	userObject "github.com/revazcus/task-tracker/backend/common/domainObject/shortUser"
)

type UserGateway interface {
	GetUserById(ctx context.Context, Id string) (*userObject.ShortUser, error)
}
