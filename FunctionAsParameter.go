/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Seno", filter)
}
