package errors

import "fmt"

const UnknownErrorCode = ErrorCode("unknown_error")

type ErrorCode string

func (c ErrorCode) String() string {
	return string(c)
}

type Error struct {
	lvl  ErrorLvl
	code ErrorCode
	msg  string
}

func NewError(code ErrorCode, msg string) *Error {
	return &Error{
		lvl:  Levels.Error(),
		code: code,
		msg:  msg,
	}
}

func NewErrorWithLvl(code ErrorCode, msg string, lvl ErrorLvl) *Error {
	return &Error{
		lvl:  lvl,
		code: code,
		msg:  msg,
	}
}

func NewErrorFrom(err error) *Error {
	return CastOrWrap(err, UnknownErrorCode)
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.msg)
}

func (e *Error) Lvl() ErrorLvl {
	return e.lvl
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Message() string {
	return e.msg
}

func (e *Error) Equals(err *Error) bool {
	if err == nil {
		return false
	}
	return e.Code() == err.Code() && e.Message() == err.Message()
}
