/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("Error")
	}
}

func main()  {
	runApp(true)
}


