package govern

type (
	fieldType[T any, U comparable] struct {
		name   string
		getter Getter[T, U]
		inner  Type[U]
	}
)

func (f fieldType[T, U]) Name() string {
	return f.name
}

func (f fieldType[T, U]) Check(x T) *Error {
	value := f.getter(x)
	if err := f.inner.Check(value); err != nil {
		return newErr(err, f.name, value)
	}

	return nil
}
