package httphandling

import "testing"

func TestHttpStupidFunction(t *testing.T) {
	testNum := httpStupidFunction(4)
	if testNum != 7 {
		t.Error("Expected 7, got ", testNum)
	}
}
