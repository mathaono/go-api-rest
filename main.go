package main

import (
	"fmt"

	"github.com/mathaono/go-api-rest/database"
	"github.com/mathaono/go-api-rest/routes"
)

func main() {
	/*models.Personalities = []models.Personality{
		{ID: 1, Name: "Nome 1", History: "Historia 1"},
		{ID: 2, Name: "Nome 2", History: "Historia 2"},
		{ID: 3, Name: "Nome 3", History: "Historia 3"},
		{ID: 4, Name: "Nome 4", History: "Historia 4"},
		{ID: 5, Name: "Nome 5", History: "Historia 5"},
	}*/

	fmt.Println("Iniciando Servidor")
	routes.HandleRequest()

	database.ConnectDataBase()

}
