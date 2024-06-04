package main

import (
	"context"
)

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
	q, err := s.db.GetQuoteByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Quote{Quote: q.Content, Author: q.Name, DateAdded: q.DateAdded}, nil
}

func (s *QuoteService) CreateQuote(ctx context.Context, quote *Quote) error {
	// TODO: add db intergration
	return nil
}
