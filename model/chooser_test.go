package model

import (
	"testing"
	"os"
)

func TestWhichOne(t *testing.T)  {
	c := Chooser{}
	var files []*os.File = nil

	f1, err := os.Open("../resource/test/dummy_one.txt")
	if err != nil {
		t.Error("Failed to open file for test")
	}

	f2, err := os.Open("../resource/test/dummy_one.txt")
	if err != nil {
		t.Error("Failed to open file for test")
	}

	files = append(files, f1, f2)

	c.WhichOne(files)
}
