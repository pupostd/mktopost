package model

import (
	"fmt"
	"log"
	"os"
)

//Searcher A class to search for a name on some folders.
type Searcher struct {
	filesFound []*os.File
}

// LookFor searches for a requested file.
func (s *Searcher) LookFor(name string) string {
	log.Printf("Received name to search %s", name)

	var content string

	if s.search(name) {
		p := Parser{}
		if len(s.filesFound) > 1 {
			c := Chooser{}
			content = p.ToHTML(c.WhichOne(s.filesFound))
		} else {
			content = p.ToHTML(s.filesFound[0])
		}
	}

	return content
}

func (s *Searcher) search(name string) bool {
	var file *os.File
	var err error

	st := Storage{}
	folders := st.Where()

	for i := 0; i < len(folders); i++ {
		place := folders[i]
		file, err = os.Open(place + "/" + name + st.FolderContent(place))

		if file != nil {
			s.filesFound = append(s.filesFound, file)
		}
	}

	if err != nil && s.filesFound == nil {
		log.Printf(fmt.Errorf("file could not be found").Error())
		return false
	}

	return true
}
