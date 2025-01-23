package initService

import (
	"fmt"
	"user-service/adapters/controllers/grpc"
)

type GrpcServer struct {
	UserServer *grpc.UserServer
}

type GrpcServerInit struct {
	dc *DependencyContainer

	grpcServers   *GrpcServer
	initFunctions []func(*DependencyContainer) error
}

func NewGrpcServerInit(dc *DependencyContainer) *GrpcServerInit {
	grpcServer := &GrpcServer{}
	return &GrpcServerInit{
		dc:          dc,
		grpcServers: grpcServer,
		initFunctions: []func(container *DependencyContainer) error{
			grpcServer.initUserServer,
		},
	}
}

func (i *GrpcServerInit) InitServices() error {
	for _, initFunc := range i.initFunctions {
		if err := initFunc(i.dc); err != nil {
			return err
		}
	}

	i.dc.GrpcServers = i.grpcServers
	return nil
}

func (i *GrpcServerInit) StartServices() error {
	return i.dc.GrpcServers.UserServer.Start()
}

func (i *GrpcServer) initUserServer(dc *DependencyContainer) error {
	portStr, err := dc.EnvRegistry.GetEnv(GrpcPort)
	if err != nil {
		dc.LogError(err)
		return err
	}

	port := fmt.Sprintf(":%s", portStr)
	userController := grpc.NewUserController(dc.UseCases.UserUseCase)
	userServer, err := grpc.NewUserServerBuilder().
		Port(port).
		UserController(userController).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.UserServer = userServer
	return nil
}
