package utils

import "golang.org/x/crypto/bcrypt"

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

func BcryptMake(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(b), err
}

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
