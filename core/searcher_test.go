package core

import (
	"testing"
)

func TestLookFor(t *testing.T) {
	searcher := Searcher{}
	post := searcher.LookFor("post_one")
	if post == "" {
		t.Errorf("Search not done well.")
	}

}
