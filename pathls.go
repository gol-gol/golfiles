package golfiles

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
PathLs returns a sort-of-linear representational mapping
of directory "/a"
with result like ["/a/b", "/a/b", "/a/b/cc", "/a/b/ccc/d"]
*/
func PathLs(basePath string) ([]string, error) {
	return doPathWalk(basePath, -1, []string{}, false)
}

/*
PathLsN is variety of PathLs function where the depth of traversal can be limited.
It need to be provided path and depth to which traversal is required.
Depth can be passed -1 if need to not limit. But use PathLs instead then.
*/
func PathLsN(basePath string, depth int) ([]string, error) {
	return doPathWalk(basePath, depth, []string{}, false)
}

/*
PathLsType is variety of PathLs function where filetype looked for can be limited to a list.
The filetype list could look like [".mp4", ".avi"]
*/
func PathLsType(basePath string, filetypes []string) ([]string, error) {
	return doPathWalk(basePath, -1, filetypes, true)
}

/*
PathLsNType is variety of PathLsN function where filetype looked for can be limited to a list.
*/
func PathLsNType(basePath string, depth int, filetypes []string) ([]string, error) {
	return doPathWalk(basePath, depth, filetypes, true)
}

/*
PathLsTypeNot is variety of PathLs function where filetype looked for can be limited to a list.
The filetype list could look like [".mp4", ".avi"]
*/
func PathLsTypeNot(basePath string, filetypes []string) ([]string, error) {
	return doPathWalk(basePath, -1, filetypes, false)
}

/*
PathLsNTypeNot is variety of PathLsN function where filetype looked for can be limited to a list.
*/
func PathLsNTypeNot(basePath string, depth int, filetypes []string) ([]string, error) {
	return doPathWalk(basePath, depth, filetypes, false)
}

/*
doPathWalk is base core engine for varied exposed features by PathLs.* functions.
*/
func doPathWalk(dirPath string, recurseDepth int, filetypes []string, filetypeBool bool) ([]string, error) {
	var pathMap = []string{}
	fullPath, err := filepath.Abs(dirPath)
	if err != nil {
		return pathMap, err
	}
	if !PathExists(fullPath) {
		errMsg := fmt.Sprintf("Path '%s' gets translated to '%s', which doesn't exist.", dirPath, fullPath)
		return pathMap, errors.New(errMsg)
	}

	callback := func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return walkFile(fullPath, path, fi, recurseDepth, filetypes, filetypeBool, &pathMap)
	}

	errWalk := filepath.Walk(fullPath, callback)
	return pathMap, errWalk
}

/*
walkFile gets called on all entries of filepath.Walk, it populates inferred pathMap for a cumulative result.
*/
func walkFile(root string, path string, fi os.FileInfo, recurseDepth int, filetypes []string, filetypeBool bool, pathMap *[]string) error {
	if fi.IsDir() {
		return nil
	}
	relativePath, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}

	currentDepth := len(strings.Split(relativePath, "/")) - 1
	if recurseDepth != -1 && recurseDepth < currentDepth {
		return nil
	}
	if isExtGroup(path, filetypes) == filetypeBool {
		*pathMap = append(*pathMap, path)
	}
	return nil
}

/*
isExtGroup checks if provided ext belongs in array ext_list.
*/
func isExtGroup(filename string, ext_list []string) bool {
	ext := filepath.Ext(filename)
	for _, _ext := range ext_list {
		if _ext == ext {
			return true
		}
	}
	return false
}
