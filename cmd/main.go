package main

import (
	"database/sql"
	"fmt"

	"github.com/ramoncgusmao/pfa-go/internal/order/infra/database"
	"github.com/ramoncgusmao/pfa-go/internal/order/usecase"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	calculate_price := usecase.NewCalculateFinalPriceUsecase(repository)
	input := usecase.OrderInputDTO{
		ID:    "159",
		Price: 100,
		Tax:   10,
	}
	order, err := calculate_price.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the final price is: %0.2f\n", order.FinalPrice)

}
