/*
Copyright 2015-2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
	"k8s.io/md-check/checks"
	"k8s.io/md-check/checks/lines"
	"k8s.io/md-check/checks/md"
)

var (
	rootDir = flag.String("root-dir", "", "Root directory containing documents to be processed.")
)

const (
	// TODO(Kashomon): Use some sort of commandline interface (e.g.,
	// https://github.com/urfave/cli)
	Usage = `Usage: md-check --root-dir=<directory-path>`
)

func main() {
	var err error
	flag.Parse()

	errReporter := &checks.StdErrReporter{}

	absRootDir, err := filepath.Abs(*rootDir)
	if err != nil {
		errReporter.FatalErrorStr(fmt.Sprintf("Error: %v\n", err))
	}

	w := &checks.Walker{
		Dir:    absRootDir,
		Filter: checks.IsMarkdown,
		Rptr:   errReporter,
		CheckFns: []checks.CheckFunc{
			md.ParseCheck,
			lines.CheckWhitespace,
		},
	}

	err = w.Walk()
	if err != nil {
		errReporter.FatalErrorStr(fmt.Sprintf("Error: %v\n", err))
	}

	if errReporter.ReportedErr() {
		fail()
	} else {
		success()
	}
}

// validateRootDir validates that the rootDir passed in is reasonable and
// then returns the absoute path.
func validateRootDir(path string) (string, error) {
	if *rootDir == "" {
		return "", fmt.Errorf("Root dir must be specified. Was empty\n")
	}

	absRootDir, err := filepath.Abs(*rootDir)
	if err != nil {
		return "", err
	}

	var stat os.FileInfo
	stat, err = os.Stat(path)
	if err != nil {
		return "", err
	}

	if !stat.IsDir() {
		return "", fmt.Errorf("root-dir must be a directory, but was not!")
	}

	return absRootDir, nil
}

func success() {
	fmt.Printf(`----------------
-- ✔ Success! --
----------------`)
}

func fail() {
	fmt.Printf(`-------------
-- ✘ Fail! --
-------------`)
}
