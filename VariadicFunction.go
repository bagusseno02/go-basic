/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers{
		total += number
	}
	return total
}

func main() {

	total := sumAll(10,10,10,10,10,10)
	fmt.Println(total)

	slice := []int{10,20,30,40,50}
	total = sumAll(slice...)
	fmt.Println(total)

}

