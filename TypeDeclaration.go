/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main(){

	type noKTP string
	type married bool

	var ktpSeno noKTP = "3204140728930001"
	fmt.Println(ktpSeno)
	fmt.Println(noKTP("12345"))
	fmt.Println(married(true))

}

