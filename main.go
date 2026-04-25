package main

import (
	"log"
	"net/http"

	"github.com/erixalv/crud-golang/config"
	"github.com/erixalv/crud-golang/handlers"
	"github.com/erixalv/crud-golang/models"
	"github.com/gorilla/mux"
)

//ponto de entrada da aplicação
func main() {
	dbConnection := config.SetupDatabase()
	_, err := dbConnection.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatalf("Erro ao criar tabela: %s", err)
	}

	//defer -> executa a função Close() quando a função em torno dela (a própria main) finalizar
	defer dbConnection.Close()

	taskHandler := handlers.NewTaskHandler(dbConnection)

	router := mux.NewRouter()
	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}