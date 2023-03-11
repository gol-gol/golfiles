package golfiles

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestMkDirWithPermission(t *testing.T) {
	dirpath := "./temp-test-mkdirperm"
	err := MkDirWithPermission(dirpath, 0700)
	if err != nil {
		t.Errorf("FAILED to mkdir %s with 0700 permissions.", dirpath)
	}
	stat, errStat := os.Stat(dirpath)
	if errStat != nil {
		t.Errorf("FAILED to mkdir %s.", dirpath)
	}
	if !stat.Mode().IsDir() {
		t.Errorf("FAILED to create %s as directory.", dirpath)
	}
	if stat.Mode().Perm() != 0700 {
		t.Errorf("FAILED to create %s with correct permissions.", dirpath)
	}
	os.Remove(dirpath)
}

func TestMkDir(t *testing.T) {
	dirpath := "./temp-test-mkdir"
	err := MkDir(dirpath)
	if err != nil {
		t.Errorf("FAILED to mkdir %s.", dirpath)
	}
	stat, errStat := os.Stat(dirpath)
	if errStat != nil {
		t.Errorf("FAILED to mkdir %s.", dirpath)
	}
	if !stat.Mode().IsDir() {
		t.Errorf("FAILED to create %s as directory.", dirpath)
	}
	if stat.Mode().Perm() != 0755 {
		t.Errorf("FAILED to create %s with correct permissions.", dirpath)
	}
	os.Remove(dirpath)
}

func TestCopyFile(t *testing.T) {
	dirpath := "./temp-test-copy"
	srcDir := path.Join(dirpath, "src")
	dstDir := path.Join(dirpath, "dst")
	src := path.Join(srcDir, "somefile")
	dst := path.Join(dstDir, "thisfile")
	MkDir(srcDir)
	MkDir(dstDir)
	ioutil.WriteFile(src, []byte("data"), 0644)

	CopyFile(src, dst)
	stat, errStat := os.Stat(dst)
	if errStat != nil {
		t.Errorf("FAILED to copy to %s.", dst)
	}
	if !stat.Mode().IsRegular() {
		t.Errorf("FAILED to create %s as file.", dst)
	}
	if stat.Mode().Perm() != 0644 {
		t.Errorf("FAILED to create %s with correct permissions.", dst)
	}

	data, errRead := ioutil.ReadFile(dst)
	if errRead != nil {
		t.Errorf("FAILED to read from %s", dst)
	}
	if string(data) != "data" {
		t.Errorf("FAILED to copy correct data to %s", dst)
	}

	os.Remove(src)
	os.Remove(srcDir)
	os.Remove(dst)
	os.Remove(dstDir)
	os.Remove(dirpath)
}

func TestMoveFile(t *testing.T) {
	dirpath := "./temp-test-move"
	srcDir := path.Join(dirpath, "src")
	dstDir := path.Join(dirpath, "dst")
	src := path.Join(srcDir, "somefile")
	dst := path.Join(dstDir, "thisfile")
	MkDir(srcDir)
	MkDir(dstDir)
	ioutil.WriteFile(src, []byte("data"), 0644)

	MoveFile(src, dst)

	_, errStatSrc := os.Stat(src)
	if errStatSrc == nil {
		t.Errorf("FAILED to move from %s to %s.", src, dst)
	}

	stat, errStat := os.Stat(dst)
	if errStat != nil {
		t.Errorf("FAILED to copy to %s.", dst)
	}
	if !stat.Mode().IsRegular() {
		t.Errorf("FAILED to create %s as file.", dst)
	}
	if stat.Mode().Perm() != 0644 {
		t.Errorf("FAILED to create %s with correct permissions.", dst)
	}

	data, errRead := ioutil.ReadFile(dst)
	if errRead != nil {
		t.Errorf("FAILED to read from %s", dst)
	}
	if string(data) != "data" {
		t.Errorf("FAILED to copy correct data to %s", dst)
	}

	os.Remove(src)
	os.Remove(srcDir)
	os.Remove(dst)
	os.Remove(dstDir)
	os.Remove(dirpath)
}
