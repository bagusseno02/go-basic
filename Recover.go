/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func endApps() {
	message := recover()
	if message != nil {
		fmt.Println("Error Dengan Message", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApps(error bool){
	defer endApps()
	if error{
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApps(true)
	fmt.Println("Bagus Seno")
}

