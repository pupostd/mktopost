package core

import (
	"fmt"
	"os"
	"path/filepath"
)

//Searcher A class to search and read the resource requested.
type Searcher struct {
	placesToSearch []string
}

// SearchRequested searches for a requested file.
func (s *Searcher) SearchRequested(fileName string) string {

	s.loadPlacesToSearch()

	fmt.Printf("Received name to search %s\n", fileName)

	var fileNameFound string
	var requestedFile *os.File
	var err error

	for i := 0; i < len(s.placesToSearch); i++ {
		place := s.placesToSearch[i]
		requestedFile, err = os.Open(place + fileName)

		if requestedFile != nil {
			break
		}
	}

	if err != nil || requestedFile == nil {
		fmt.Printf(fmt.Errorf("File could not be found.\n").Error())
	}

	if requestedFile != nil {
		_, fileNameFound = filepath.Split(requestedFile.Name())
	}

	return fileNameFound
}

func (s *Searcher) addPlacesToSearch(path string) {
	s.placesToSearch = append(s.placesToSearch, path)
}

func (s *Searcher) loadPlacesToSearch() {
	s.placesToSearch = []string{"/home/pupo/"}
	s.addPlacesToSearch("/tmp/to_publish_hidden/")
	s.addPlacesToSearch("/tmp/")

	fmt.Printf("Length of placesToSearch is %d\n", len(s.placesToSearch))
}
