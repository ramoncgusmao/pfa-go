package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

/*
func main() {
	go contador("a")
	go contador("b")

	go contador("c")
	time.Sleep(time.Second * 12)
}
*/
