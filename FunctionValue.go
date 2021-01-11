/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func getGoodBye(name string) string{
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodBye
	fmt.Println(goodbye("Bagus Seno"))
}
