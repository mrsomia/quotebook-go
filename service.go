package main

import "context"

type Service interface {
  GetQuote(context.Context, int) (*Quote, error)
  CreateQuote(context.Context, *Quote) error
}

type QuoteService struct {
  db Persistor
}

func NewQuoteService(db Persistor) Service {
  return &QuoteService{
    db: db,
  }
}

func (s *QuoteService) GetQuote(ctx context.Context, id int) (*Quote, error) {
  // TODO: add db intergiration
  return &Quote{Quote:"Progress over Pride", Author: "Lebron James"}, nil
}

func (s *QuoteService) CreateQuote(ctx context.Context, quote *Quote) error {
  // TODO: add db intergration
  return nil
}