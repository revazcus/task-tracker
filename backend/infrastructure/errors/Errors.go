package errors

import "strings"

type Errors struct {
	errors []*Error
}

func NewErrors() *Errors {
	return &Errors{
		errors: make([]*Error, 0),
	}
}

func NewErrorsFrom(errs ...error) *Errors {
	commonErrs := NewErrors()
	commonErrs.AddErrors(errs)
	return commonErrs
}

func (e *Errors) Error() string {
	errStrings := make([]string, 0)
	for _, err := range e.errors {
		errStrings = append(errStrings, err.Error())
	}
	return strings.Join(errStrings, "\n")
}

func (e *Errors) AddNewError(errorCode ErrorCode, msg string) {
	e.AddError(NewError(errorCode, msg))
}

func (e *Errors) AddErrors(errs []error) {
	for _, err := range errs {
		e.AddError(err)
	}
}

func (e *Errors) AddError(err error) {
	switch err := err.(type) {
	case *Error:
		e.errors = append(e.errors, err)
	case *Errors:
		e.errors = append(e.errors, err.errors...)
	default:
		commonErr := NewErrorFrom(err)
		e.errors = append(e.errors, commonErr)
	}
}

func (e *Errors) Contains(err error) bool {
	if err == nil {
		return false
	}

	commonErr, isCommonErr := err.(*Error)
	if !isCommonErr {
		return false
	}

	for _, err := range e.errors {
		if err.Equals(commonErr) {
			return true
		}
	}

	return false
}

func (e *Errors) ContainsByCode(errorCode ErrorCode) bool {
	for _, err := range e.errors {
		if err.Code() == errorCode {
			return true
		}
	}
	return false
}

func (e *Errors) IsEmpty() bool {
	return len(e.errors) == 0
}

func (e *Errors) IsPresent() bool {
	return len(e.errors) > 0
}

func (e *Errors) Size() int {
	return len(e.errors)
}

func (e *Errors) ToArray() []*Error {
	errsCopy := make([]*Error, len(e.errors))
	copy(errsCopy, e.errors)
	return errsCopy
}
