package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomStringSource(10)
	if len(s) != 10 {
		t.Error("Random string is not 10 characters long")
	}
}
