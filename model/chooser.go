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
	for _, v := range files {
		chosen = v // TODO implement rule to choose
		break
	}
	return chosen
}
