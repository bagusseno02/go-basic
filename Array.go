/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	var names = [4]string{
		"Bagus",
		"Seno",
		"Prasetyo",
		"Diwiryo",
	}
	fmt.Println(names[0], names[1])

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(len(names))

}
