package initInfra

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DependenciesInitializer struct {
	appId             string
	infraInitChain    []InfraInit
	servicesInitChain []ServicesInit
	infraSetter       InfraSetter

	infraContainer *InfraContainer
}

type InfraSetter interface {
	SetInfraContainer(infraContainer *InfraContainer)
}

type ServicesInit interface {
	InitServices() error
	StartServices() error
}

func (i *DependenciesInitializer) InitDependencies() error {
	if err := i.initInfraContainer(i.infraSetter); err != nil {
		i.infraContainer.CloseInfra()
		return err
	}

	if err := i.initServices(); err != nil {
		i.infraContainer.CloseInfra()
		return err
	}

	i.ListenStopSig()

	time.Sleep(200 * time.Millisecond)
	i.infraContainer.LogInfo(fmt.Sprintf("%s service started", i.appId))
	return nil
}

func (i *DependenciesInitializer) initInfraContainer(infraSetter InfraSetter) error {
	infraContainer := NewInfraContainer(i.appId)
	infraSetter.SetInfraContainer(infraContainer)
	i.infraContainer = infraContainer

	for _, initChain := range i.infraInitChain {
		if err := initChain.InitInfra(infraContainer); err != nil {
			infraContainer.LogError(err)
			return err
		}
	}

	for _, initChain := range i.infraInitChain {
		if err := initChain.StartInfra(infraContainer); err != nil {
			infraContainer.LogError(err)
			return err
		}
	}

	return nil
}

func (i *DependenciesInitializer) initServices() error {
	for _, serviceInitChain := range i.servicesInitChain {
		if err := serviceInitChain.InitServices(); err != nil {
			i.infraContainer.LogError(err)
			return err
		}
	}

	for _, serviceInitChain := range i.servicesInitChain {
		if err := serviceInitChain.StartServices(); err != nil {
			i.infraContainer.LogError(err)
			return err
		}
	}

	return nil
}

func (i *DependenciesInitializer) ListenStopSig() {
	go func() {
		sigHandler := make(chan os.Signal, 1)
		signal.Notify(sigHandler, syscall.SIGTERM, syscall.SIGINT)

		select {
		case <-sigHandler:
			i.infraContainer.LogInfo(fmt.Sprintf("%s service stopped", i.appId))
			time.Sleep(500 * time.Microsecond)

			i.infraContainer.CloseInfra()
		}
	}()
}
