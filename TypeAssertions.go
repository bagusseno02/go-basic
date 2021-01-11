/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {

	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	//resultInt := result.(int)
	//fmt.Println(resultInt)

	results := random()
	switch value := results.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Bool", value)
	}

}
