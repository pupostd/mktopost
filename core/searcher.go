package core

import (
	"fmt"
	"log"
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

	log.Printf("Received name to search %s", fileName)

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
		log.Printf(fmt.Errorf("file could not be found").Error())
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

	conf := PropertiesFile{}
	conf.LoadJSON("")

	for _, v := range conf.Properties {
		s.addPlacesToSearch(v.Value)
	}

	log.Printf("Length of placesToSearch is %d", len(s.placesToSearch))
}
