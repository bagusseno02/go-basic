/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	seno := Man{"Bagus Seno"}
	seno.Married()
	fmt.Println(seno.Name)
}
