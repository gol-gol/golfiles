package golfiles

import (
	"log"
	"os"
	"strings"
)

/*
CwdBasename returns name of current directory, not path to it.
*/
func CwdBasename() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	dirlist := strings.Split(dir, string(os.PathSeparator))
	return dirlist[len(dirlist)-1]
}

/*
PathExists is a non error-throwing simple boolean proxy for existence of a filesystem level path.
*/
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
