package controllers

import (
	"net/http"

	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookController struct {
	service services.BookService
}

func NewBookController(s services.BookService) *BookController {
	return &BookController{s}
}

func (c *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	book, err := c.service.GetBookById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)

}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var bookRequest requestdto.BookRequest
	if err := ctx.ShouldBindJSON(&bookRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookResponse, err := c.service.CreateBook(bookRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, bookResponse)
}
