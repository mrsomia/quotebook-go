package main

import "context"

type Persistor interface {
  CreateQuote(context.Context, *Quote) error
}

type QuotePersistor struct {
  db map[int]Quote
}

func NewQuotePersistor() *QuotePersistor {
  return &QuotePersistor{db:make(map[int]Quote),}
}

