package initInfra

import "infrastructure/errors"

var (
	ErrAppIdIsRequired             = errors.NewError("SYS", "AppId is required")
	ErrInfraInitChainIsRequired    = errors.NewError("SYS", "InfraInitChain is required")
	ErrServicesInitChainIsRequired = errors.NewError("SYS", "ServicesInitChain is required")
	ErrInfraSetterIsRequired       = errors.NewError("SYS", "InfraSetter is required")
)
