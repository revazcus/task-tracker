package restModel

type RestServerConfig struct {
	HttpWriteTimeout          int
	HttpReadTimeout           int
	HttpServerShutdownTimeout int
	port                      int
}

func NewDefaultRestConfig(port int) *RestServerConfig {
	return &RestServerConfig{
		HttpWriteTimeout:          30,
		HttpReadTimeout:           30,
		HttpServerShutdownTimeout: 10,
		port:                      port,
	}
}

func (c *RestServerConfig) Port() int {
	return c.port
}
