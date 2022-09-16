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

// Author godoc
// @ID create-author
// @Router /v1/author [POST]
// @Summary create author
// @Description create author by inserting name
// @Tags author
// @Accept json
// @Produce json
// @Param author body book_service.CreateAuthorRequest true "author"
// @Success 200 {object} models.ResponseModel{data=models.Author} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateAuthor(c *gin.Context) {
	var author book_service.CreateAuthorRequest

	if err := c.BindJSON(&author); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.AuthorService().Create(c.Request.Context(), &author)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating author", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllAuthor godoc
// @ID get-author
// @Router /v1/author [GET]
// @Summary get author all
// @Description get author
// @Tags author
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllAuthorResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllAuthor(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.AuthorService().GetAll(
		c.Request.Context(),
		&book_service.GetAllAuthorRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all authors", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-AuthorByID godoc
// @ID get-author-byID
// @Router /v1/author/{author_id} [GET]
// @Summary get author by ID
// @Description get author
// @Tags author
// @Accept json
// @Produce json
// @Param author_id path string true "author_id"
// @Success 200 {object} models.ResponseModel{data=models.Author} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAuthor(c *gin.Context) {
	var author book_service.Author
	author_id := c.Param("author_id")

	if !util.IsValidUUID(author_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "author_id  is not valid", errors.New("author_id is not valid"))
		return
	}

	resp, err := h.services.AuthorService().Get(
		context.Background(),
		&book_service.AuthorId{
			Id: author_id,
		},
	)

	err = ProtoToStruct(&author, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	if !handleError(h.log, c, err, "error while getting author") {
		return
	}

	err = ProtoToStruct(&author, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", &author)

}

// Update Author godoc
// @ID update_author
// @Router /v1/author/{author_id} [PUT]
// @Summary Update Author
// @Description Update Author by ID
// @Tags author
// @Accept json
// @Produce json
// @Param author_id path string true "author_id"
// @Param author body models.CreateAuthor true "author"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateAuthor(c *gin.Context) {
	var status models.Status
	var author models.CreateAuthor

	author_id := c.Param("author_id")

	if !util.IsValidUUID(author_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid author id", errors.New("author id is not valid"))
		return
	}

	err := c.BindJSON(&author)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	resp, err := h.services.AuthorService().Update(
		context.Background(),
		&book_service.Author{
			Id:        author_id,
			FirstName: author.FirstName,
			LastName:  author.LastName,
		},
	)

	if !handleError(h.log, c, err, "error while getting author") {
		return
	}

	err = ProtoToStruct(&status, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Updated")
}

// Delete Author godoc
// @ID delete-author
// @Router /v1/author/{author_id} [DELETE]
// @Summary delete author
// @Description Delete Author
// @Tags author
// @Accept json
// @Produce json
// @Param author_id path string true "author_id"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteAuthor(c *gin.Context) {

	var author models.Author
	author_id := c.Param("author_id")

	if !util.IsValidUUID(author_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "author id is not valid", errors.New("author id is not valid"))
		return
	}

	resp, err := h.services.AuthorService().Delete(
		context.Background(),
		&book_service.AuthorId{
			Id: author_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting author") {
		return
	}

	err = ProtoToStruct(&author, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Deleted")
}
