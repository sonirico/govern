package govern

type (
	field[T any] interface {
		Check(T) *Error

		Name() string
	}

	enumField[T any] interface {
		field[T]

		enumField()
	}

	maybeField[T any] interface {
		field[T]

		maybeField()
	}

	Schema[T any] struct {
		fields map[string]field[T]
	}
)

func (s *Schema[T]) Check(x T) Result {
	var (
		errors []Error
	)

	for _, f := range s.fields {
		if err := f.Check(x); err != nil {
			errors = append(errors, *err)
		}
	}

	return Result{Errors: errors}
}

func (s *Schema[T]) Enum(f enumField[T]) *Schema[T] {
	s.fields[f.Name()] = f
	return s
}

func (s *Schema[T]) Maybe(f maybeField[T]) *Schema[T] {
	s.fields[f.Name()] = f
	return s
}

func (s *Schema[T]) String(name string, getter Getter[T, string], opts StringOpts) *Schema[T] {
	s.fields[name] = fieldType[T, string]{
		name:   name,
		getter: getter,
		inner:  String(opts),
	}
	return s
}

func New[T any]() *Schema[T] {
	return &Schema[T]{fields: make(map[string]field[T])}
}
