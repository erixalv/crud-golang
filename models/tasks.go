package models

type Tasks struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Status bool   `json:"status"`
}

func NewTask(title, desc string) *Tasks {
	return &Tasks{Title: title, Desc: desc}
}

const (
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description VARCHAR(250) NOT NULL,
		status BOOLEAN NOT NULL DEFAULT FALSE
	);`
)