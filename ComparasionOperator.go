/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main(){

	var firstname = "Bagus Seno"
	var lastname = "Prasetyo"

	var result = firstname == lastname
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 > value2)

}

