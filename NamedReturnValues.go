/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func getFullName() (firstname, lastname string){
	firstname = "Bagus"
	lastname = "Seno"
	return
}

func main() {
	firstname, lastname := getFullName()
	fmt.Println(firstname,lastname)
}

