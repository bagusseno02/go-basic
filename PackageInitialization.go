/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import (
	"fmt"
	"go-basic/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
