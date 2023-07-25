package public

import "testing"

func TestGenSaltPassword(t *testing.T) {
	s := GenSaltPassword("aksldjalskdj", "shizhenfei123")
	t.Log(s)
}
