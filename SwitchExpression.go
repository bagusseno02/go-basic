/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	name := "Seno"

	switch name {
	case "Seno":
		fmt.Println("Hello", name)
	case "Bagus":
		fmt.Println("Hello Bagus")
	default:
		fmt.Println("Boleh Kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

}
