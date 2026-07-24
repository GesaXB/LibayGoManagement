package controllers

import (
	"net/http"
	"strconv"

	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	book, err := c.service.GetBookById(uint(id))
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
