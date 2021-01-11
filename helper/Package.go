/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */
package helper

import "fmt"

/* Access modifier cukup dengan huruf besar kecil di depan penamaan variable */
var version = 1
var Application = "Belajar Golang Dasar"

func SayWelcome(name string) string {
	return "Hello " + name
}

func sayWelcome(name string)  {
	fmt.Println("Hello", name)
}
