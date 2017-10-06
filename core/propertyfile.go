package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const PropertiesFileName = "../properties"
const extensionJSON = ".json"
const extensionINI = ".ini"

type Property struct {
	Name  string
	Value string
}

type PropertiesFile struct {
	Properties []Property
}

func (f *PropertiesFile) LoadJSON(fileName string) {

	b, name := readFile(fileName, extensionJSON)
	err := json.Unmarshal(b, &f)

	if err != nil {
		log.Fatalf("Could not read %s properly! Error: %T", name, err)
	}

	for _, v := range f.Properties {
		log.Printf("Property key: %s, value: %s", v.Name, v.Value)
	}
}

// TODO: load an .ini property file
/*func (f *PropertiesFile) LoadPropertyINI(fileName string) {

	b, name := readFile(fileName, extensionINI)

	log.Printf(name)

	os.File{}


}*/

func readFile(fileName string, ext string) ([]byte, string) {
	name := PropertiesFileName + ext
	if fileName != "" {
		name = fileName
	}
	b, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatalf("Could not find %s file at project root! Error: %T", PropertiesFileName + extensionINI, err)
	}

	return b, name
}