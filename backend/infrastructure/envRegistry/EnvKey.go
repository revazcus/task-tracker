package envRegistry

type EnvKey string

func (e EnvKey) String() string {
	return string(e)
}
