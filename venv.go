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

// Walk from the current directory backwards until we find a directory which
// contains a src directory.
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

func findPath(paths []string, p string) (int, bool) {
	for pos, elem := range paths {
		if elem == p {
			return pos, true
		}
	}

	return -1, false
}

// Make a new PATH string, therefor remove the old go bin path if it exists
// and append the new go bin path at the end.
func NewPath(pathEnv, o, n string) string {
	vals := strings.Split(pathEnv, string(os.PathListSeparator))

	pos, found := findPath(vals, o)
	if found {
		vals = append(vals[:pos], vals[pos+1:]...)
	}

	vals = append(vals, n)

	return strings.Join(vals, string(os.PathListSeparator))

}
