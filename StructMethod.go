/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

type OfficeManagement struct {
	CTO, CEO, Karyawan string
	Absensi int
}

func(officeManagement OfficeManagement) sayHello(){
	fmt.Println("Hello, My Name is,", officeManagement.CEO)
}

func main()  {
	seno := OfficeManagement{CEO: "Seno"}
	seno.sayHello()
}