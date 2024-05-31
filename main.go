package main

import (
	"fmt"
)

func run() {
  fmt.Println("Running")
  svc := NewQuoteService(nil)
  svc = NewLoggingService(svc)

  apiServer := NewApiServer(svc)
  apiServer.Start(":8080")

}

func main() {
  run()
}
