 
package api

import (
	"encoding/json"
	"github.com/kanwulf01/go-rest-api/src/helpers"
	"github.com/kanwulf01/go-rest-api/src/database/models"
	"net/http"
	"strings"
	"fmt"
	"github.com/gorilla/mux"
)

type Data struct {
	Success bool          `json:"success"`
	Data    []models.Todo `json:"data"`
	Errors  []string      `json:"errors"`
}

type DataPregunta struct {
	Success bool          `json:"success"`
	DataPregunta    []models.Pregunta `json:"pregunta"`
	Errors  []string      `json:"errors"`
}

type DataRespuesta struct {
	Success bool          `json:"success"`
	DataRespuesta    []models.Respuesta `json:"respuesta"`
	Errors  []string      `json:"errors"`
}

//Objeto que empaqueta 1 objeto de tipo pregunta y un arreglo de objetos respuestas
type ResponseJoin struct {
	Pregunta models.Pregunta `json:"pregunta"`
	Respuestas []models.Respuesta `json:"respuestas"`

}


//Serializador POST de Pregunta con todos y cada uno de sus campos
func CreatePregunta(w http.ResponseWriter, req *http.Request) {
	bodyTodo, success := helpers.DecodeBodyPregunta(req)
	if success != true {
		

		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data DataPregunta = DataPregunta{Errors: make([]string, 0)}
	bodyTodo.NombreP = strings.TrimSpace(bodyTodo.NombreP)
	bodyTodo.DescripcionP = strings.TrimSpace(bodyTodo.DescripcionP)
	bodyTodo.Respuesta = strings.TrimSpace(bodyTodo.Respuesta)
	fmt.Println(bodyTodo.Respuesta)
	if !helpers.IsValidField(bodyTodo.DescripcionP) && !helpers.IsValidField(bodyTodo.NombreP) && !helpers.IsValidField(bodyTodo.Respuesta){
		
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	pregunta, success := models.InsertPregunta(bodyTodo.NombreP, bodyTodo.DescripcionP, bodyTodo.Respuesta)
	fmt.Println(success)
	fmt.Println(pregunta)
	if success != true {
		data.Errors = append(data.Errors, "could not create todo")
	}

	data.Success = success
	data.DataPregunta = append(data.DataPregunta, pregunta)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

//Serializador de Todo
func CreateTodo(w http.ResponseWriter, req *http.Request) {
	bodyTodo, success := helpers.DecodeBody(req)
	if success != true {
		

		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	//fmt.Println(success)

	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	fmt.Println(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	fmt.Println("AAAA")
	todo, success := models.Insert(bodyTodo.Description)
	fmt.Println(success)
	if success != true {
		data.Errors = append(data.Errors, "could not create todo")
	}

	data.Success = success
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}


func GetTodos(w http.ResponseWriter, req *http.Request) {
	var todos []models.Todo = models.GetAll()
	fmt.Println(todos)
	var data = Data{true, todos, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//Funcion que retorna un Get con 1 Preguntas y todas sus diferentes Respuestas
func GetPreguntaXrespuestas(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request Get")
	var pregunta models.Pregunta
	var respuestas []models.Respuesta 
	pregunta, respuestas = models.GetPreguntaxRespuesta()

	var res = ResponseJoin{pregunta, respuestas}
	json, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)


	fmt.Println(pregunta)
	fmt.Println(respuestas)
}

//GET A PREGUNTAS
func GetPreguntas(w http.ResponseWriter, req *http.Request) {
	var preguntas []models.Pregunta = models.GetAllPreguntas()
	fmt.Println(preguntas)
	var data = DataPregunta{true, preguntas, make([]string,0)}
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}


/*
func UpdateTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	todo_id := vars["id"]

	bodyTodo, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	todo, success := models.Update(todo_id, bodyTodo.Description)
	if success != true {
		data.Errors = append(data.Errors, "could not update todo")
	}

	data.Success = success
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}
*/

func GetTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var todo models.Todo
	var success bool
	todo, success = models.Get(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "todo not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

/*
func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data = Data{Errors: make([]string, 0)}

	todo, success := models.Delete(id)
	if success != true {
		data.Errors = append(data.Errors, "could not delete todo")
	}

	data.Success = success
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
*/

func GetPregunta(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data DataPregunta

	var pregunta models.Pregunta
	var success bool
	pregunta, success = models.GetPregunta(id)
	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "todo not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.DataPregunta = append(data.DataPregunta, pregunta)
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func CreateRespuesta(w http.ResponseWriter, req *http.Request) {
	bodyRequest, success := helpers.DecodeBodyRespuesta(req)
	fmt.Println("success")
	fmt.Println(success)
	if success != true {
		

		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	
	var data DataRespuesta = DataRespuesta{Errors: make([]string, 0)}
	bodyRequest.Respuesta = strings.TrimSpace(bodyRequest.Respuesta)
	
	fmt.Println(bodyRequest.Respuesta)
	if !helpers.IsValidField(bodyRequest.Respuesta) && !helpers.IsValidFieldNumeric(bodyRequest.Idp) {
		
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	respuesta, success := models.InsertRespuesta(bodyRequest.Respuesta, bodyRequest.Idp)
	fmt.Println(success)
	fmt.Println(respuesta)
	if success != true {
		data.Errors = append(data.Errors, "could not create todo")
	}

	data.Success = success
	data.DataRespuesta = append(data.DataRespuesta, respuesta)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}