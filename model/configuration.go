package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const configurationFile = "../configuration.json"

type Configuration struct {
	Properties []Property
}

type Property struct {
	Name  string
	Value string
}

func (c *Configuration) For(name string) string {
	c.load()

	var found string
	for _, v := range c.Properties {
		if name == v.Name {
			found = v.Value
			break
		}
	}

	return found
}

func (c *Configuration) load() {
	b, name := c.readFileAsBuffer(configurationFile)
	err := json.Unmarshal(b, &c)

	if err != nil {
		log.Fatalf("Could not load %s properly! Error: %T", name, err)
	}

}

func (c Configuration) readFileAsBuffer(name string) ([]byte, string) {
	b, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatalf("Could not find %s file at project root! Error: %T", name, err)
	}

	return b, name
}
