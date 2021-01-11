/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func factorialLoop(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialLoop(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
}
