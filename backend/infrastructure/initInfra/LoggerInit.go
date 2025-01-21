package initInfra

import (
	commonLogger "infrastructure/logger"
	"infrastructure/logger/zapLogger"
)

type LoggerInit struct {
	loggerService *commonLogger.LoggerService
}

func NewLoggerInit() *LoggerInit {
	return &LoggerInit{}
}

func (i *LoggerInit) InitInfra(ic *InfraContainer) error {
	securityContextStr, err := ic.EnvRegistry.GetEnv(SecurityContextEnv)
	if err != nil {
		ic.LogError(err)
		return err
	}

	loggerService := commonLogger.NewLoggerService(ic.stopChan)

	// TODO переписать второй параметр NewZapLogger на securityContext
	zap := zapLogger.NewZapLogger(ic.appID, securityContextStr)
	loggerService.AddLogger("zap", zap)

	ic.LoggerService = loggerService
	ic.Logger = commonLogger.NewLogger(loggerService.GetInputChan())
	return nil
}

func (i *LoggerInit) StartInfra(ic *InfraContainer) error {
	ic.LoggerService.Start()
	return nil
}
