package errors

import (
	"errors"
	"fmt"
)

type BusinessError struct {
	doRollBack    bool
	originalError error
	message       string
}

func (e *BusinessError) DoRollBack() bool {
	return e.doRollBack
}

func (e *BusinessError) OriginalError() error {
	return e.originalError
}

func (e *BusinessError) Message() string {
	return e.message
}

func (e *BusinessError) Error() string {
	return e.Message()
}

func NewBusinessError(format string, a ...interface{}) *BusinessError {
	return &BusinessError{true, nil, fmt.Sprintf(format, a...)}
}

func NewBusinessErrorf(originalError error, format string, a ...interface{}) *BusinessError {
	return &BusinessError{true, originalError, fmt.Sprintf(format, a...)}
}

func NewNoRollBackBusinessError(format string, a ...interface{}) *BusinessError {
	return &BusinessError{false, nil, fmt.Sprintf(format, a...)}
}

func NewNoRollBackBusinessErrorf(originalError error, format string, a ...interface{}) *BusinessError {
	return &BusinessError{false, originalError, fmt.Sprintf(format, a...)}
}

func IsBusinessError(err error) bool {
	_, ok := AsBusinessError(err)
	return ok
}

func IsSystemError(err error) bool {
	_, ok := AsBusinessError(err)
	return err != nil && !ok
}

func AsBusinessError(err error) (*BusinessError, bool) {
	var businessError *BusinessError
	ok := errors.As(err, &businessError)
	return businessError, ok
}
