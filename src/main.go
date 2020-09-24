package main

import (
	"fmt"
	//"github.com/kanwulf01/go-rest-api/src/database"
	"github.com/gorilla/mux"
	"github.com/kanwulf01/go-rest-api/src/api"
	"net/http"
)



func main() {
	fmt.Println("Hola Mundo")
	//db := database.GetConnection()

	var port string = "8080"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods("POST")
	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods("GET")
	apiRouter.HandleFunc("/todoslista/", api.GetTodos).Methods("GET")
	apiRouter.HandleFunc("/createPregunta", api.CreatePregunta).Methods("POST")
	apiRouter.HandleFunc("/listaPreguntas", api.GetPreguntas).Methods("GET")
	apiRouter.HandleFunc("/pregunta/{id}", api.GetPregunta).Methods("GET")
	apiRouter.HandleFunc("/createRespuesta", api.CreateRespuesta).Methods("POST")
	apiRouter.HandleFunc("/getPxR", api.GetPreguntasXrespuestas).Methods("GET")
	apiRouter.HandleFunc("/getPreguntas", api.GetPreguntaXrespuestas1).Methods("GET")
	apiRouter.HandleFunc("/correcta/{idp}/{res}", api.CorrectAnswers).Methods("GET")

	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":" + port, router)

}

