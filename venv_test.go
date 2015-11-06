package drawer

import (
	"os"
	"path"
	"testing"
)

func newTestEnv(t *testing.T) string {
	startDir := "tests/d1/d1.1/d1.1.1"
	err := os.MkdirAll(startDir, 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Mkdir("tests/d1/src", 0755)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Mkdir("tests/d1/d1.2", 0755)
	if err != nil {
		t.Fatal(err)
	}

	return startDir
}

func removeTestEnv(t *testing.T) {
	err := os.RemoveAll("tests")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_FindSrcDir(t *testing.T) {
	startDir := newTestEnv(t)
	defer removeTestEnv(t)

	dir, err := FindSrcDir(startDir)
	if err != nil {
		t.Fatal(err)
	}

	cur, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	expect := path.Join(cur, "tests", "d1")

	if expect != dir {
		t.Fatal("Expect", expect, "was", dir)
	}
}

func Test_ReplaceBinPath(t *testing.T) {
	p := "/a/b:/c/d/:/v:a/b"

	r := NewPath(p, "/a/b", "/blub/blob")
	expect := "/c/d/:/v:a/b:/blub/blob"

	if expect != r {
		t.Fatal("Expect", expect, "was", r)
	}
}
