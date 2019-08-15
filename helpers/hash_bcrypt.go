package helpers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHashMake from plain string
func BcryptHashMake(plainText string) string {
	str := []byte(plainText)
	hashedStr, err := bcrypt.GenerateFromPassword(str, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hashedStr)
}

// BcryptHashCompare between plainText and hashedText
func BcryptHashCompare(plainText, hashedText string) bool {
	matchError := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
	if matchError != nil {
		return false
	}
	return true
}
