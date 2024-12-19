package logModel

import (
	"context"
	"fmt"
)

type LogLvl int

const (
	DebugLvl LogLvl = iota - 1
	InfoLvl
	WarnLvl
	ErrorLvl
	DPanicLvl
	PanicLvl
	FatalLvl
)

func (l LogLvl) String() string {
	switch l {
	case DebugLvl:
		return "debug"
	case InfoLvl:
		return "info"
	case WarnLvl:
		return "warn"
	case ErrorLvl:
		return "error"
	case DPanicLvl:
		return "dpanic"
	case PanicLvl:
		return "panic"
	case FatalLvl:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

const (
	FieldErrKey       = "error"
	FieldComponentKey = "component"
	FieldFilenameKey  = "filename"
)

type LogData struct {
	Ctx    context.Context
	Msg    string
	Fields []*LogField
	Lvl    LogLvl
}

type LogField struct {
	Key     string
	Integer int
	Float   float64
	String  string
	Object  interface{}
}
