/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
	Married bool
}

func main()  {

	var customer Customer
	customer.Name = "Bagus Seno"
	customer.Address = "Bandung"
	customer.Age = 27
	customer.Married = true
	fmt.Println(customer)

	/* Struct Literal */
	Cust := Customer{
		Name: "Joko",
		Address: "Jakarta",
		Age: 30,
	}
	fmt.Println(Cust)

}

