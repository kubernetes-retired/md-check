// Package md contains checks for markdown.
package md

import (
	"fmt"

	"github.com/russross/blackfriday"
	"k8s.io/md-check/checks"
)

// ParseCheck ensures that
func ParseCheck(rptr checks.ErrReporter, fname string, contents []byte) error {
	defer func() {
		if err := recover(); err != nil {
			rptr.ErrorStr(fmt.Sprintf("Markdown Parse Error: %v\n", err))
		}
	}()
	// Check to make sure the markdown parses successfully, but don't do anything
	// special. Note: If parsing fails, blackfriday will panic (unfortunately).
	blackfriday.MarkdownBasic(contents)
	return nil // Continue walking the tree
}

// Ensure that ParseCheck is a CheckFunc
var _ checks.CheckFunc = ParseCheck
