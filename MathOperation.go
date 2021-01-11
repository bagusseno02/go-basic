/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main(){

	var a = 10
	var b = 100
	var c = a + b
	fmt.Println("Operasi Penjumlahan", c)

	var nilai1 = 10
	nilai1 += 10
	fmt.Println("Augmented Assignment", nilai1)

	var nilai2 = 10
	nilai2++
	fmt.Println("Unary Assignment", nilai2)

}

