package lines

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"k8s.io/md-check/checks"
)

// Just check for trailing tabs and spaces.
var wsSuffix = regexp.MustCompile(`( |\t)$`)

// updateWhitespace remove all trailing whitespace.
func updateWhitespace(file string, mlines mungeLines) (mungeLines, error) {
	var out mungeLines
	for _, mline := range mlines {
		if mline.preformatted {
			out = append(out, mline)
			continue
		}
		newline := strings.TrimRightFunc(mline.data, unicode.IsSpace)
		out = append(out, newMungeLine(newline, mline.preformatted))
	}
	return out, nil
}

// CheckWhitespace is a simple checker that checks whether there is trailing
// whitespace.
func CheckWhitespace(rptr checks.ErrReporter, fname string, contents []byte) error {
	mlines := getMungeLines(string(contents))
	for i, mline := range mlines {
		if wsSuffix.MatchString(mline.data) {
			rptr.ErrorStr(fmt.Sprintf("File: %v, Line: #%v: <%v> contains trailing whitespace\n",
				fname, i, strings.TrimRight(mline.data, "\n"))) // trim newlines
		}
	}
	return nil
}
