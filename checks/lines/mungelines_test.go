package lines

import (
	"reflect"
	"testing"
)

func TestNewMungeLines(t *testing.T) {
	text := "foo\n" +
		"```\n" +
		"bar\n" +
		"```\n" +
		"#ImaHeader\n" +
		"<!-- BEGIN MUNGE: zed>\n" +
		"<!-- END MUNGE: zod>\n" +
		"[ima](link)"

	l := getMungeLines(text)

	expected := []mungeLine{
		mungeLine{data: "foo"},
		mungeLine{data: "```", preformatted: true},
		mungeLine{data: "bar", preformatted: true},
		mungeLine{data: "```"},
		mungeLine{data: "#ImaHeader", header: true},
		mungeLine{data: "<!-- BEGIN MUNGE: zed>", beginTag: true},
		mungeLine{data: "<!-- END MUNGE: zod>", endTag: true},
		mungeLine{data: "[ima](link)", link: true},
	}

	for i := range expected {
		ei := expected[i]
		li := l[i]
		if !reflect.DeepEqual(ei, li) {
			t.Errorf("(%v). Wanted %v: Found: %v", i, ei, li)
		}
	}
}
