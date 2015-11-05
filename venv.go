package venv

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	SrcName = "src"
	NoSrc   = errors.New("No GOPATH directory found")
)

func FindSrcDir(start string) (string, error) {

	tmp := filepath.ToSlash(start)
	tmp = filepath.Clean(tmp)

	src, err := find(strings.Split(tmp, string(os.PathSeparator)))
	if err != nil {
		return "", err
	}

	src, err = filepath.Abs(src)
	if err != nil {
		return "", err
	}

	return src, nil

}

func find(p []string) (string, error) {
	if len(p) == 0 {
		return "", NoSrc
	}

	currPath := filepath.Join(p...)
	infos, err := ioutil.ReadDir(currPath)
	if err != nil {
		return "", err
	}

	r := isSrc(infos)
	if r {
		return currPath, nil
	}

	return find(p[0 : len(p)-1])
}

func isSrc(infos []os.FileInfo) bool {
	for _, info := range infos {
		if info.Name() == SrcName {
			return true
		}
	}

	return false
}
