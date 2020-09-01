package models

import (
	"github.com/kanwulf01/go-rest-api/src/database"
	"fmt"
	"log"
)

//Objeto pregunta
type Pregunta struct {
	ID          int    `json:"id"`
	NombreP string `json:"nombrep"`
	DescripcionP string `json:"descripcionp"`
	Respuesta string `json:"respuesta"`
}



//Insert Pregunta
func InsertPregunta(nombrep string, descripcionp string, respuesta string) (Pregunta, bool) {
	db := database.GetConnection()

	var pregunta_id int
	db.QueryRow("INSERT INTO pregunta(nombrep,descripcionp,respuesta) VALUES($1,$2,$3) RETURNING id", nombrep, descripcionp,respuesta).Scan(&pregunta_id)

	if pregunta_id == 0 {
		return Pregunta{}, false
	}
	fmt.Println(pregunta_id)
	return Pregunta{pregunta_id, nombrep,descripcionp,respuesta}, true
}

//Listar Preguntas

func GetAllPreguntas() ([]Pregunta) {
	db := database.GetConnection()

	rows, err := db.Query("SELECT * FROM pregunta")
	if err != nil{
		log.Fatal(err)
	}

	defer rows.Close()
	fmt.Println("rows")
	//fmt.Println(rows)
	var preguntas []Pregunta
	for rows.Next() {
		
		p := Pregunta{}

		var ID int
		var nombrep string
		var descripcionp string
		var respuesta string

		err := rows.Scan(&ID, &nombrep, &descripcionp, &respuesta)
		fmt.Println("err-scan")
		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(nombrep + descripcionp + respuesta)
		p.ID = ID
		p.NombreP = nombrep
		p.DescripcionP = descripcionp
		p.Respuesta = respuesta

		preguntas = append(preguntas, p)
	}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		return preguntas
}

//Lista Pregunta
func GetPregunta(id string) (Pregunta, bool) {
	db := database.GetConnection()

	row := db.QueryRow("SELECT * FROM pregunta where id = $1", id)
	fmt.Println(row)
	var ID int
	var nombrep string
	var descripcionp string
	var respuesta string

	err := row.Scan(&ID, &nombrep, &descripcionp, &respuesta)
	fmt.Println(err)
	if err != nil {
		return Pregunta{}, false
	}

	return Pregunta{ID,nombrep,descripcionp,respuesta}, true
	
}

//Lista pregunta y sus respuestas (join pregunta respuesta)
func GetPreguntaxRespuesta() (Pregunta,[]Respuesta) {
	db := database.GetConnection()

	rows, err := db.Query("SELECT * FROM pregunta p inner join respuesta r on p.id=r.idp")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pregunta := Pregunta{}
	var r []Respuesta
	
	for rows.Next() {

		respuestas := Respuesta{}

		var ID int
		var nombrep string
		var descripcionp string
		var respuesta string

		var IDS int
		var answer string
		var idp int

		err := rows.Scan(&ID,&nombrep,&descripcionp,&respuesta, &IDS, &answer, &idp)
		if err != nil {
			log.Fatal(err)
		}

		//Llena Objecto Pregunta

		pregunta.ID = ID
		pregunta.NombreP = nombrep
		pregunta.DescripcionP = descripcionp
		pregunta.Respuesta = respuesta

		//Llena arreglo de respuestas
		respuestas.ID = IDS
		respuestas.Respuesta = answer
		respuestas.Idp = idp

		r = append(r,  respuestas)
	}

	return pregunta,r
}