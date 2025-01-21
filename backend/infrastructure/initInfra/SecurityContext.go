package initInfra

import (
	"fmt"
	"infrastructure/errors"
)

type SecurityContext string

type SecurityContextEnum map[string]SecurityContext

const (
	localSecurityContext = "local"
	devSecurityContext   = "dev"
	qaSecurityContext    = "qa"
	prodSecurityContext  = "prod"
)

var SecurityContexts = SecurityContextEnum{
	localSecurityContext: localSecurityContext,
	devSecurityContext:   devSecurityContext,
	qaSecurityContext:    qaSecurityContext,
	prodSecurityContext:  prodSecurityContext,
}

func (e SecurityContextEnum) Local() SecurityContext {
	return e[localSecurityContext]
}

func (e SecurityContextEnum) Dev() SecurityContext {
	return e[devSecurityContext]
}

func (e SecurityContextEnum) QA() SecurityContext {
	return e[qaSecurityContext]
}

func (e SecurityContextEnum) Prod() SecurityContext {
	return e[prodSecurityContext]
}

func (e SecurityContextEnum) Of(securityContextStr string) (SecurityContext, error) {
	stage, ok := e[securityContextStr]
	if !ok {
		return "", errors.NewError("SYS", fmt.Sprintf("Unsupported SecurityContext = '%s'", securityContextStr))
	}
	return stage, nil
}
