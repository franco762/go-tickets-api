package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/franco762/go-tickets-api/internal/transport/http"
)

// App : la estructura que contiene
// punteros a conexiones de base de datos
type App struct{}

// Run - Controla el inicio de la aplicación
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
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
