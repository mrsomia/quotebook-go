package main

import (
	"fmt"
)

func run() {
	fmt.Println("Running")
	svc := NewQuoteService(nil)
	svc = NewLoggingService(svc)

	apiServer := NewApiServer(svc)
	if err := apiServer.Start(":8080"); err != nil {
		fmt.Println(err)
	}

}

func main() {
	run()
}
