package errors

import "errors"

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
