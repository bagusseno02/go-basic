/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter := 1; counter <= 11; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	/* For Range */
	slice := []string{"Bagus", "Seno", "Prasetyo"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println("index", index, "-", name)
	}

}
