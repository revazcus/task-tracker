package initInfra

import (
	"fmt"
	"infrastructure/restServer"
	restModel "infrastructure/restServer/model"
	"strconv"
)

type RestServerInit struct {
}

func NewRestServerInit() *RestServerInit {
	return &RestServerInit{}
}

func (i *RestServerInit) InitInfra(ic *InfraContainer) error {
	portStr, err := ic.EnvRegistry.GetEnv(ServicePortEnv)
	if err != nil {
		ic.LogError(err)
		return err
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		ic.LogError(err)
		return err
	}

	restConfig := restModel.NewDefaultRestConfig(port)

	ginServer, err := restServer.NewRestServerBuilder().
		ServerConfig(restConfig).
		Logger(ic.Logger).
		JwtService(ic.JWTService).
		Build()

	ic.RestServer = ginServer

	ic.LogInfo(fmt.Sprintf("RestServer initialized Port: %d", restConfig.Port()))

	return nil
}

func (i *RestServerInit) StartInfra(ic *InfraContainer) error {
	return nil
}
