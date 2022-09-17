package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaloldinov/Udevs_task/api_gateway/api/models"
	"github.com/jaloldinov/Udevs_task/api_gateway/genproto/book_service"
	"github.com/jaloldinov/Udevs_task/api_gateway/pkg/util"
)

// Book godoc
// @ID create-books
// @Router /v1/book [POST]
// @Summary create book
// @Description create book
// @Tags book
// @Accept json
// @Produce json
// @Param book body book_service.CreateBookRequest true "book"
// @Success 200 {object} models.ResponseModel{data=book_service.Book} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateBook(c *gin.Context) {
	var book book_service.CreateBookRequest

	if err := c.BindJSON(&book); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.BookService().Create(c.Request.Context(), &book)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating profession", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllBook godoc
// @ID get-book
// @Router /v1/book [GET]
// @Summary get book all
// @Description get book
// @Tags book
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=book_service.GetAllBookResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllBook(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.BookService().GetAll(
		c.Request.Context(),
		&book_service.GetAllBookRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all books", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-BookByID godoc
// @ID get-book-byID
// @Router /v1/book/{book_id} [GET]
// @Summary get book by ID
// @Description get book
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path string true "book_id"
// @Success 200 {object} models.ResponseModel{data=models.BookModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBook(c *gin.Context) {
	var book models.BookModel
	book_id := c.Param("book_id")

	if !util.IsValidUUID(book_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "book_id  is not valid", errors.New("book_id is not valid"))
		return
	}

	resp, err := h.services.BookService().Get(
		context.Background(),
		&book_service.BookId{
			Id: book_id,
		},
	)

	err = ProtoToStruct(&book, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	if !handleError(h.log, c, err, "error while getting author") {
		return
	}

	err = ProtoToStruct(&book, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Update Book godoc
// @ID update_book
// @Router /v1/book/{book_id} [PUT]
// @Summary Update Book
// @Description Update Book by ID
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path string true "book_id"
// @Param book body models.CreateBookModel true "book"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateBook(c *gin.Context) {
	var status models.Status
	var book models.CreateBookModel

	book_id := c.Param("book_id")

	if !util.IsValidUUID(book_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid book id", errors.New("book id is not valid"))
		return
	}

	err := c.BindJSON(&book)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	resp, err := h.services.BookService().Update(
		context.Background(),
		&book_service.Book{
			Id:         book_id,
			Name:       book.Name,
			CategoryId: book.CategoryId,
			AuthorId:   book.AuthorId,
		},
	)

	if !handleError(h.log, c, err, "error while getting book") {
		return
	}

	err = ProtoToStruct(&status, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Updated")
}

/// Delete Book godoc
// @ID delete-book
// @Router /v1/book/{book_id} [DELETE]
// @Summary delete book
// @Description Delete Book
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path string true "book_id"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteBook(c *gin.Context) {

	var book models.BookModel
	book_id := c.Param("book_id")

	if !util.IsValidUUID(book_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "book id is not valid", errors.New("book id is not valid"))
		return
	}

	resp, err := h.services.BookService().Delete(
		context.Background(),
		&book_service.BookId{
			Id: book_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting book") {
		return
	}

	err = ProtoToStruct(&book, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Deleted")
}
