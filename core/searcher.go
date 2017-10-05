package core

import "fmt"

//Searcher A class to search and read the resource requested.
type Searcher struct {
	placesToSearch []string
}

// SearchRequested searches for a requested file.
func (s Searcher) SearchRequested(fileName string) string {
	s.loadPlacesToSearch()

	fmt.Printf("Received name to search %s\n", fileName)

	return fileName
}

func (s Searcher) addPlacesToSearch(path string) {
	s.placesToSearch = append(s.placesToSearch, path)
}

func (s Searcher) loadPlacesToSearch() {
	//s.placesToSearch := [1]string{"searcher.properties"}
	s.addPlacesToSearch("to_publish")
	s.addPlacesToSearch("to_publish_hidden")
}
