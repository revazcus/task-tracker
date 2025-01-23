package gateways

import (
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
)

type Builder struct {
	grpcGateway *BaseGRPCGateway
	errors      *errors.Errors
}

func NewBuilder() *Builder {
	return &Builder{
		grpcGateway: &BaseGRPCGateway{},
		errors:      errors.NewErrors(),
	}
}

func (b *Builder) Url(url string) *Builder {
	b.grpcGateway.url = url
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.grpcGateway.logger = logger
	return b
}

func (b *Builder) Build() (*BaseGRPCGateway, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.grpcGateway, nil
}

func (b *Builder) checkRequiredFields() {
	if b.grpcGateway.url == "" {
		b.errors.AddError(errors.NewError("SYS", "BaseGrpcGatewayBuilder: Url is required"))
	}
	if b.grpcGateway.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "BaseGrpcGatewayBuilder: Logger is required"))
	}
}
