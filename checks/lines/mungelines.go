package lines

import (
	"regexp"
	"strings"
)

var (
	// Finds all preformatted block start/stops.
	preformatRE    = regexp.MustCompile("^\\s*```")
	notPreformatRE = regexp.MustCompile("^\\s*```.*```")
	// Is this line a header?
	mlHeaderRE = regexp.MustCompile(`^#`)
	// Is there a link on this line?
	mlLinkRE   = regexp.MustCompile(`\[[^]]*\]\([^)]*\)`)
	beginTagRE = regexp.MustCompile(`<!-- BEGIN MUNGE:`)
	endTagRE   = regexp.MustCompile(`<!-- END MUNGE:`)

	blankMungeLine = newMungeLine("", false)
)

// mungeLines is a collection of mungeline structs.
type mungeLines []mungeLine

// mungeLine represents a single line from a file.
type mungeLine struct {
	lineNum      int
	data         string
	preformatted bool
	header       bool
	link         bool
	beginTag     bool
	endTag       bool
}

// newMungeLine creates a newMungeLine.
func newMungeLine(line string, preformatted bool) mungeLine {
	return mungeLine{
		data:         line,
		preformatted: preformatted,
		header:       mlHeaderRE.MatchString(line),
		link:         mlLinkRE.MatchString(line),
		beginTag:     beginTagRE.MatchString(line),
		endTag:       endTagRE.MatchString(line),
	}
}

// getMungeLines takes the fullText of a file and converts it to mungeLines.
func getMungeLines(document string) mungeLines {
	var out mungeLines
	preformatted := false

	lines := splitLines(document)
	// We indicate if any given line is inside a preformatted block or
	// outside a preformatted block
	for _, line := range lines {
		if !preformatted {
			if preformatRE.MatchString(line) && !notPreformatRE.MatchString(line) {
				preformatted = true
			}
		} else {
			if preformatRE.MatchString(line) {
				preformatted = false
			}
		}
		ml := newMungeLine(line, preformatted)
		out = append(out, ml)
	}
	return out
}

// splitLines splits a document into a slice of lines.
func splitLines(document string) []string {
	lines := strings.Split(document, "\n")
	// Skip trailing empty string from Split-ing
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
