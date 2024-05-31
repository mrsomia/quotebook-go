package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	// INFO: move this router config to an add routes function
	apiServer.mux.HandleFunc("POST /quote", apiServer.handleGetQuote)
	return apiServer
}

func (s *ApiServer) Start(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}
	return server.ListenAndServe()
}

func (s *ApiServer) handleGetQuote(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	b := struct {
		ID int `json:"id"`
	}{}

	if err := d.Decode(&b); err != nil {
		s := fmt.Sprintf("Unable to read JSON\nerror: %v", err.Error())
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": s})
		return
	}

	quote, err := s.svc.GetQuote(context.Background(), b.ID)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, quote)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
