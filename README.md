# Restful API basica de Tickets

API basica que permite crear, eliminar, actualizar y listar tickes, usando postgresql como gestor de base de datos, golang, y docker.

## Comenzando 🚀

Instrucciones de como configurar y correr el proyecto en local

### Pre-requisitos 📋

Para correr este proyecto es ecesario tener instalado Docker y Docker compose.

### Instalación 🔧

Para contruir la imagen docker ejecutar el siguiente comando:

```
docker build -t tickets-api . 
```

Para correr el proyaceto usar el comando de docker compose:

```
docker-compose up --build
```

## Construido con 🛠️

* [Golang](https://golang.org/) - El lenguaje de programación usado.
* [Gorm](https://gorm.io/) - Libreria ORM para base de datos.
* [Gorilla/Mux](https://github.com/gorilla/mux) - Enrutador HTTP y comparador de URL para construir servidores web Go.

## Autor ✒️

* **Fabian Franco** - *Desarrollo* - [Fabian Franco](https://github.com/franco762)
