package cmd

import "testing"

func TestExtractProjectNameFromPath(t *testing.T) {
	testdata := []struct {
		giveString string
		wontString string
	}{
		{"foo.sqlite", "foo"},
		{"/home/bar/.kani/kani.sqlite", "bar"},
		{"/home/bar/.kani/baz.sqlite", "baz"},
		{"/home/bar/qux.sqlite", "qux"},
		{"/home/bar/kani.sqlite", "unknown"},
	}
	for _, td := range testdata {
		gotString := extractProjectNameFromPath(td.giveString)
		if gotString != td.wontString {
			t.Errorf("extractProjectNameFromPath(%s) did not match, wont %s got %s", td.giveString, td.wontString, gotString)
		}
	}
}
