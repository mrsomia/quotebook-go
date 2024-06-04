package main

import (
	"context"
	"database/sql"
	_ "embed"

	_ "github.com/mattn/go-sqlite3"

	"GitHub.com/mrsomia/quotes/dbqueries"
)

type Persistor interface {
	CreateQuote(context.Context, *Quote) error
}

type QuotePersistor struct {
	db *dbqueries.Queries
}

// go:embed schema.sql
var ddl string

func NewQuotePersistor(ctx context.Context) (*QuotePersistor, error) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}

	queries := dbqueries.New(db)
	return &QuotePersistor{db: queries}, nil
}

func (p *QuotePersistor) CreateQuote(ctx context.Context, quote *Quote) error {
	// TODO
	return nil
}

func (p *QuotePersistor) GetQuoteByID(ctx context.Context, id int) (*Quote, error) {
	return nil, nil
}
