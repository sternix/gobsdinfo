package net

import (
	"testing"
)

func TestMbuf(t *testing.T) {
	_, err := Mbufs()
	if err != nil {
		t.Error(err)
	}
}

func TestIsr(t *testing.T) {
	_, err := Isrs()
	if err != nil {
		t.Error(err)
	}
}
