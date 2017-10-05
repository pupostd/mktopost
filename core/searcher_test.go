package core

import (
	"testing"
)

func Test(t *testing.T) {
	searcher := Searcher{}
	post := searcher.SearchRequested("post_one.md")
	if post != "post_one.md" {
		t.Errorf("Search not done well.")
	}

}
