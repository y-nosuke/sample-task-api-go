package errors

import (
	"fmt"
)

type ErrorCode int

const (
	BadRequest ErrorCode = iota
	Unauthorized
	Forbidden
	NotFound
	Conflict
)

type BusinessError struct {
	errorCode ErrorCode
	// message code
	doRollBack    bool
	originalError error
	message       string
}

func (e *BusinessError) ErrorCode() ErrorCode {
	return e.errorCode
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

func newBusinessError(errorCode ErrorCode, format string, a ...interface{}) *BusinessError {
	return &BusinessError{errorCode, true, nil, fmt.Sprintf(format, a...)}
}

func newBusinessErrorf(errorCode ErrorCode, originalError error, format string, a ...interface{}) *BusinessError {
	return &BusinessError{errorCode, true, originalError, fmt.Sprintf(format, a...)}
}

// func newNoRollBackBusinessError(errorCode ErrorCode, format string, a ...interface{}) *BusinessError {
// 	return &BusinessError{errorCode, false, nil, fmt.Sprintf(format, a...)}
// }
//
// func newNoRollBackBusinessErrorf(errorCode ErrorCode, originalError error, format string, a ...interface{}) *BusinessError {
// 	return &BusinessError{errorCode, false, originalError, fmt.Sprintf(format, a...)}
// }

func NewBadRequestError(format string, a ...interface{}) *BusinessError {
	return newBusinessError(BadRequest, format, a...)
}

func NewBadRequestErrorf(originalError error, format string, a ...interface{}) *BusinessError {
	return newBusinessErrorf(BadRequest, originalError, format, a...)
}

func NewUnauthorizedError(format string, a ...interface{}) *BusinessError {
	return newBusinessError(Unauthorized, format, a...)
}

func NewUnauthorizedErrorf(originalError error, format string, a ...interface{}) *BusinessError {
	return newBusinessErrorf(Unauthorized, originalError, format, a...)
}

func NewForbiddenError(format string, a ...interface{}) *BusinessError {
	return newBusinessError(Forbidden, format, a...)
}

func NewForbiddenErrorf(originalError error, format string, a ...interface{}) *BusinessError {
	return newBusinessErrorf(Forbidden, originalError, format, a...)
}

func NewNotFoundError(format string, a ...interface{}) *BusinessError {
	return newBusinessError(NotFound, format, a...)
}

func NewNotFoundErrorf(originalError error, format string, a ...interface{}) *BusinessError {
	return newBusinessErrorf(NotFound, originalError, format, a...)
}

func NewConflictError(format string, a ...interface{}) *BusinessError {
	return newBusinessError(NotFound, format, a...)
}

func NewConflictErrorf(originalError error, format string, a ...interface{}) *BusinessError {
	return newBusinessErrorf(NotFound, originalError, format, a...)
}
