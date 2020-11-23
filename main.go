package main

import (
	"golang/structure"
)

func main() {
	person := &structure.Person{}
	person.Name = "胡康"
	person.Age = 29
	person.Gender = 1
	person.FmtPersion()

	var i structure.Methods
	i = person
	i.FmtPersion()
}
