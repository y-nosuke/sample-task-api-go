package presenter

type ErrorPresenterImpl struct {
	BusinessErrorPresenterImpl
	SystemErrorHandlerPresenterImpl
}

func NewErrorPresenterImpl() *ErrorPresenterImpl {
	return &ErrorPresenterImpl{
		BusinessErrorPresenterImpl{},
		SystemErrorHandlerPresenterImpl{},
	}
}
