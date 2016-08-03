package checks

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// A filter to select only markdown files.
func IsMarkdown(fpath string) bool {
	if strings.HasSuffix(fpath, ".md") {
		return true
	}
	return false
}

// A filter to select only YAML files.
func IsYAML(fpath string) bool {
	if strings.HasSuffix(fpath, ".yaml") || strings.HasSuffix(fpath, ".yml") {
		return true
	}
	return false
}

// Walker
type Walker struct {
	Dir      string
	Filter   func(filename string) bool
	Rptr     ErrReporter
	CheckFns []CheckFunc
}

// Walk a file tree and call the walkFn
func (w *Walker) Walk() error {
	return filepath.Walk(w.Dir, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if w.Filter(p) {
			var contents []byte
			contents, err = ioutil.ReadFile(p)
			if err != nil {
				return err
			}
			for _, fn := range w.CheckFns {
				checkErr := fn(w.Rptr, p, contents)
				if checkErr != nil {
					return checkErr // recall that returning errors stops execption
				}
			}
		}
		return nil
	})
}

// CheckFunc is given a filename and an error reporter
type CheckFunc func(rptr ErrReporter, filename string, contents []byte) error
