
## golfiles

> facade for regular used Go wrapper calls for specific Filesystem operations or info functions
>
> extracted from an older pandora box of such packages at [abhishekkr/gol](https://github.com/abhishekkr/gol)

### Public Functions

* `CwdBasename() string`
* `PathExists(path string) bool`

* `PathLs(basePath string) (map[int]string, error)`
* `PathLsN(basePath string, depth int) (map[int]string, error)`
* `PathLsType(basePath string, filetypes []string) (map[int]string, error)`
* `PathLsNType(basePath string, depth int, filetypes []string) (map[int]string, error)`
* `PathLsTypeNot(basePath string, filetypes []string) (map[int]string, error)`
* `PathLsNTypeNot(basePath string, depth int, filetypes []string) (map[int]string, error)`

---
