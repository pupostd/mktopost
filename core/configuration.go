package core

import (
	"log"
	"encoding/json"
	"io/ioutil"
)

type Property struct {
	Name string
	Value string
}

type Configuration struct {
	Properties []Property
}

func (conf *Configuration) Load() {

	b, err := ioutil.ReadFile("../properties.json")

	if err != nil {
		log.Fatalf("Could not find properties.json file at project root! Error: %T", err)
	}

	err = json.Unmarshal(b, &conf)

	if err != nil {
		log.Fatalf("Could not read properties.json properly! Error: %T", err)
	}

	for _, v := range conf.Properties {
		log.Printf("Property key: %s, value: %s", v.Name, v.Value)
	}
}
