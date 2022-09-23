package utils

import "golang.org/x/crypto/bcrypt"

func BcryptMake(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 7)
	return string(b), err
}

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func BcryptMustMake(password string) string {
	b, _ := BcryptMake(password)
	return b
}
