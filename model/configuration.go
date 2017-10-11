package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const configurationFile = "../configuration.json"

// Configuration A class that is responsible to load the config file properties.
type Configuration struct {
	Properties []Property
}

// Property type is responsible for getting property name and value.
type Property struct {
	Name  string
	Value string
}

// For reads the value for a property name given.
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
