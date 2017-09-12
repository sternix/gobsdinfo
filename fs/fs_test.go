package fs

import "testing"

func TestDF(t *testing.T) {
	_, err := Info()
	if err != nil {
		t.Error(err)
	}
}
