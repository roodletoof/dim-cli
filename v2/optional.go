package dimcli

type Optional[T any] struct {
	isSome bool
	value T
}

func Some[T any](value T) Optional[T] {
	return Optional[T]{
		isSome: true,
		value: value,
	}
}

func None[T any]() Optional[T] {
	return Optional[T]{
		isSome: false,
	}
}

func (o *Optional[T]) Get() (T, bool) {
	return o.value, o.isSome
}
