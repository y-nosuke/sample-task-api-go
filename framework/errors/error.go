package errors

const (
	Unauthorized = iota
	Forbidden
	NotFound
	Conflict
)

type AppError struct {
	error
	Status  int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(status int, message string, err error) *AppError {
	return &AppError{err, status, message}
}
