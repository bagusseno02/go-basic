/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	person := NewMap("Bagus")
	fmt.Println(person)

	var persons = NewMap("Bagus")
	if persons == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(persons)
	}

}
