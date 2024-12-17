package errors

import "fmt"

var (
	UnsupportedLvlErrorCode ErrorCode = "unsupported_error_lvl"
)

func ErrUnsupportedLevel(lvlStr string) error {
	errMessage := fmt.Sprintf("Unsuppported ErrorLvl = `%s`", lvlStr)
	return NewError(UnsupportedLvlErrorCode, errMessage)
}

type ErrorLvl string

func (l ErrorLvl) String() string {
	return string(l)
}

const (
	infoErrorLvl       = "info"
	warnErrorLvl       = "warn"
	errorErrorLvl      = "error"
	criticalErrorLevel = "critical"
)

type ErrorLvlEnum map[string]ErrorLvl

var Levels = ErrorLvlEnum{
	infoErrorLvl:       infoErrorLvl,
	warnErrorLvl:       warnErrorLvl,
	errorErrorLvl:      errorErrorLvl,
	criticalErrorLevel: criticalErrorLevel,
}

func (e ErrorLvlEnum) Info() ErrorLvl {
	return e[infoErrorLvl]
}

func (e ErrorLvlEnum) Warn() ErrorLvl {
	return e[warnErrorLvl]
}

func (e ErrorLvlEnum) Error() ErrorLvl {
	return e[errorErrorLvl]
}

func (e ErrorLvlEnum) Critical() ErrorLvl {
	return e[criticalErrorLevel]
}

func (e ErrorLvlEnum) Of(lvlCode string) (ErrorLvl, error) {
	value, ok := e[lvlCode]
	if !ok {
		return "", ErrUnsupportedLevel(lvlCode)
	}
	return value, nil
}
