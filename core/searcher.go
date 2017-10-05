package core

type Searcher struct {
  placesToSearch []string
}

func (s Searcher) SearchRequested(fileName string) string {
  s.loadPlacesToSearch()
  return fileName
}

func (s Searcher) addPlacesToSearch(path string) {
  s.placesToSearch = append(s.placesToSearch, path)
}

func (s Searcher) loadPlacesToSearch() {
  //s.placesToSearch := [1]string{"searcher.properties"}
  s.addPlacesToSearch("searcher.properties")
  s.addPlacesToSearch("default.properties")
}
