package user

import "golang.org/x/crypto/bcrypt"

func createHash(s string) string {
	bytePassword := []byte(s)
	hashPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashPassword)
}
