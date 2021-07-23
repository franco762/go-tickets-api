package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - Almacena los punteros al servicio de tickets
type Handler struct {
	Router *mux.Router
}

// NewHandler - Retorna un puntero a un handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - Configura todas las rutas de la aplicación
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}
