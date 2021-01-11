/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import "fmt"

type Address2 struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address2) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address2{
		City:     "Ciamis",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address2 = &alamat
	ChangeAddressToIndonesia(alamatPointer)
	fmt.Println(alamatPointer)
}
