package main

import (
	"context"
	"fmt"
	"log"
)

func run() {
	fmt.Println("Running")
	db, err := NewQuotePersistor(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	svc := NewQuoteService(db)
	svc = NewLoggingService(svc)

	apiServer := NewApiServer(svc)
	if err := apiServer.Start(":8080"); err != nil {
		fmt.Println(err)
	}

}

func main() {
	run()
}
