package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type people struct {
	People []string `yaml:"people"`
}

func (p *people)readinputpeople(file string) *people {
	peopleFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading people file:", err)
	}
	err = yaml.Unmarshal(peopleFile, p)
	if err != nil {
		fmt.Println("Error Unmarshalling people file:", err)
	}

	return p

}

func main() {
  var p people
	p.readinputpeople("settings/people.yml")
	fmt.Println(p)

}
