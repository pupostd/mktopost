package model

import (
	"io/ioutil"
	"log"
	"strings"
)

const storageFolders = "storage.folders"

type Storage struct {
	folders []string
}

func (s *Storage) Where() []string {
	conf := Configuration{}
	s.folders = strings.Split(
		conf.For(storageFolders), ";")

	return s.folders
}

func (s *Storage) Save(name string, content []byte, here string) {
	err := ioutil.WriteFile(here+"/"+name, []byte(content), 0644)

	if err != nil {
		log.Fatalf("Could not write file %s to %s", name, here)
	}
}

func (s Storage) FolderContent(name string) string {
	var suffix string
	cut := strings.Split(name, "/")

	l := len(cut)
	if l > 0 {
		suffix = cut[l-1]

		if suffix == "" && ((l - 2) >= 0) {
			suffix = cut[len(cut)-2]
		}

		suffix = "." + suffix
	}

	return suffix
}
