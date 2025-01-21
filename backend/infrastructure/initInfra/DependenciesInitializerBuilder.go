package initInfra

import "infrastructure/errors"

type DependenciesInitializerBuilder struct {
	appId             string
	infraInitChain    []InfraInit
	servicesInitChain []ServicesInit
	infraSetter       InfraSetter

	errors *errors.Errors
}

func NewDependenciesInitializer() *DependenciesInitializerBuilder {
	return &DependenciesInitializerBuilder{
		errors: errors.NewErrors(),
	}
}

func (b *DependenciesInitializerBuilder) AppId(appId string) *DependenciesInitializerBuilder {
	b.appId = appId
	return b
}

func (b *DependenciesInitializerBuilder) InfraInitChain(infraInitChain []InfraInit) *DependenciesInitializerBuilder {
	b.infraInitChain = infraInitChain
	return b
}

func (b *DependenciesInitializerBuilder) ServicesInitChain(serviceInitChain []ServicesInit) *DependenciesInitializerBuilder {
	b.servicesInitChain = serviceInitChain
	return b
}

func (b *DependenciesInitializerBuilder) InfraSetter(infraSetter InfraSetter) *DependenciesInitializerBuilder {
	b.infraSetter = infraSetter
	return b
}

func (b *DependenciesInitializerBuilder) Init() error {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return b.errors
	}

	initializer := b.createInitializer()
	return initializer.InitDependencies()
}

func (b *DependenciesInitializerBuilder) checkRequiredFields() {
	if b.appId == "" {
		b.errors.AddError(ErrAppIdIsRequired)
	}
	if b.infraInitChain == nil {
		b.errors.AddError(ErrInfraInitChainIsRequired)
	}
	if b.servicesInitChain == nil {
		b.errors.AddError(ErrServicesInitChainIsRequired)
	}
	if b.infraSetter == nil {
		b.errors.AddError(ErrInfraSetterIsRequired)
	}
}

func (b *DependenciesInitializerBuilder) createInitializer() *DependenciesInitializer {
	return &DependenciesInitializer{
		appId:             b.appId,
		infraInitChain:    b.infraInitChain,
		servicesInitChain: b.servicesInitChain,
		infraSetter:       b.infraSetter,
	}
}
