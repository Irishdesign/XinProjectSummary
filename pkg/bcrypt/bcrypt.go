package bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash), err
}

func Check(p string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))

	if err != nil {
		return false
	} else {
		return true
	}
}
