package main

import (
	"database/sql"
	"encoding/json"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails"
)

type Persona struct {
	Id   int
	Name string
	Sex  string
}

func create(name string, sex string) string {
	database, err := sql.Open("sqlite3", "./data.db")

	if err != nil {
		panic(err.Error())
	}

	query, err := database.Prepare("CREATE TABLE IF NOT EXISTS personas (id INTEGER PRIMARY KEY AutoIncrement, nombre TEXT, sexo TEXT)")

	if err != nil {
		panic(err.Error())
	}

	query.Exec()

	query, err = database.Prepare("INSERT INTO personas (nombre, sexo) VALUES (?,?) ")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(name, sex)

	return "Creado con exito"
}

func getPersona() string {
	database, _ := sql.Open("sqlite3", "./data.db")

	rows, err := database.Query("SELECT * FROM personas")

	if err != nil {
		panic(err.Error())
	}

	persona := Persona{}
	DataPersona := []Persona{}

	for rows.Next() {
		var id int
		var nombre, sexo string

		err := rows.Scan(&id, &nombre, &sexo)

		if err != nil {
			panic(err.Error())
		}

		persona.Id = id
		persona.Name = nombre
		persona.Sex = sexo

		DataPersona = append(DataPersona, persona)

	}

	jsonPersona, err := json.Marshal(DataPersona)

	if err != nil {
		panic(err.Error())
	}

	return string(jsonPersona)

}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "mary_system",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(getPersona)
	app.Bind(create)
	app.Run()
}
