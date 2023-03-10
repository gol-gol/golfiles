package golfiles

import (
	"os"
	"testing"
)

func TestCwdBasename(t *testing.T) {
	if CwdBasename() != "golfiles" {
		t.Fatal("FAILED for CwdBasename, or the file is not present in dir named: golfiles")
	}
}

func TestPathExists(t *testing.T) {
	if PathExists("/this/should/not/exist") {
		t.Fatal("FAILED for PathExists when given missing path, or dir exists: /this/should/not/exist")
	}

	dir, _ := os.Getwd()
	if !PathExists(dir) {
		t.Fatalf("FAILED for PathExists when given current path, dir: %s", dir)
	}
}
