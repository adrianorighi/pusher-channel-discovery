package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type Server struct {
	HostName string
	Port     int64
}

func Start(srv *Server) error {

	mux := chi.NewMux()

	mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Health check passes`))
	})

	mux.Post("/generate", func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()

		var response struct {
			Status int64     `json:"status"`
			ID     uuid.UUID `json:"id"`
		}

		response.Status = 1
		response.ID = uuid.New()

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", srv.Port), mux)
}
