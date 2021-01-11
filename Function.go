/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloSeno() {
	fmt.Println("Hello Seno")
}

func main() {
	sayHello()
	for i := 0; i < 10; i++ {
		sayHelloSeno()
	}
}
