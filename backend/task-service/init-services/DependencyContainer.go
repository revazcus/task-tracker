package initService

import (
	"infrastructure/initInfra"
	saramaClient "infrastructure/kafka"
)

type DependencyContainer struct {
	*initInfra.InfraContainer

	Repositories    *Repositories
	UseCases        *UseCases
	RestControllers *RestControllers
	GrpcGateways    *GrpcGateways

	// TODO разобраться почему работает именно так, а не через интерфейс
	KafkaClient *saramaClient.SaramaClient
}

func (dc *DependencyContainer) SetInfraContainer(infraContainer *initInfra.InfraContainer) {
	dc.InfraContainer = infraContainer
}

func GetInfraInitChains() []initInfra.InfraInit {
	return []initInfra.InfraInit{
		initInfra.NewEnvironmentsInit(),
		initInfra.NewSecurityContextInit(),
		initInfra.NewLoggerInit(),
		initInfra.NewMongoInit(),
		initInfra.NewJwtServiceInit(),
		initInfra.NewRestServerInit(),
	}
}

func GetServicesInitChains(dc *DependencyContainer) []initInfra.ServicesInit {
	return []initInfra.ServicesInit{
		NewEnvironmentsInit(dc),
		NewRepositoriesInit(dc),
		// TODO разобраться, почему именно в таком порядке инициализации кафка поднимается, а в другом нет
		NewKafkaInit(dc),
		// TODO разобраться, почему именно в таком порядке инициализации GRPC поднимается, а в другом нет
		NewGrpcGatewayInit(dc),
		NewUseCasesInit(dc),
		NewRestControllersInit(dc),
		NewRestRoutingInit(dc),
	}
}
