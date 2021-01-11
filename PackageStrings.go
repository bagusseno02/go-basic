/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Bagus Seno", "Prasetyo"))
	fmt.Println(strings.Split("Bagus Seno", " "))
	fmt.Println(strings.ToLower("Bagus Seno Prasetyo"))
	fmt.Println(strings.ToUpper("Bagus Seno Prasetyo"))
	fmt.Println(strings.Trim("   Bagus Seno  ", " "))
	fmt.Println(strings.ReplaceAll("Bagus Bagus Bagus Bagus", "Bagus", "Seno"))
}

