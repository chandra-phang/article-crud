package parse

type Params[T any] struct {
	key      string
	values   []string
	typeconv TypeConv[T]
}

func NewParams[T any](key string, values []string, typeconv TypeConv[T]) Params[T] {
	return Params[T]{
		key,
		values,
		typeconv,
	}
}

func (p Params[T]) All() ([]T, error) {
	if len(p.values) == 0 {
		return nil, newParamNotFound(p.key)
	}

	var ret []T
	for _, v := range p.values {
		vv, err := p.typeconv(v)
		if err != nil {
			return nil, newParamInvalidType(p.key, err)
		}
		ret = append(ret, vv)
	}
	return ret, nil
}

func (p Params[T]) First() (T, error) {
	var ret T

	all, err := p.All()
	if err != nil {
		return ret, err
	}
	return all[0], nil
}
