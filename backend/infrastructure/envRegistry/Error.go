package envRegistry

import (
	"fmt"
	"infrastructure/errors"
)

func ErrEnvInOSNotFound(key EnvKey) error {
	errMsg := fmt.Sprintf("`%s` not found in OS environments", key)
	return errors.NewError("SYS", errMsg)
}

func ErrEnvInRegistryNotFound(key EnvKey) error {
	errMsg := fmt.Sprintf("`%s` not found in registry", key)
	return errors.NewError("SYS", errMsg)
}

func ErrInvalidEnvTypeValue(key EnvKey, cause error) error {
	errMsg := fmt.Sprintf("unable to cast env value to bool type, key: %q, cause: %q", key, cause)
	return errors.NewError("SYS", errMsg)
}
