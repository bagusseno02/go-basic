/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main(){

	var nilai1 int32 = 32768
	var nilai2 int64 = int64(nilai1)

	var nilai16 int16 = int16(nilai2)

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai16)

	var name = "Bagus Seno"
	var e byte =  name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)

}
