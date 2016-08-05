package lines

import (
	"strings"
	"testing"
)

type wsErrorRep struct {
	lines []string
}

func (w *wsErrorRep) ErrorStr(s string) {
	w.lines = append(w.lines, s)
}

func (w *wsErrorRep) FatalErrorStr(s string) {
	w.lines = append(w.lines, s)
}

func (w *wsErrorRep) ReportedErr() bool {
	return len(w.lines) > 0
}

func TestWhitespace(t *testing.T) {
	text := `
Time flies
like an arrow.	
Fruit flies 
like a banana.
`

	erp := &wsErrorRep{}
	CheckWhitespace(erp, "foo.md", []byte(text))

	if !erp.ReportedErr() {
		t.Errorf("Expected an error to have been reported")
	}

	if len(erp.lines) != 2 {
		t.Errorf("Expected exactly two errors to have been reported")
	}

	if !strings.Contains(erp.lines[0], "arrow") {
		t.Errorf("Expected first error to contain arrow")
	}

	if !strings.Contains(erp.lines[1], "Fruit") {
		t.Errorf("Expected second error to contain Fruit")
	}
}
