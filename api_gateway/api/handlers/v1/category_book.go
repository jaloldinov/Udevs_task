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

// Category godoc
// @ID create-categorys
// @Router /v1/category [POST]
// @Summary create category
// @Description create category
// @Tags category
// @Accept json
// @Produce json
// @Param category body book_service.CreateCategoryRequest true "category"
// @Success 200 {object} models.ResponseModel{data=book_service.Category} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var category book_service.CreateCategoryRequest

	if err := c.BindJSON(&category); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.CategoryService().Create(c.Request.Context(), &category)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating profession", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllCategory godoc
// @ID get-category
// @Router /v1/category [GET]
// @Summary get category all
// @Description get category
// @Tags category
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=book_service.GetAllCategoryResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllCategory(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.CategoryService().GetAll(
		c.Request.Context(),
		&book_service.GetAllCategoryRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all categorys", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-CategoryByID godoc
// @ID get-category-byID
// @Router /v1/category/{category_id} [GET]
// @Summary get category by ID
// @Description get category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.ResponseModel{data=models.Category} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetCategory(c *gin.Context) {
	var category models.Category
	category_id := c.Param("category_id")

	if !util.IsValidUUID(category_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "category_id  is not valid", errors.New("category_id is not valid"))
		return
	}

	resp, err := h.services.CategoryService().Get(
		context.Background(),
		&book_service.CategoryId{
			Id: category_id,
		},
	)

	err = ProtoToStruct(&category, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	if !handleError(h.log, c, err, "error while getting author") {
		return
	}

	err = ProtoToStruct(&category, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Update Category godoc
// @ID update_category
// @Router /v1/category/{category_id} [PUT]
// @Summary Update Category
// @Description Update Category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Param category body models.CreateCategory true "category"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var status models.Status
	var category models.CreateCategory

	category_id := c.Param("category_id")

	if !util.IsValidUUID(category_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "invalid category id", errors.New("category id is not valid"))
		return
	}

	err := c.BindJSON(&category)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}
	resp, err := h.services.CategoryService().Update(
		context.Background(),
		&book_service.Category{
			Id:           category_id,
			CategoryName: category.Name,
		},
	)

	if !handleError(h.log, c, err, "error while getting category") {
		return
	}

	err = ProtoToStruct(&status, resp)
	if !handleError(h.log, c, err, "error while parsing to struct") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Updated")
}

/// Delete Category godoc
// @ID delete-category
// @Router /v1/category/{category_id} [DELETE]
// @Summary delete category
// @Description Delete Category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteCategory(c *gin.Context) {

	var category models.Category
	category_id := c.Param("category_id")

	if !util.IsValidUUID(category_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "category id is not valid", errors.New("category id is not valid"))
		return
	}

	resp, err := h.services.CategoryService().Delete(
		context.Background(),
		&book_service.CategoryId{
			Id: category_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting category") {
		return
	}

	err = ProtoToStruct(&category, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", "Deleted")
}
