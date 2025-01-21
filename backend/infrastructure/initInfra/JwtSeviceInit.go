package initInfra

import "infrastructure/security/jwtService"

type JwtServiceInit struct {
}

func NewJwtServiceInit() *JwtServiceInit {
	return &JwtServiceInit{}
}

func (i *JwtServiceInit) InitInfra(ic *InfraContainer) error {
	jwtSecret, err := ic.EnvRegistry.GetEnv(JWTSecretKeyEnv)
	if err != nil {
		ic.LogError(err)
		return err
	}

	jwtService, err := jwtService.NewBuilder().Secret(jwtSecret).Build()
	if err != nil {
		ic.LogError(err)
		return err
	}

	ic.JWTService = jwtService
	return nil
}

func (i *JwtServiceInit) StartInfra(ic *InfraContainer) error {
	return nil
}
