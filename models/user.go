package models

type Tasks struct {
	ID 			int 		`json:"id"`
	Title 		string		`json:"title"`
	Desc		string		`json:"desc"`
	Status		bool		`json:"status"`
}