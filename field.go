package govern

type (
	Field[T any] interface {
		Check(T) *Error
	}

	Type[T comparable] interface {
		Check(T) error
	}

	Getter[T any, U comparable] func(x T) U
)
