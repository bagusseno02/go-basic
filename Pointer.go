/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 = &address1

	address2.City = "Subang"

	fmt.Println(address1)
	fmt.Println(address2)

	/* Pass by reference dengan pointer */
	address3 := &address1

	address3.City = "Subang"

	fmt.Println(address1)
	fmt.Println(address3)

	address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	fmt.Println(address2)

	/* New pada pointer */
	var address4 = new(Address)
	fmt.Println(address4)
	address4.City = "Jakarta"
	fmt.Println(address4)

}
