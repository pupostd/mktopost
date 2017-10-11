package model

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//Searcher A class to search for a name on some folders.
type Searcher struct {
	filesFound []*os.File
}

// LookFor searches for a requested file.
func (s *Searcher) LookFor(name string) string {
	log.Printf("Received name to search %s", name)

	var file *os.File

	if s.search(name) {
		if len(s.filesFound) > 1 {
			c := Chooser{}
			file = c.WhichOne(s.filesFound)
		} else {
			file = s.filesFound[0]
		}

		return s.content(file)
	}

	return ""
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

func (s *Searcher) content(file *os.File) string {
	if strings.HasSuffix(file.Name(), ".md") {
		p := Parser{}
		return p.ToHTML(file)

	}

	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, file)

	if err != nil {
		log.Fatal("Problems reading file content.")
	}

	return string(buf.Bytes())
}
