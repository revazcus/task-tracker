package envRegistryInterface

import "infrastructure/envRegistry"

type EnvRegistry interface {
	GetEnv(key envRegistry.EnvKey) (string, error)
	FindAndSetEnv(key envRegistry.EnvKey) error
	SetEnv(key envRegistry.EnvKey, val string)
}
