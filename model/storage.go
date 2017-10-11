package model

import (
	"io/ioutil"
	"log"
	"strings"
)

const storageFolders = "storage.folders"

// Storage A class that is responsible to know the folders to store files.
type Storage struct {
	folders []string
}

// Where return a list of folders where files are storaged.
func (s *Storage) Where() []string {
	conf := Configuration{}
	s.folders = strings.Split(
		conf.For(storageFolders), ";")

	return s.folders
}

// Save a file in the storage using a _here_ path
func (s *Storage) Save(name string, content []byte, here string) {
	err := ioutil.WriteFile(here+"/"+name, []byte(content), 0644)

	if err != nil {
		log.Fatalf("Could not write file %s to %s", name, here)
	}
}

// FolderContent grasp the type of file that should be on a folder based on its name
// eg. ../resource/md/ folder should have .md files.
// so, a ../resource/md/myfile.md will as name will return .md here.
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
