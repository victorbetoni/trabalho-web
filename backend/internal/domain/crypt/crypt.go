package crypt

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(input string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(input), 10)
}
