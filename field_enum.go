package govern

type (
	EnumField[T any, U comparable] struct {
		name   string
		inner  EnumType[U]
		getter Getter[T, U]
	}

	EnumType[T comparable] struct {
		values []T
		inner  Type[T]
	}
)

func (f EnumField[T, U]) Name() string {
	return f.name
}

func (f EnumField[T, U]) Check(x T) *Error {
	raw := f.getter(x)

	if err := f.inner.Check(raw); err != nil {
		return newErr(err, f.name, raw)
	}

	return nil
}

func (f EnumField[T, U]) enumField() {}

func (t EnumType[T]) Values() []T {
	return t.values
}

func (t EnumType[T]) Check(x T) error {
	if err := t.inner.Check(x); err != nil {
		return err
	}

	for _, v := range t.values {
		if v == x {
			return nil
		}
	}

	// TODO: wrap
	return ErrValueNotInEnum
}

func FieldEnum[T any, U comparable](
	name string,
	inner Type[U],
	getter Getter[T, U],
	values ...U,
) EnumField[T, U] {
	return EnumField[T, U]{
		name:   name,
		getter: getter,
		inner:  Enum[U](inner, values...),
	}
}

func Enum[T comparable](inner Type[T], values ...T) EnumType[T] {
	return EnumType[T]{
		values: values,
		inner:  inner,
	}
}
