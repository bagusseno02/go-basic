/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	name := "Bagus Seno"

	if name == "Bagus Seno" {
		fmt.Println("Hello", name)
	} else if name == "Prasetyo" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Perkenalkan diri anda")
	}

	/* If Short Statement */
	var length = len(name)
	if length > 20 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

	if length := len(name); length > 20 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
