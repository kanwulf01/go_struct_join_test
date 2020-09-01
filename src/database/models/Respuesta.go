package models



import (
	"github.com/kanwulf01/go-rest-api/src/database"
	"fmt"
	"log"
)

type Respuesta struct {
	ID          int    `json:"id"`
	Respuesta string `json:"respuesta"`
	Idp 		int 	`json:"idp"`
}

//Insert Respuesta
func InsertRespuesta(respuesta string, idp int) (Respuesta, bool) {
	db := database.GetConnection()

	var respuesta_id int
	db.QueryRow("INSERT INTO respuesta(respuesta,idp) VALUES($1,$2) RETURNING id", respuesta, idp).Scan(&respuesta_id)

	if respuesta_id == 0 {
		return Respuesta{}, false
	}
	fmt.Println(respuesta_id)
	return Respuesta{respuesta_id,"",0}, true
}

//Lista todas las respuestas
func GetAllRespuestas() ([]Respuesta) {
	db := database.GetConnection()

	rows, err := db.Query("SELECT * FROM pregunta")
	if err != nil{
		log.Fatal(err)
	}

	defer rows.Close()
	fmt.Println("rows")
	//fmt.Println(rows)
	var respuestas []Respuesta
	for rows.Next() {
		
		r := Respuesta{}

		var ID int
		var respuesta string
		var idp int

		err := rows.Scan(&ID, &respuesta, &idp)
		fmt.Println("err-scan")
		fmt.Println(err)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(nombrep + descripcionp + respuesta)
		r.ID = ID
		r.Respuesta = respuesta
		r.Idp = idp

		respuestas = append(respuestas, r)
	}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		return respuestas
}