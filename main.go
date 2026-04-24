package main

import (
	"log"
	"net/http"

	"github.com/erixalv/crud-golang/config"
)

//ponto de entrada da aplicação
func main() {
	dbConnection := config.SetupDatabase()
	//defer -> executa a função Close() quando a função acima dela (a própria main) finalizar
	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":8000", nil))
}