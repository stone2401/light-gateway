package public

import (
	"crypto/sha256"
	"fmt"
)

func GenSaltPassword(salt, passowrd string) string {
	s1 := sha256.New()
	s1.Write([]byte(passowrd))
	st1 := fmt.Sprintf("%x", s1.Sum(nil))
	s2 := sha256.New()
	s2.Write([]byte(st1 + salt))
	return fmt.Sprintf("%x", s1.Sum(nil))
}
