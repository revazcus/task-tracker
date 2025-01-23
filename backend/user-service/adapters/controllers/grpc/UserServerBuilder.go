package grpc

import (
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
)

type UserServerBuilder struct {
	port           string
	userController *UserController
	logger         loggerInterface.Logger

	errors *errors.Errors
}

func NewUserServerBuilder() *UserServerBuilder {
	return &UserServerBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *UserServerBuilder) Port(port string) *UserServerBuilder {
	b.port = port
	return b
}

func (b *UserServerBuilder) UserController(userController *UserController) *UserServerBuilder {
	b.userController = userController
	return b
}

func (b *UserServerBuilder) Logger(logger loggerInterface.Logger) *UserServerBuilder {
	b.logger = logger
	return b
}

func (b *UserServerBuilder) Build() (*UserServer, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.createFromBuilder(), nil
}

func (b *UserServerBuilder) checkRequiredFields() {
	if b.port == "" {
		b.errors.AddError(errors.NewError("SYS", "UserServerBuilder: Port is required"))
	}
	if b.userController == nil {
		b.errors.AddError(errors.NewError("SYS", "UserServerBuilder: UserController is required"))
	}
	if b.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "UserServerBuilder: Logger is required"))
	}
}

func (b *UserServerBuilder) createFromBuilder() *UserServer {
	return &UserServer{
		port:           b.port,
		userController: b.userController,
		logger:         b.logger,
	}
}
