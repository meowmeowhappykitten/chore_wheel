package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type people struct {
	People []string `yaml:"people"`
}

type chore struct {
	Description string `yaml:"description"`
	Frequency   string `yaml:"frequency"`
}

type chores struct {
	Chores []map[string]chore
}

func (p *people) readinputpeople(file string) *people {
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

func (c *chores) readinputchores(file string) *chores {
	choreFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading chores file:", err)
	}
	err = yaml.Unmarshal(choreFile, c)
	if err != nil {
		fmt.Println("Error Unmarshalling chores file:", err)
	}

	return c

}

func main() {
	var p people
	var c chores
	p.readinputpeople("settings/people.yml")
	c.readinputchores("settings/chores.yml")
	fmt.Println(p)
	fmt.Println(c)

}
