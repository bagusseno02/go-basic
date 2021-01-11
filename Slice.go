/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	months := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	var slice1 = months[2:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Bagus")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Bagus"
	newSlice[1] = "Seno"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([] string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
