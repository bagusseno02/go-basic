/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func getFullname()(string, string){
	return "Bagus Seno", "Prasetyo Diwiryo"
}

func getFullname2()(string, string){
	return "Bagus Seno", "Prasetyo"
}

func main() {

	firstname, lastname := getFullname()
	fmt.Println(firstname, lastname)

	/* Function multiple without return value */
	firstName, _ := getFullname2()
	fmt.Println(firstName)

}

