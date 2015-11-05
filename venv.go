package venv

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	SrcName = "src"
	NoSrc   = errors.New("No GOPATH directory found")
)

func FindSrcDir(start string) (string, error) {

	tmp := filepath.ToSlash(start)
	tmp = filepath.Clean(tmp)
	tmp, err := filepath.Abs(tmp)
	if err != nil {
		return "", err
	}

	src, err := find(tmp)
	if err != nil {
		return "", err
	}

	src, err = filepath.Abs(src)
	if err != nil {
		return "", err
	}

	return src, nil

}

func find(p string) (string, error) {
	if len(p) == 1 && p == string(os.PathSeparator) {
		return "", NoSrc
	}

	infos, err := ioutil.ReadDir(p)
	if err != nil {
		return "", err
	}

	r := isSrc(infos)
	if r {
		return p, nil
	}

	return find(filepath.Dir(p))
}

func isSrc(infos []os.FileInfo) bool {
	for _, info := range infos {
		if info.Name() == SrcName {
			return true
		}
	}

	return false
}

func AppendEnvList(key, val string) error {
	newVal := os.Getenv(key) + string(os.PathListSeparator) + val
	return os.Setenv(key, newVal)
}
