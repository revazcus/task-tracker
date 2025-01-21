package envRegistry

import "os"

type EnvRegistry struct {
	envs map[EnvKey]string
}

func NewEnvRegistry() *EnvRegistry {
	return &EnvRegistry{
		envs: make(map[EnvKey]string),
	}
}

func (r *EnvRegistry) GetEnv(key EnvKey) (string, error) {
	env, exists := r.envs[key]
	if !exists {
		return "", ErrEnvInRegistryNotFound(key)
	}
	return env, nil
}

func (r *EnvRegistry) FindAndSetEnv(key EnvKey) error {
	env := os.Getenv(key.String())
	if env == "" {
		return ErrEnvInOSNotFound(key)
	}
	r.SetEnv(key, env)
	return nil
}

func (r *EnvRegistry) SetEnv(key EnvKey, val string) {
	r.envs[key] = val
}
