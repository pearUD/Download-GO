package files

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	onceD sync.Once
	sep   = "/"
	oSep  = string(os.PathSeparator)
)

func OpenFile(path, name string) *os.File {
	if sep == oSep {
		sep = "\\"
	}
	path = strings.ReplaceAll(path, sep, oSep)
	path = OpenDir(path, 0)
	name, suffix := getSuffix(name)
	if suffix != "" {
		suffix = "." + suffix
	}
	file, err := OpenNewFile(path, name, suffix, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return file
}

func OpenNewFile(path, name, suffix string, x int) (*os.File, error) {
	nameN := path + string(os.PathSeparator) + name
	if x != 0 {
		nameN += strconv.Itoa(x)
	}
	nameN += suffix
	if _, err := os.Stat(nameN); err == nil {
		return OpenNewFile(path, name, suffix, x+1)
	}
	return os.OpenFile(nameN, os.O_CREATE|os.O_WRONLY, 0666)
}

func OpenDir(path string, x int) string {
	pathN := path
	if x != 0 {
		pathN += strconv.Itoa(x)
	}
	if i, err := os.Stat(pathN); err == nil {
		if !i.IsDir() {
			onceD.Do(func() {
				log.Printf("已有同名文件/文件夹")
			})
			return OpenDir(path, x+1)
		}
		return pathN
	}
	err := os.MkdirAll(pathN, 0666)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return pathN
}

func getSuffix(name string) (fName, suffix string) {
	if strings.Contains(name, ".") {
		for r := len(name) - 1; r > 0; r-- {
			if name[r-1:r] == "." {
				return name[:r-1], name[r:]
			}
		}
	}
	return name, ""
}
