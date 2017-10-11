package model

import (
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWhere(t *testing.T) {
	s := Storage{}
	folders := s.Where()
	if folders == nil {
		t.Error("Empty array! Some paths should be returned.")
	}
}

func TestSave(t *testing.T) {

	os.Remove("../resource/test/post_one.md.html")

	f, err := os.Open("../resource/md/post_one.md")
	if err != nil {
		t.Error("Could not load file md for testing save.")
	}

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, f)
	if err != nil {
		t.Error("Could not transform markdown to html.")
	}

	out := blackfriday.Run(buf.Bytes())
	name := filepath.Base(f.Name())
	name = (strings.Split(name, "."))[0] + ".html"

	s := Storage{}
	s.Save(name, out, "../resource/test/")
	f.Close()
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
