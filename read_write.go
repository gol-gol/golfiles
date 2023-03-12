package golfiles

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func CreateBinaryFile(filename string, blob []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	_, err = f.Write(blob)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	return err
}

func AppendBinaryToFile(filename string, blob []byte) error {
	if filename == "" {
		return errors.New("Filename can't be empty.")
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	_, err = f.Write(blob)
	if err != nil {
		f.Close()
		return err
	}
	err = f.Close()
	return err
}

func AppendToFile(filename, txt string) error {
	if filename == "" {
		return errors.New("Filename can't be empty.")
	}
	fileHandle, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, txt)
	writer.Flush()
	return nil
}

func ReadBinaryFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}
	filesize := fileinfo.Size()

	buffer := make([]byte, filesize)
	if _, err := file.Read(buffer); err != nil {
		return []byte{}, err
	}
	return buffer, nil
}

func ReadFileBuffer(filepath string) (*bytes.Buffer, error) {
	const chunksize int = 1024
	var err error
	var count int

	data, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer := bytes.NewBuffer(make([]byte, 0))
	part := make([]byte, chunksize)

	for {
		count, err = reader.Read(part)
		if err != nil {
			break
		}
		buffer.Write(part[:count])
	}
	if err != io.EOF {
		return nil, err
	}
	return buffer, nil
}
