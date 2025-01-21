package initInfra

import (
	"infrastructure/envRegistry"
	"os"
)

const (
	SecurityContextEnv envRegistry.EnvKey = "SECURITY_CONTEXT"
	JWTSecretKeyEnv    envRegistry.EnvKey = "JWT_SECRET_KEY"
	ServicePortEnv     envRegistry.EnvKey = "SERVICE_PORT"

	MongoURIEnv    envRegistry.EnvKey = "MONGO_URI"
	MongoDBNameEnv envRegistry.EnvKey = "MONGO_DB_NAME"
)

type EnvironmentsInit struct {
}

func NewEnvironmentsInit() *EnvironmentsInit {
	return &EnvironmentsInit{}
}

func (i *EnvironmentsInit) InitInfra(ic *InfraContainer) error {
	i.findAndSetEnv(JWTSecretKeyEnv, ic)
	i.findAndSetEnv(MongoURIEnv, ic)
	i.findAndSetEnv(MongoDBNameEnv, ic)
	i.findAndSetEnv(ServicePortEnv, ic)
	i.findAndSetEnv(SecurityContextEnv, ic)
	return nil
}

func (i *EnvironmentsInit) StartInfra(ic *InfraContainer) error {
	return nil
}

func (i *EnvironmentsInit) findAndSetEnv(envName envRegistry.EnvKey, ic *InfraContainer) {
	envValue := os.Getenv(envName.String())
	ic.EnvRegistry.SetEnv(envName, envValue)
}
