package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/ramoncgusmao/pfa-go/cmd/pkg/rabbitmq"
	"github.com/ramoncgusmao/pfa-go/internal/order/infra/database"
	"github.com/ramoncgusmao/pfa-go/internal/order/usecase"
)

func main() {
	maxWorkers := 6
	wg := sync.WaitGroup{}
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	calculate_price := usecase.NewCalculateFinalPriceUsecase(repository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, out)
	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go worker(out, calculate_price, i)
	}
	wg.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) error {
	for msg := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
			return nil
		}
		input.Tax = 10.0
		_, err = uc.Execute(input)

		if err != nil {
			fmt.Println("Error execute message", err)
			return nil
		}
		msg.Ack(false)
	}
	return nil
}
