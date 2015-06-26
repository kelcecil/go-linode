package api

import (
	"testing"
)

func TestUpperCase(t *testing.T) {
	word := "hello"
	if UpperCaseFirstLetter(word) != "Hello" {
		t.Fail()
	}
}
