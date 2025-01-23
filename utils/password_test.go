package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	hashPwd, _ := HashPassword("1234")

	if VerifyPassword(hashPwd, "1234") != nil {
		t.Error("Hash verify failed")
	}
}
