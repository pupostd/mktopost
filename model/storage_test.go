package model

import "testing"

func TestWhere(t *testing.T) {
	s := Storage{}
	folders := s.Where()
	if folders == nil {
		t.Error("Empty array! Some paths should be returned.")
	}
}

func TestFolderContent(t *testing.T) {
	s := Storage{}
	md := s.FolderContent("../resource/md/")
	if md != ".md" {
		t.Error("Error on getting suffix for folder")
	}

	md = s.FolderContent("../resource/md")
	if md != ".md" {
		t.Error("Error on getting suffix for folder")
	}
}
