package models

import (
	"github.com/kanwulf01/go-rest-api/src/database"
	"fmt"
	"log"
)

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
// funcion sql para insertar datos por POST
func Insert(description string) (Todo, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("INSERT INTO todos(description) VALUES($1) RETURNING id", description).Scan(&todo_id)

	if todo_id == 0 {
		return Todo{}, false
	}
	fmt.Println(todo_id)
	return Todo{todo_id, ""}, true
}

//funcion sql para llevar datos con GET
func Get(id string) (Todo, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM todos where id= $1",id)

	var ID int
	var description string
	err := row.Scan(&ID, &description)
	if err != nil {
		return Todo{},false
	}

	return Todo{ID,description}, true
}

func GetAll() ([]Todo) {
	db := database.GetConnection()
	sizerows, err := db.Query("select count(*) from (SELECT * FROM todos) as subconsul;")

	fmt.Println(sizerows)

	rows, err := db.Query("SELECT * FROM todos ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		t := Todo{}

		var ID int
		var description string

		err := rows.Scan(&ID, &description)
		if err != nil {
			log.Fatal(err)
		}

		t.ID = ID
		t.Description = description

		todos = append(todos, t)

	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return todos
}