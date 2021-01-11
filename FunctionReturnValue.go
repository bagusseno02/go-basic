/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(getHello("Bagus Seno Prasetyo Diwiryo"))
}

