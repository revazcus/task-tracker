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

	address := fmt.Sprintf(":%s", portStr)
	grpcController := grpc.NewUserController(dc.UseCases.UserUseCase)
	userServer := grpc.NewUserServer(address, grpcController, dc.Logger)

	i.UserServer = userServer
	return nil
}
