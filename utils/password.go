package utils

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

// HashPassword generates a bcrypt hash of the password using cost from BCRYPT_COST env if present.
func HashPassword(password string) (string, error) {
	cost := bcrypt.DefaultCost
	if c := os.Getenv("BCRYPT_COST"); c != "" {
		if v, err := strconv.Atoi(c); err == nil {
			cost = v
		}
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CheckPassword compares a bcrypt hashed password with its possible plaintext equivalent.
func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
