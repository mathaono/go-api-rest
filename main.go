package main

import (
	"fmt"

	"github.com/mathaono/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando Servidor")
	routes.HandleRequest()
}
