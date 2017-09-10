package mem

import (
	"fmt"
	"testing"
)

func TestInterrupt(t *testing.T) {
	intr, err := Interrupts()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(intr.Version)
}

func TestFork(t *testing.T) {
	fork, err := Forks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fork.Version)
}

func TestMalloc(t *testing.T) {
	malloc, err := Mallocs()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(malloc.Version)
}

func TestObjects(t *testing.T) {
	objs, err := Objects()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(objs.Version)
}

func TestSumStats(t *testing.T) {
	sum, err := SumStats()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sum.Version)
}
