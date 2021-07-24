package main

import (
	"fmt"
	"net/http"

	"franco762/go-tickets-api/internal/database"
	"franco762/go-tickets-api/internal/ticket"
	transportHTTP "franco762/go-tickets-api/internal/transport/http"
)

// App : la estructura que contiene
// punteros a conexiones de base de datos
type App struct{}

// Run - Controla el inicio de la aplicación
func (app *App) Run() error {
	fmt.Println("Configurando APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	ticketService := ticket.NewService(db)

	handler := transportHTTP.NewHandler(ticketService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("No se pudo configurar el servidor")
		return err
	}

	return nil
}

func main() {
	fmt.Println("GO TICKETS API")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Errror al iniciar la API")
		fmt.Println(err)
	}
}
