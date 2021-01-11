/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import (
	"fmt"
	"strconv"
)

func main() {

	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("Boolean:", boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("10000", 10, 64)
	if err == nil {
		fmt.Println("Number", number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

}
