package main

import (
	"context"
	"fmt"
	"log"
)

func run() {
  fmt.Println("Running")
  svc := NewQuoteService(nil)
  svc = NewLoggingService(svc)

  q, err := svc.GetQuote(context.TODO(), 1)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%+v\n", q)

}

func main() {
  run()
}
