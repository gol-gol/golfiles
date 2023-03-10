package golfiles

import (
	"log"
	"os"
	"strings"
)

func CwdBasename() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirlist := strings.Split(dir, string(os.PathSeparator))
	return dirlist[len(dirlist)-1]
}
