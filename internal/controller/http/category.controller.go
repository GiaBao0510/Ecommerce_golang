package http

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/dto"
	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryController struct {
	categoryService service.ICategoryService
	logger          *zap.Logger
}

func NewCategoryController(categoryService service.ICategoryService, logger *zap.Logger) *CategoryController {
	return &CategoryController{categoryService: categoryService, logger: logger}
}

func (c *CategoryController) GetAll(ctx *gin.Context) error {
	result, err := c.categoryService.GetAll(ctx)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "All categories retrieved successfully", result)
	return nil
}

func (c *CategoryController) GetByID(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	result, err := c.categoryService.GetByID(ctx, param.ID)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Category retrieved successfully", result)
	return nil
}

func (c *CategoryController) GetByName(ctx *gin.Context) error {
	result, err := c.categoryService.GetByName(ctx, ctx.Param("name"))
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Category retrieved successfully", result)
	return nil
}

func (c *CategoryController) Create(ctx *gin.Context) error {
	input := models.Category{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	result, err := c.categoryService.Create(ctx, &input)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Category created successfully", gin.H{"id": result})
	return nil
}

func (c *CategoryController) Update_Put(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	input := models.Category{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.categoryService.Update_Put(ctx, param.ID, &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Category updated successfully", nil)
	return nil
}

func (c *CategoryController) Delete(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	if err := c.categoryService.Delete(ctx, param.ID); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Category deleted successfully", nil)
	return nil
}
