package restServer

import (
	"infrastructure/errors"
	loggerInterface "infrastructure/logger/interface"
	restModel "infrastructure/restServer/model"
	jwtServiceInterface "infrastructure/security/jwtService/interface"
)

type Builder struct {
	serverConfig *restModel.RestServerConfig
	logger       loggerInterface.Logger
	jwtService   jwtServiceInterface.JWTService

	errors *errors.Errors
}

func NewRestServerBuilder() *Builder {
	return &Builder{
		errors: errors.NewErrors(),
	}
}

func (b *Builder) ServerConfig(serverConfig *restModel.RestServerConfig) *Builder {
	b.serverConfig = serverConfig
	return b
}

func (b *Builder) Logger(logger loggerInterface.Logger) *Builder {
	b.logger = logger
	return b
}

func (b *Builder) JwtService(jwtService jwtServiceInterface.JWTService) *Builder {
	b.jwtService = jwtService
	return b
}

func (b *Builder) Build() (*RestServer, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}
	return b.createFromBuilder(), nil
}

func (b *Builder) checkRequiredFields() {
	if b.serverConfig == nil {
		b.errors.AddError(errors.NewError("SYS", "RestServerBuilder: ServerConfig is required"))
	}
	if b.logger == nil {
		b.errors.AddError(errors.NewError("SYS", "RestServerBuilder: Logger is required"))
	}
	if b.jwtService == nil {
		b.errors.AddError(errors.NewError("SYS", "RestServerBuilder: JwtService is required"))
	}
}

func (b *Builder) createFromBuilder() *RestServer {
	restServer := &RestServer{
		serverConfig: b.serverConfig,
		logger:       b.logger,
		jwtService:   b.jwtService,
	}
	restServer.createHttpServer()
	return restServer
}
