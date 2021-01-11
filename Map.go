/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	person := map[string] string{
		"name":    "Bagus Seno",
		"address": "Bandung",
	}

	person["title"] = "Backend Developer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))

	var book = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Bagus Seno Prasetyo"
	book["ups"] = "salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

}
