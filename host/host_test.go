package host

import (
	"testing"
)

func TestUptimeInfo(t *testing.T) {
	_, err := UptimeInfo()
	if err != nil {
		t.Fatal(err)
	}
}
