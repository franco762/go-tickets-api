package http

import (
	"encoding/json"
	"fmt"
	"franco762/go-tickets-api/internal/ticket"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Handler - Almacena los punteros al servicio de tickets
type Handler struct {
	Router  *mux.Router
	Service *ticket.Service
}

// Response - Almacena las respuestas de la api
type Response struct {
	Message string
	Error   string
}

// NewHandler - Retorna un puntero a un handler
func NewHandler(service *ticket.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - Configura todas las rutas de la aplicación
func (h *Handler) SetupRoutes() {
	fmt.Println("Configurando rutas")
	h.Router = mux.NewRouter()

	// Tickets rutas
	h.Router.HandleFunc("/api/ticket", h.GetAllTickets).Methods("GET")
	h.Router.HandleFunc("/api/ticket", h.PostTicket).Methods("POST")
	h.Router.HandleFunc("/api/ticket/{id}", h.GetTicket).Methods("GET")
	h.Router.HandleFunc("/api/ticket/{id}", h.UpdateTicket).Methods("PUT")
	h.Router.HandleFunc("/api/ticket/{id}", h.DeleteTicket).Methods("DELETE")

	// Ruta de prueba
	h.Router.HandleFunc("/api/test", func(w http.ResponseWriter, r *http.Request) {
		if err := sendOkResponse(w, Response{Message: "API funcionando"}); err != nil {
			panic(err)
		}
	})
}

// GetTicket - obtiene un ticket por ID
func (h *Handler) GetTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "No se puede convertir ID a UINT", err)
		return
	}

	ticket, err := h.Service.GetTicket(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error al obtener ticket por Id", err)
		return
	}

	if err := sendOkResponse(w, ticket); err != nil {
		panic(err)
	}
}

// GetAllTickets - obtiene todos los tickets
func (h *Handler) GetAllTickets(w http.ResponseWriter, r *http.Request) {
	tickets, err := h.Service.GetAllTickets()
	if err != nil {
		sendErrorResponse(w, "Error al obtener todos los tickets", err)
		return
	}

	if err := sendOkResponse(w, tickets); err != nil {
		panic(err)
	}
}

// PostTicket - crea un ticket
func (h *Handler) PostTicket(w http.ResponseWriter, r *http.Request) {
	var tic ticket.Ticket
	if err := json.NewDecoder(r.Body).Decode(&tic); err != nil {
		sendErrorResponse(w, "Error al decifrar JSON body", err)
		return
	}

	ticket, err := h.Service.PostTicket(tic)
	if err != nil {
		sendErrorResponse(w, "Error al crear ticket", err)
		return
	}

	if err := sendOkResponse(w, ticket); err != nil {
		panic(err)
	}
}

// UpdateTicket - actualiza un ticket por ID
func (h *Handler) UpdateTicket(w http.ResponseWriter, r *http.Request) {
	var tic ticket.Ticket
	if err := json.NewDecoder(r.Body).Decode(&tic); err != nil {
		sendErrorResponse(w, "Error al decifrar JSON body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	ticketID, err := strconv.ParseUint(id, 10, 60)
	if err != nil {
		sendErrorResponse(w, "No se puede convertir ID a UINT", err)
		return
	}

	ticket, err := h.Service.UpdateTicket(uint(ticketID), tic)
	if err != nil {
		sendErrorResponse(w, "Error al actualizar ticket", err)
		return
	}

	if err := sendOkResponse(w, ticket); err != nil {
		panic(err)
	}
}

// DeleteTicket - elimina un ticket por ID
func (h *Handler) DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	ticketID, err := strconv.ParseUint(id, 10, 60)
	if err != nil {
		sendErrorResponse(w, "No se puede convertir ID a UINT", err)
		return
	}

	err = h.Service.DeleteTicket(uint(ticketID))
	if err != nil {
		sendErrorResponse(w, "Error al eliminar ticket", err)
		return
	}

	if err = sendOkResponse(w, Response{Message: "Ticket eliminado satisfactoriamente"}); err != nil {
		panic(err)
	}
}

// Helpers
func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
