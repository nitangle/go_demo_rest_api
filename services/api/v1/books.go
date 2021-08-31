package v1

import (
	"database/sql"
	"fmt"

	models "golang_rest_api/models"

)

type Book = models.Book  

type BookService struct{
	book Book
}


var db = models.GetDB()

func (bs *BookService) GetBook() Book{
	return bs.book
}

func Books() ([]Book,error){
	var books []Book
	rows,err := db.Query("SELECT * from books")
	if err != nil{
		return nil,fmt.Errorf("books get %v", err)
	}
	// fmt.Println("something found",rows)
	defer rows.Close()
	for rows.Next(){
		// fmt.Print("ok inside\n")
		var book Book
		if err := rows.Scan(&book.ID,&book.Title,&book.Author,&book.Category,&book.LentTo,&book.LentOn);err != nil{
			fmt.Println("ok sad",err)
			return nil,fmt.Errorf("books get %v",err)
		}
		// fmt.Print("book is",book)
		books = append(books,book)
	}
	if err := rows.Err(); err != nil{
		return nil,fmt.Errorf("books get %v",err)
	}
	return books, nil
}

func(bs *BookService) AddBook(b Book) (int64,error){
	result, err := db.Exec("INSERT INTO books (title, author, category, lentTo, lentOn) VALUES (?, ?, ?, ?, ?)", b.Title, b.Author, b.Category, b.LentTo, b.LentOn)
    if err != nil {
        return 0, fmt.Errorf("addBook: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addBook: %v", err)
    }
	bs.book = b
	bs.book.ID = id
    return id, nil
}

func(bs *BookService) BookByID (id string) (error){
	var book = bs.book
	row := db.QueryRow("SELECT * from books WHERE id=?",id)
	err := row.Scan(&book.ID,&book.Title,&book.Author,&book.Category,&book.LentTo,&book.LentOn)
	switch{
	case err==sql.ErrNoRows:
		return fmt.Errorf("bookByid: %v no such book found", id)

	case err!=nil:
		return fmt.Errorf("bookByid: %v", err)
	default:
		bs.book = book
		return nil

	}

}