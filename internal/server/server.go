package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abhinavdevarakonda/maplet/internal/export"
)

type Server struct {
	graph export.Graph
}

func New(graph export.Graph) *Server {
	return &Server{graph: graph}
}

func(s *Server) Start(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/graph", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// it either returns (error) nil or some non nil value
		if err := json.NewEncoder(w).Encode(s.graph); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	mux.Handle("/", http.FileServer(http.Dir("frontend")))

	log.Printf("Serving maplet server on http://%s\n", addr)

	return http.ListenAndServe(addr, mux)
}
