package presenter

type ErrorPresenter interface {
	BusinessErrorPresenter
	SystemErrorPresenter
}
