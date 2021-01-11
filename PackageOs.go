/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import (
	"fmt"
	"os"
)

func main() {

	/* Get Os Argument */
	args := os.Args
	fmt.Println("Argument :", args)

	/* Get Hostname */
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println(err.Error())
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)

}
