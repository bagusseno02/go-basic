/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func blackList1(name string) bool {
	return name == "anjing"
}

func main() {
	blackList := blackList1
	registerUser("anjing", blackList)
}
