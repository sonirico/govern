package govern

type (
	MaybeField[T any, U comparable] struct {
		name   string
		getter Getter[T, *U]
		inner  Type[*U]
	}

	MaybeType[T comparable] struct {
		required bool
		inner    Type[T]
	}
)

func (f MaybeField[T, U]) Name() string {
	return f.name
}

func (f MaybeField[T, U]) Check(x T) *Error {
	raw := f.getter(x)

	if err := f.inner.Check(raw); err != nil {
		return newErr(err, f.name, raw)
	}

	return nil
}

func (f MaybeField[T, U]) maybeField() {}

func (f MaybeType[T]) Check(x *T) error {
	if f.required && x == nil {
		return ErrRequired
	}

	if x == nil {
		return nil
	}

	if err := f.inner.Check(*x); err != nil {
		return err
	}

	return nil
}

func FieldMaybe[T any, U comparable](
	name string,
	inner Type[U],
	required bool,
	getter Getter[T, *U],
) MaybeField[T, U] {
	return MaybeField[T, U]{
		name:   name,
		getter: getter,
		inner:  Maybe[U](inner, required),
	}
}

func Maybe[T comparable](inner Type[T], required bool) MaybeType[T] {
	return MaybeType[T]{
		inner:    inner,
		required: required,
	}
}
