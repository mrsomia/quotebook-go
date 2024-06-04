package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func run(ctx context.Context, w io.Writer, getEnv func(string) string) error {
	fmt.Println("Running")
	port := getEnv("PORT")
	if port == "" {
		port = ":8080"
	}

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	db, err := NewQuotePersistor(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}

	svc := NewQuoteService(db)
	svc = NewLoggingService(svc)

	srv := NewApiServer(svc, port)

	go func() {
		fmt.Printf("Listening and serving on %s\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	// attempts a graceful shutdown
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), time.Second*10)
		defer shutdownCtxCancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(w, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
