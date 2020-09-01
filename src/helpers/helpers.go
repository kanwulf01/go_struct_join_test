package helpers

import (
	"encoding/json"
	"github.com/kanwulf01/go-rest-api/src/database/models"
	"net/http"
	"strings"
)

func DecodeBody(req *http.Request) (models.Todo, bool) {
	var todo models.Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		return models.Todo{}, false
	}

	return todo, true
}

//Decodificacion del modelo Pregunta
func DecodeBodyPregunta(req *http.Request) (models.Pregunta, bool) {
	var pregunta models.Pregunta
	err := json.NewDecoder(req.Body).Decode(&pregunta)
	if err != nil {
		return models.Pregunta{},false
	}

	return pregunta, true
}

func DecodeBodyRespuesta(req *http.Request) (models.Respuesta, bool) {
	var respuesta models.Respuesta
	err := json.NewDecoder(req.Body).Decode(&respuesta)
	if err != nil {
		return models.Respuesta{}, false
	}

	return respuesta, true
}

func IsValidDescription(description string) bool {
	desc := strings.TrimSpace(description)
	if len(desc) == 0 {
		return false
	}

	return true
}

func IsValidFieldNumeric(field int) bool {
	if field == 0 {
		return false
	}

	return true
}

//Valida los campos de tipo string si son vacios
func IsValidField(field string) bool {
	desc := strings.TrimSpace(field)
	if len(desc) == 0 {
		return false
	}

	return true
}