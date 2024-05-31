package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type ApiServer struct {
  svc Service
  mux *http.ServeMux
}

func NewApiServer(svc Service) *ApiServer {
  s := http.NewServeMux()
  apiServer := &ApiServer{
    svc: svc,
    mux: s,
  }
  return apiServer
}

func (s *ApiServer) Start(addr string) error {
  s.mux.HandleFunc("GET /quote", s.handleGetQuote)
  server := &http.Server{
    Addr: addr,
    Handler: s.mux,
  }
  return server.ListenAndServe()
}

func (s *ApiServer) handleGetQuote(w http.ResponseWriter, r *http.Request) {
  // TODO: read int from request
  quote, err := s.svc.GetQuote(context.Background(), 1)
  if err != nil {
    writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
    log.Fatal(err)
  }
  writeJSON(w, http.StatusOK, quote)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
  w.WriteHeader(s)
  w.Header().Add("Content-Type", "application/json")
  return json.NewEncoder(w).Encode(v)
}
