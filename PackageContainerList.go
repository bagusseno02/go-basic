/**
 * @author <a href="mailto:bagus.seno39@gmail.com>seno</a>
 * Created on 24/11/20
 * Project go-basic
 */

package main

import (
	"container/list"
	"fmt"
)

func main()  {
	data := list.New()
	data.PushBack("Bagus")
	data.PushBack("Seno")
	data.PushBack("Prasetyo")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for e := data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}

