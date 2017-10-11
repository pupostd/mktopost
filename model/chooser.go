package model

import (
	"os"
	"log"
)

type Chooser struct {

}

func (c *Chooser) WhichOne(files []*os.File) *os.File {
	if files != nil && len(files) > 1 {
		return c.choose(files)
	} else if files != nil {
		return files[0]
	} else {
		log.Fatalf("No file received to choose.")
		return nil
	}
}

func (c *Chooser) choose(files []*os.File) *os.File {
	var chosen *os.File
	var last int64 = 0

	for _, v := range files {
		info, err := v.Stat()
		if err != nil {
			log.Fatal("Could not read file stats to choose.")
		}

		changed := info.ModTime().Unix()
		if changed > last {
			last = changed
			chosen = v
		}
	}
	return chosen
}
