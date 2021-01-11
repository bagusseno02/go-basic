/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 23/11/20
 * Project go-basic
 */

package main

import "fmt"

func main() {

	/* Break */
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke -", i)
	}

	/* Continue */
	for j := 0; j < 20; j++ {
		if j % 2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke -", j)
	}

}
