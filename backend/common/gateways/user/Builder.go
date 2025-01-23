package commonUserGateway

import (
	"common/gateways"
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
)

type Builder struct {
	userGateway *UserGateway
	errors      *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		userGateway: &UserGateway{},
		errors:      errors.NewErrors(),
	}
}

func (b *Builder) BaseGrpcGateway(baseGrpcGateway *gateways.BaseGRPCGateway) *Builder {
	b.userGateway.baseGRPC = baseGrpcGateway
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.userGateway.logger = logger
	return b
}

func (b *Builder) Build() (*UserGateway, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.userGateway, nil
}

func (b *Builder) checkRequiredFields() {
	if b.userGateway.baseGRPC == nil {
		b.errors.AddError(errors.NewError("SYS", "UserGatewayBuilder: BaseGRPC is required"))
	}
	if b.userGateway.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "UserGatewayBuilder: Logger is required"))
	}
}
