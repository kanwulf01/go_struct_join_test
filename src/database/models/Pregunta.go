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

//Objeto que empaqueta 1 objeto de tipo pregunta y un arreglo de objetos respuestas
type ResponseJoin struct {
	Pregunta Pregunta `json:"pregunta"`
	Respuestas []Respuesta `json:"respuestas"`

}

type JoinNormal struct {
	Pregunta Pregunta `json:"pregunta"`
	Respuesta Respuesta `json:"respuesta"`
}

type Response struct {

	Array ResponseJoin `json:"array"`

}

type ResponseNormal struct {
	Arrays JoinNormal `json:"arrays"`
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
	sizerows, err := db.Query("select count(*) from (SELECT * FROM pregunta) as subconsul;")
	var count int
	var c int
	for sizerows.Next() {
		sizerows.Scan(&count)
		fmt.Println(c)
	}
	
	defer sizerows.Close()
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
//Para listar todas las Preguntas y cada una con su respuesta, debo meter todo 
//dentro de un arreglo pregunta y respuestas
//Lista pregunta y sus respuestas (join pregunta respuesta)
func GetPreguntasxRespuesta() ([]Response) {
	db := database.GetConnection()

	sizerows, err := db.Query("select count(*) from (select * from pregunta p inner join respuesta r on p.id = r.idp) as x")
	var count int
	var limit int
	for sizerows.Next() {
		sizerows.Scan(&count)
		fmt.Println(count)
	}
	
	defer sizerows.Close()	

	rows, err := db.Query("SELECT * FROM pregunta p inner join respuesta r on p.id=r.idp")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	pregunta := Pregunta{}
	responsejoin := ResponseJoin{}//encapsula pregunta y arreglo respuesta
	res := Response{} //encapsula arreglo responsejoin
	var respuestasArray []Respuesta
	//var responsejoinArray []ResponseJoin
	var resArray []Response
	
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
		limit++
		fmt.Println("limit")
		fmt.Println(limit)

		//Llena Objecto Pregunta solo si no esta el id repetido,
		//es decir descarta filas de pregunta repetidas por el join
		//id 1 == ID 1
		//asigne valores Objeto Pregunta, Objecto respuesta
		//guarde objeto Respuesta en arreglo de respuesta
		
		if (pregunta.ID == ID || pregunta.ID == 0) && limit != count {
			fmt.Println("Pregunta repetida")
			pregunta.ID = ID
			pregunta.NombreP = nombrep
			pregunta.DescripcionP = descripcionp
			pregunta.Respuesta = respuesta

			//Llenar arreglo de respuestas por esta pregunta
			respuestas.ID = IDS
			respuestas.Respuesta = answer
			respuestas.Idp = idp

			respuestasArray = append(respuestasArray, respuestas)
		}else if pregunta.ID != ID && limit == count {
			fmt.Println("Pregunta diferente y final")
			//antes de borrar guardar esto dentro del empaquetador la clase responsejoin
			responsejoin.Pregunta = pregunta
			responsejoin.Respuestas = respuestasArray
			//despues guardar responsejoin en la clase empaquetadora res
			res.Array = responsejoin
			//despues guardar esto dentro de un arreglo de la clase Response
			resArray = append(resArray,res)

			respuestasArray = nil

			pregunta.ID = ID
			pregunta.NombreP = nombrep
			pregunta.DescripcionP = descripcionp
			pregunta.Respuesta = respuesta

			//Llenar arreglo de respuestas por esta pregunta
			respuestas.ID = IDS
			respuestas.Respuesta = answer
			respuestas.Idp = idp
			fmt.Println("hi")
			respuestasArray = append(respuestasArray, respuestas)

			//si el ide de una pregunta es igual a la ultima posicion
			//entonces debo guardar todo en el arreglo
			responsejoin.Pregunta = pregunta
			responsejoin.Respuestas = respuestasArray
			//despues guardar responsejoin en la clase empaquetadora res
			res.Array = responsejoin
			//despues guardar esto dentro de un arreglo de la clase Response
			resArray = append(resArray,res)
			
		}else if pregunta.ID != ID  && limit != count {
			fmt.Println("Pregunta diferente")
			//////////////////
			responsejoin.Pregunta = pregunta
			responsejoin.Respuestas = respuestasArray
			//despues guardar responsejoin en la clase empaquetadora res
			res.Array = responsejoin
			//despues guardar esto dentro de un arreglo de la clase Response
			resArray = append(resArray,res)

			respuestasArray = nil

			pregunta.ID = ID
			pregunta.NombreP = nombrep
			pregunta.DescripcionP = descripcionp
			pregunta.Respuesta = respuesta

			//Llenar arreglo de respuestas por esta pregunta
			respuestas.ID = IDS
			respuestas.Respuesta = answer
			respuestas.Idp = idp
			respuestasArray = append(respuestasArray, respuestas)
		}else if limit == count {
			fmt.Println("pregunta final")
			//si el ide de una pregunta es igual a la ultima posicion
			//entonces debo guardar todo en el arreglo
			responsejoin.Pregunta = pregunta
			responsejoin.Respuestas = respuestasArray
			//despues guardar responsejoin en la clase empaquetadora res
			res.Array = responsejoin
			//despues guardar esto dentro de un arreglo de la clase Response
			resArray = append(resArray,res)
		}

	}

	return resArray
}

func GetPreguntaxRespuesta1() ([]ResponseNormal) {
	db := database.GetConnection()

	rows, err := db.Query("SELECT * FROM pregunta p inner join respuesta r on p.id=r.idp")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	joinN := JoinNormal{}
	responseN := ResponseNormal{}
	var response []ResponseNormal
	pregunta := Pregunta{}
	//var r []Respuesta
	
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
			log.Println("cannot read current row")
			return nil
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

		joinN.Pregunta = pregunta
		joinN.Respuesta = respuestas
		responseN.Arrays = joinN
		response = append(response,responseN)
	}

	return response
}

func pointer(xPtr *string) {//pass address
	*xPtr = "hi"//new valye to the pointer
  } 

func RespuestaCorrecta(id int, res string) (string) {

	db := database.GetConnection()
	//var response string
	rows,err := db.Query("SELECT respuesta FROM pregunta where id=$1 and respuesta = $2",id,res)

	if err != nil{
		log.Fatal(err)
	}
	var response string
	var respuesta string
	for rows.Next() {

		err := rows.Scan(&respuesta)
		if err != nil{
			log.Fatal(err)
			
		}
		
		
		fmt.Println(respuesta)
		if respuesta != "" {
			fmt.Println("encontro pregunta")
			response = "Respuesta Correcta"
			//pointer(&response)
		}
		response = "Respuesta Incorrecta"
		

	}

	return response
	
}