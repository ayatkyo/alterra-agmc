package utils

func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, item := range collection {
		if predicate(item) {
			return item, true
		}
	}

	var result T
	return result, false
}

func FindIndex[T any](collection []T, predicate func(T) bool) (int, bool) {
	for i, item := range collection {
		if predicate(item) {
			return i, true
		}
	}

	return -1, false
}

func Delete[T any](collection *[]T, predicate func(T) bool) {
	list := []T{}

	for _, item := range *collection {
		if !predicate(item) {
			list = append(list, item)
		}
	}

	*collection = list
}
