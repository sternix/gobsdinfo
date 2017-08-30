package proc

import (
	ps "github.com/mitchellh/go-ps"
	"testing"
)

/*
Bench xo vs sysctl versions
*/

func BenchmarkProcessesViaSysctl(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := ps.Processes()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProcessesViaXo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := All()
		if err != nil {
			b.Fatal(err)
		}
	}
}
