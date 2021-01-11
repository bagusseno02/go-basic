/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

type HasName interface {
	GetName() string
}

func sayWelcome(hasName HasName){
	fmt.Println("Welcome,", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main()  {
	var seno Person
	seno.Name = "Bagus Seno"
	var hewan Animal
	hewan.Name = "Anjing"
	sayWelcome(seno)
	sayWelcome(hewan)
}