package funcman

func Map[I, O any](input []I, f func(input I) O) []O {
	result, _ := MapWithError(input, func(input I) (O, error) { return f(input), nil })
	return result
}

func MapWithError[I, O any](input []I, f func(input I) (O, error)) ([]O, error) {
	result := make([]O, len(input), cap(input))

	for i, val := range input {
		var err error
		result[i], err = f(val)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// TODO: reduce, for each, filter
