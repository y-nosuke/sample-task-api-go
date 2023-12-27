package errors

import "golang.org/x/xerrors"

type _error struct {
	error error
}

func (e *_error) Error() string {
	return e.error.Error()
}

func _new(format string, a ...interface{}) *_error {
	return &_error{error: xerrors.Errorf(format, a)}
}

type BusinessError struct {
	_error
}

func BusinessErrorf(format string, a ...interface{}) *BusinessError {
	return &BusinessError{_error{_new(format, a)}}
}

type SystemError struct {
	_error
}

func SystemErrorf(format string, a ...interface{}) *SystemError {
	return &SystemError{_error{_new(format, a)}}
}
