package model

import (
	"os"
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

func (s *Storage) Save(file *os.File, here string) {}

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
