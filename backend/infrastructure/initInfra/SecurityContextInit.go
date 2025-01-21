package initInfra

type SecurityContextInit struct {
}

func NewSecurityContextInit() *SecurityContextInit {
	return &SecurityContextInit{}
}

func (i *SecurityContextInit) InitInfra(ic *InfraContainer) error {
	securityContextStr, err := ic.EnvRegistry.GetEnv(SecurityContextEnv)
	if err != nil {
		ic.LogError(err)
		return err
	}

	securityContext, err := SecurityContexts.Of(securityContextStr)
	if err != nil {
		ic.LogError(err)
		return err
	}

	ic.SecurityContext = securityContext
	return nil
}

func (i *SecurityContextInit) StartInfra(ic *InfraContainer) error {
	return nil
}
