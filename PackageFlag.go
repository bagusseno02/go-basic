/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import (
	"flag"
	"fmt"
)

func main()  {

	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your username")
	password := flag.String("password", "root", "Put your password")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)

}

