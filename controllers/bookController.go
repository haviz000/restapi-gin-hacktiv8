package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/restapi-gin-hacktiv8/database"
	"github.com/haviz000/restapi-gin-hacktiv8/models"
)

func GetBooks(ctx *gin.Context) {
	var books []models.Book

	sqlQuery := `SELECT * FROM books`

	db := database.ConnectDB()

	result, err := db.Query(sqlQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for result.Next() {
		var book models.Book
		err := result.Scan(
			&book.BookID,
			&book.Title,
			&book.Author,
			&book.Desc)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, book)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Books": books,
	})
}

func GetBookById(ctx *gin.Context) {
	var book models.Book
	var idBook = ctx.Param("idBook")
	var isNotFound = true

	id, err := strconv.Atoi(idBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":   "Id buku harus angka",
			"error_messages": err,
		})
		return
	}
	sqlQuery := `SELECT * FROM books WHERE id=$1`

	db := database.ConnectDB()

	result := db.QueryRow(sqlQuery, id)
	err = result.Scan(&book.BookID, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if book.BookID != 0 {
		isNotFound = false
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "data not found",
			"error_messages": fmt.Sprintf("buku dengan id %d tidak ditemukan", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Book": book,
	})
}

func CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlQuery := `
	INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	RETURNING *
	`

	db := database.ConnectDB()

	var bookResult models.Book

	result := db.QueryRow(sqlQuery, book.Title, book.Author, book.Desc)
	err := result.Scan(&bookResult.BookID, &bookResult.Title, &bookResult.Author, &bookResult.Desc)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.String(http.StatusOK, "created")
}

func UpdateBook(ctx *gin.Context) {
	var book models.Book
	var idBook = ctx.Param("idBook")
	var isNotFound = true

	id, err := strconv.Atoi(idBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":   "Id buku harus angka",
			"error_messages": err,
		})
		return
	}
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlQuery := `
	UPDATE books
	SET title=$1, author=$2, description=$3
	WHERE id=$4
	RETURNING *
	`

	db := database.ConnectDB()

	var bookResult models.Book
	err = db.QueryRow(sqlQuery, book.Title, book.Author, book.Desc, id).Scan(
		&bookResult.BookID,
		&bookResult.Title,
		&bookResult.Author,
		&bookResult.Desc,
	)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if bookResult.BookID != 0 {
		isNotFound = false
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data not found",
			"error_message": fmt.Sprintf("buku dengan id %d tidak ditemukan", id),
		})
		return
	}

	ctx.String(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	var idBook = ctx.Param("idBook")
	var isNotFound = true

	id, err := strconv.Atoi(idBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":   "Id buku harus angka",
			"error_messages": err,
		})
		return
	}

	sqlQuery := `
	DELETE FROM books
	WHERE id=$1
	RETURNING id
	`

	db := database.ConnectDB()

	var bookId int
	err = db.QueryRow(sqlQuery, id).Scan(&bookId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if bookId != 0 {
		isNotFound = false
	}

	if isNotFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "data not found",
			"error_message": fmt.Sprintf("buku dengan id %d tidak ditemukan", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, "deleted")
}
