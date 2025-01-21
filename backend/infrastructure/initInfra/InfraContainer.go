package initInfra

import (
	"context"
	"fmt"
	"infrastructure/envRegistry"
	envRegistryInterface "infrastructure/envRegistry/interface"
	commonLogger "infrastructure/logger"
	loggerInterface "infrastructure/logger/interface"
	mongoRepo "infrastructure/mongo"
	restServerInterface "infrastructure/restServer/interface"
	jwtServiceInterface "infrastructure/security/jwtService/interface"
	"time"
)

const closeTimeout = 500 * time.Millisecond

type InfraInit interface {
	InitInfra(ic *InfraContainer) error
	StartInfra(ic *InfraContainer) error
}

type InfraContainer struct {
	SecurityContext SecurityContext

	Logger        loggerInterface.Logger
	LoggerService *commonLogger.LoggerService
	MongoRepo     *mongoRepo.MongoRepo
	RestServer    restServerInterface.Server
	JWTService    jwtServiceInterface.JWTService
	EnvRegistry   envRegistryInterface.EnvRegistry

	stopChan chan struct{}
	appID    string

	TraceEnable bool
	KafkaEnable bool
}

func NewInfraContainer(appID string) *InfraContainer {
	infra := &InfraContainer{
		appID:       appID,
		stopChan:    make(chan struct{}),
		EnvRegistry: envRegistry.NewEnvRegistry(),
	}
	return infra
}

func (ic *InfraContainer) AppID() string {
	return ic.appID
}

func (ic *InfraContainer) LogInfo(msg string) {
	if ic.Logger == nil {
		fmt.Println(msg)
	} else {
		ic.Logger.Info(context.Background(), msg)
	}
}

func (ic *InfraContainer) LogError(err error) {
	if ic.Logger == nil {
		fmt.Printf("error: %v\n", err.Error())
	} else {
		ic.Logger.Error(context.Background(), err)
	}
}

func (ic *InfraContainer) LogWarning(msg string) {
	if ic.Logger == nil {
		fmt.Printf("warning: %s\n", msg)
	} else {
		ic.Logger.Warning(context.Background(), msg)
	}
}

func (ic *InfraContainer) StopChan() chan struct{} {
	return ic.stopChan
}

func (ic *InfraContainer) CloseInfra() {
	close(ic.stopChan)
	time.Sleep(closeTimeout)
}
