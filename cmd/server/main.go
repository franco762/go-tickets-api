package main

import "fmt"

// App : la estructura que contiene
// punteros a conexiones de base de datos
type App struct{}

// Run - Controla el inicio de la aplicación
func (app *App) Run() error {
	fmt.Println("Configurando aplicación")
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
