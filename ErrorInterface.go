/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
	var contohError = errors.New("Ups Error")
	fmt.Println(contohError.Error())
}
