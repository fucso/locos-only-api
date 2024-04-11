package base

type domain interface{}

type builder[T domain] interface {
	Build() (*T, error)
}
