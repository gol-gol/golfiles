package golfiles

import (
	"fmt"
	"os"
	"testing"
)

var (
	RandomBytes = []byte{0, 10, 13, 41}
	TmpFile     = "testfile.tmp"
)

func TestCreateBinaryFileAndReadBinaryFile(t *testing.T) {
	if errCreate := CreateBinaryFile(TmpFile, RandomBytes); errCreate != nil {
		t.Error("FAILED to create binary file.")
	}

	bites, errRead := ReadBinaryFile(TmpFile)
	if errRead != nil {
		t.Error("FAILED to read binary file.")
	}
	for idx, b := range bites {
		if b != RandomBytes[idx] {
			t.Error("FAILED to read correct bytes, or failed to write correctly.")
		}
	}

	os.Remove(TmpFile)
}

func TestAppendBinaryToFile(t *testing.T) {
	CreateBinaryFile(TmpFile, []byte("Say"))
	if err := AppendBinaryToFile(TmpFile, []byte(" What!")); err != nil {
		t.Error("FAILED to append to file.")
	}
	bites, errRead := ReadBinaryFile(TmpFile)
	if errRead != nil {
		t.Error("FAILED to read appended file.")
	} else if string(bites) != "Say What!" {
		fmt.Println(string(bites))
		t.Error("FAILED to read required data from appended file.")
	}
	os.Remove(TmpFile)
}

func TestAppendToFile(t *testing.T) {
	CreateBinaryFile(TmpFile, []byte("Say"))
	if err := AppendToFile(TmpFile, " What!"); err != nil {
		t.Error("FAILED to append to file.")
	}
	bites, errRead := ReadBinaryFile(TmpFile)
	if errRead != nil {
		t.Error("FAILED to read appended file.")
	} else if string(bites) != "Say What!\n" {
		t.Error("FAILED to read required data from appended file.")
	}
	os.Remove(TmpFile)
}

func TestReadFileBuffer(t *testing.T) {
	CreateBinaryFile(TmpFile, RandomBytes)
	buf, err := ReadFileBuffer(TmpFile)
	if err != nil {
		t.Error("FAILED to read buffer from file.")
	}
	bites := buf.Bytes()
	for idx, b := range bites {
		if b != RandomBytes[idx] {
			t.Error("FAILED to read correct bytes, or failed to write correctly.")
		}
	}
	os.Remove(TmpFile)
}
