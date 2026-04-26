package models

type Tasks struct {
	ID			uint 		`gorm:"primarykey" json:"id"`
	Title  		string 		`gorm:"not null" json:"title"`
	Desc   		string 		`gorm:"not null;column:description" json:"desc"`
	Status 		bool   		`gorm:"default:false" json:"status"`
	UserID		uint 		`json:"user_id"`
}
