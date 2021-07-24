# Restful API basica de Tickets

API basica que permite crear, eliminar, actualizar y listar tickets, usando postgresql como gestor de base de datos, golang, y docker.

## Comenzando 🚀

Instrucciones de como configurar y correr el proyecto en local

### Pre-requisitos 📋

Para correr este proyecto es necesario tener instalado Docker y Docker compose.

### Instalación 🔧

Para construir la imagen docker ejecutar el siguiente comando:

```
docker build -t tickets-api . 
```

Para correr el proyecto usar el comando de docker compose:

```
docker-compose up --build
```

## Rutas de la API 🚀

Listado de todas las rutas de la api.

### Ruta de prueba (GET)

```
http://localhost:8080/api/test
```

### Crear Ticket (POST)

```
http://localhost:8080/api/ticket
```
el request body para la crear un tricket es el siguiente:

```
{
    "user": "string",
    "status": "string"
}
```

### Obtener un Ticket (GET)

```
http://localhost:8080/api/ticket/{id}
```

### Obtener todos los Tickets (GET)

```
http://localhost:8080/api/ticket/
```

### Eliminar un Ticket (DELETE)

```
http://localhost:8080/api/ticket/{id}
```

### Actualizar Ticket (PUT)

```
http://localhost:8080/api/ticket
```
el request body para la actualizar un tricket es el siguiente:

```
{
    "status": "string"
}
```

## Construido con 🛠️

* [Golang](https://golang.org/) - El lenguaje de programación usado.
* [Gorm](https://gorm.io/) - Libreria ORM para base de datos.
* [Gorilla/Mux](https://github.com/gorilla/mux) - Enrutador HTTP y comparador de URL para construir servidores web Go.

## Autor ✒️

* **Fabian Franco** - *Desarrollo* - [Fabian Franco](https://github.com/franco762)
