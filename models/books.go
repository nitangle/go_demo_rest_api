package models
import (
	"database/sql"
)

type Book struct {
    ID     int64  `json:"id"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Category  string `json:"category"`
    LentTo  sql.NullInt64 `json:"lentTo"`
	LentOn sql.NullInt64 `json:"lentOn"`	

}

