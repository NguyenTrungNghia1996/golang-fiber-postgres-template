package utils

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	t.Setenv("BCRYPT_COST", "4")
	hash, err := HashPassword("secret")
	if err != nil {
		t.Fatalf("hash error: %v", err)
	}
	if !CheckPassword("secret", hash) {
		t.Fatalf("password does not match")
	}
}
