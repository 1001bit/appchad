package crypt

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandomHex(size int) (string, error) {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	return fmt.Sprintf("%x", buf), err
}
