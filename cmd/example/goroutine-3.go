package main

import "time"

func worker(workerId int, msg chan int) {
	for res := range msg {
		println("worker", workerId, "recebeu", res)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int) // criação de um canal
	for i := 0; i < 5; i++ {
		go worker(i, canal)
	}

	for i := 0; i < 30; i++ {
		canal <- i
	}
}
