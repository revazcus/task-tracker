package commonUserGateway

import (
	userObject "common/domainObject/shortUser"
	emailPrimitive "common/domainPrimitive/email"
	idPrimitive "common/domainPrimitive/id"
	profilePrimitive "common/domainPrimitive/profile"
	"common/gateways"
	"common/gateways/user/protoModel"
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
)

type UserGateway struct {
	baseGRPC   *gateways.BaseGRPCGateway
	userClient protoModel.UserServiceClient
	logger     loggerInterface.Logger
}

func (g *UserGateway) GetUserById(ctx context.Context, Id string) (*userObject.ShortUser, error) {
	user, err := g.getUserClient().GetUser(ctx, &protoModel.UserRequest{Id: Id})
	if err != nil {
		grpcStatus, _ := status.FromError(err)
		g.logger.Error(ctx, fmt.Errorf("failed to get user by id: %s, grpc status: %s", err, grpcStatus.Message()))
		return nil, errors.NewError("SYS", fmt.Sprintf("failed to get user by id: %e", err))
	}
	return g.fillShortUserFromProto(user)
}

func (g *UserGateway) fillShortUserFromProto(user *protoModel.User) (*userObject.ShortUser, error) {
	id, err := idPrimitive.EntityIdFrom(user.Id)
	if err != nil {
		return nil, err
	}

	email, err := emailPrimitive.EmailFrom(user.Email)
	if err != nil {
		return nil, err
	}

	profile, err := profilePrimitive.NewProfile(user.FirstName, user.LastName)
	if err != nil {
		return nil, err
	}

	return userObject.NewShortUser(&id, &email, profile), nil
}

func (g *UserGateway) getUserClient() protoModel.UserServiceClient {
	if g.userClient == nil {
		g.userClient = protoModel.NewUserServiceClient(g.baseGRPC.Connection())
	}
	return g.userClient
}
