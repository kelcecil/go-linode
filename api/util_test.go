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

func TestSafeVariable(t *testing.T) {
	name := "Ubuntu 8.10 64bit"
	if LabelVariableSafe(name) != "Ubuntu_8_10_64bit" {
		t.Fail()
	}
}

func TestIntToBool(t *testing.T) {
	if IntToBool(0) != "false" {
		t.Fail()
	}
	if IntToBool(1) != "true" {
		t.Fail()
	}
}
