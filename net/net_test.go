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

func TestIfStat(t *testing.T) {
	_, err := IfStats()
	if err != nil {
		t.Error(err)
	}
}

func TestUnixSocketStat(t *testing.T) {
	_, err := UnixSocketStats()
	if err != nil {
		t.Error(err)
	}
}
