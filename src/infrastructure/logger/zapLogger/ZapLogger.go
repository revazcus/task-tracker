package zapLogger

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	logModel "task-tracker/infrastructure/logger/model"
)

type CtxLoggerKey string

const (
	timeTag = "timestamp"
)

type Logger struct {
	*zap.Logger
	appId string
	env   string
}

func NewZapLogger(appId, env string) *Logger {
	config := getEncoderConfig()
	coreConsole := zapcore.NewCore(zapcore.NewJSONEncoder(config), os.Stdout, getAllLvlFunc())
	zapLogger := zap.New(zapcore.NewTee(coreConsole))

	return &Logger{
		Logger: zapLogger,
		appId:  appId,
		env:    env,
	}
}

func (l *Logger) SendMessage(logData *logModel.LogData) {
	if logData.Ctx == nil {
		logData.Ctx = context.Background()
	}

	appId, ok := logData.Ctx.Value(logModel.AppId).(string)
	if !ok || appId == "" {
		appId = l.appId
	}

	env, ok := logData.Ctx.Value(logModel.EnvName).(string)
	if !ok || env == "" {
		env = l.env
	}

	fields := []zapcore.Field{
		zap.String("service_name", appId),
		zap.String("env", env),
	}

	resFields := l.getPayloadFields(logData)
	fields = append(fields, resFields...)

	switch logData.Lvl {
	case logModel.DebugLvl:
		l.Debug(logData.Msg, fields...)
	case logModel.InfoLvl:
		l.Info(logData.Msg, fields...)
	case logModel.WarnLvl:
		l.Warn(logData.Msg, fields...)
	case logModel.ErrorLvl:
		l.Error(logData.Msg, fields...)
	case logModel.DPanicLvl:
		l.DPanic(logData.Msg, fields...)
	case logModel.PanicLvl:
		l.Panic(logData.Msg, fields...)
	case logModel.FatalLvl:
		l.Fatal(logData.Msg, fields...)
	}
}

func getEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = timeTag
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	return config
}

func getAllLvlFunc() zap.LevelEnablerFunc {
	return func(lvl zapcore.Level) bool {
		return true
	}
}

func (l *Logger) getPayloadFields(logData *logModel.LogData) []zap.Field {
	var resFields []zap.Field
	resFields = append(resFields, zap.Namespace("payload"))
	for _, field := range logData.Fields {
		if field.Integer != 0 {
			resFields = append(resFields, zap.Int(field.Key, field.Integer))
		}
		if field.String != "" {
			resFields = append(resFields, zap.String(field.Key, field.String))
		}
		if field.Float != 0.0 {
			resFields = append(resFields, zap.Float64(field.Key, field.Float))
		}
	}
	return resFields
}
