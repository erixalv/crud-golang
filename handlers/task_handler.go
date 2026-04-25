package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/erixalv/crud-golang/models"
	"github.com/gorilla/mux"
)

type TaskHandler struct {
	DB *sql.DB
}

//construtor de taskhandler
func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB : db}
}

func (taskHandler *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []models.Tasks{}

	rows, err := taskHandler.DB.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()


	//rows.Next -> método de rows que itera sobre as linhas
	for rows.Next() {
		var task models.Tasks

		//rows.Scan -> lê cada linha passando os atributos na variável (task)
		err := rows.Scan(&task.ID, &task.Title, &task.Desc, &task.Status)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (taskHandler *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Tasks

    err := json.NewDecoder(r.Body).Decode(&task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    _, err = taskHandler.DB.Exec(
        "INSERT INTO tasks (title, description) VALUES ($1, $2)",
        task.Title, task.Desc,
    )
    if err != nil {
        http.Error(w, "Error in inserting task", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (taskHandler *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var task models.Tasks
	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := taskHandler.DB.Exec(
		"UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4", 
		task.Title, task.Desc, task.Status, id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No task found with this ID", http.StatusNotFound)
		return
	}

	task.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

func (taskHandler *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	result, err := taskHandler.DB.Exec(
		"DELETE FROM tasks WHERE id = $1", id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "No task found with this ID", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}