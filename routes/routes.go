package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mathaono/go-api-rest/controllers"
)

// Mesma coisa que o package HTTP padrão do Golang - Mas aqui está sendo utilizada a função HandleFunc() da biblioteca externa Gorilla Mux
func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnPersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
