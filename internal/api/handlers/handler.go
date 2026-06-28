package handlers

import (
	"log"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger *slog.Logger
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.logger.Info("")
		_, err := w.Write([]byte("System Up and Running on port 8000"))
		if err != nil {
			log.Println("ERROR: GET /health route failed. ", err)
		}
	}
}
