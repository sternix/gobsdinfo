package arp

import "testing"

func TestArpEntries(t *testing.T) {
	_ ,err := Entries()
	if err != nil {
		t.Fatal(err)
	}
}
