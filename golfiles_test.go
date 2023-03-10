package golfiles

import "testing"

func TestCwdBasename(t *testing.T) {
	if CwdBasename() != "golfiles" {
		t.Fatal("FAILED for CwdBasename, or the file is not present in dir named: golfiles")
	}
}
