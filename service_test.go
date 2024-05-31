package main

import (
	"context"
	"testing"
)

func TestQuoteFromGetQuote(t *testing.T) {
  expectedQuote := "Progress over Pride"
	svc := NewQuoteService(nil)
	q, err := svc.GetQuote(context.Background(), 1)
	if err != nil {
    t.Fatalf("Failed, error from GetQuote:\n%v\n", err.Error())
  }
  if q.Quote != expectedQuote{
    t.Fatalf("Expected quote to be:\n%v\nFound:\n%v", expectedQuote, q.Quote)
  }
}

func TestAuthorFromGetQuote(t *testing.T) {
  expectedAuthor := "Lebron James"
	svc := NewQuoteService(nil)
	q, err := svc.GetQuote(context.Background(), 1)
	if err != nil {
    t.Fatalf("Failed, error from GetQuote:\n%v\n", err.Error())
  }
  if q.Author != expectedAuthor{
    t.Fatalf("Expected quote to be:\n%v\nFound:\n%v", expectedAuthor, q.Author)
  }
}
