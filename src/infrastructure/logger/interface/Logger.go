package loggerInterface

type Logger interface {
	LogInfo(method string, path string, requestBody string)
	LogError(err error, method string, path string, requestBody string)
}
