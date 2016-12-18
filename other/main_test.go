package main

import (
	"example.com/common/null"
	"testing"
)

func TestOK(t *testing.T) {
	if null.One != 1 {
		t.Fatal("Expected 1 got", null.One)
	}
}
