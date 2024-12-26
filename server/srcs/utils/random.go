package utils

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func RandStr() string {
	letters := "abcdefhijklmnopqrstuvwxyzABCDEFHIJKLMNOPQRSTUVWXYZ"
	token := ""

	for i := 0; i < 10; i++ {
		n := rand.Intn(len(letters))
		token += string(letters[n])
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(token), 14)
	token = string(hash)
	return token
}
