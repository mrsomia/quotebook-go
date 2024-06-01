package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{
		next: next,
	}
}

func (s *LoggingService) GetQuote(ctx context.Context, id int) (q *Quote, err error) {
	defer func(start time.Time) {
		fmt.Printf("quote=%+v err=%v took=%v\n", q, err, time.Since(start))
	}(time.Now())
	return s.next.GetQuote(ctx, id)
}

func (s *LoggingService) CreateQuote(ctx context.Context, quote *Quote) (err error) {
	defer func(start time.Time) {
		fmt.Printf("err=%v took=%d\n", err, time.Since(start))
	}(time.Now())
	return s.next.CreateQuote(ctx, quote)
}
