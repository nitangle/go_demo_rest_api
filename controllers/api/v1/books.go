package v1

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	v1s "golang_rest_api/services/api/v1"
)

type BookService = v1s.BookService

var bs BookService

// getBooks responds with the list of all books as JSON.
func GetBooks(c *gin.Context) {
	books,err := v1s.Books()
	if err!=nil{
		fmt.Printf("getBooks controller %v\n",err)
	}
    c.IndentedJSON(http.StatusOK, books)
}

//GetBookById gets book info by ID
func GetBookByID(c *gin.Context){
	id := c.Param("id")
	err := bs.BookByID(id)
	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"No such book exists!"})
		return 
	}
	c.IndentedJSON(http.StatusOK,bs.GetBook())
}

//postBooks adds a book
func PostBook(c *gin.Context){
	fmt.Printf("inside postbook %+v\n",c.Request.Body)
	book := bs.GetBook()
	if err:=c.BindJSON(&book);err!=nil{		
		fmt.Println("error while reading data into book obj",err)
		c.IndentedJSON(http.StatusBadRequest,err)
		return
	}
	id,err := bs.AddBook(book)
	if err!=nil{
		fmt.Println("id is",id)
		c.IndentedJSON(http.StatusInternalServerError,err)
		return
	}
	c.IndentedJSON(http.StatusCreated,bs.GetBook())
}

