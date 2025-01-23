package initService

import (
	"common/gateways"
	commonUserGateway "common/gateways/user"
)

type GrpcGateways struct {
	BaseGrpcGateway *gateways.BaseGRPCGateway
	UserGrpcGateway *commonUserGateway.UserGateway
}

type GrpcGatewayInit struct {
	dc *DependencyContainer

	grpcGateways  *GrpcGateways
	initFunctions []func(*DependencyContainer) error
}

func NewGrpcGatewayInit(dc *DependencyContainer) *GrpcGatewayInit {
	grpcGateways := &GrpcGateways{}
	return &GrpcGatewayInit{
		dc:           dc,
		grpcGateways: grpcGateways,
		initFunctions: []func(*DependencyContainer) error{
			grpcGateways.initBaseGrpcGateway,
			grpcGateways.initUserGrpcGateway,
		},
	}
}

func (i *GrpcGatewayInit) InitServices() error {
	for _, initFunc := range i.initFunctions {
		if err := initFunc(i.dc); err != nil {
			return err
		}
	}

	i.dc.GrpcGateways = i.grpcGateways
	return nil
}

func (i *GrpcGatewayInit) StartServices() error {
	return i.dc.GrpcGateways.BaseGrpcGateway.Start()
}

func (i *GrpcGateways) initBaseGrpcGateway(dc *DependencyContainer) error {
	userServiceUrl, err := dc.EnvRegistry.GetEnv(GrpcUserServiceUrl)
	if err != nil {
		dc.LogError(err)
		return err
	}

	baseGRPCGateway, err := gateways.NewBuilder().
		Url(userServiceUrl).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.BaseGrpcGateway = baseGRPCGateway
	return nil
}

func (i *GrpcGateways) initUserGrpcGateway(dc *DependencyContainer) error {
	userGateway, err := commonUserGateway.NewBuilder().
		BaseGrpcGateway(i.BaseGrpcGateway).
		Logger(dc.Logger).
		Build()
	if err != nil {
		dc.LogError(err)
		return err
	}

	i.UserGrpcGateway = userGateway
	return nil
}
