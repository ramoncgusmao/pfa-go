package main

import "time"

func goroutinex(ch chan string) {
	time.Sleep(time.Second * 5) // se você colocar antes ele espera se voce atribuir depois do canal ele não espera
	ch <- "Olá mundo!"
}

/*
func main() {
	canal := make(chan string) // criação de um canal

	go goroutinex(canal)

	result := <-canal
	println(result)
}
*/
