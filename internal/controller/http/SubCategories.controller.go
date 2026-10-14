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

type SubCategoryController struct {
	subCategoryService service.ISubCategoryService
	logger             *zap.Logger
}

func NewSubCategoryController(subCategoryService service.ISubCategoryService, logger *zap.Logger) *SubCategoryController {
	return &SubCategoryController{subCategoryService: subCategoryService, logger: logger}
}

func (c *SubCategoryController) GetAll(ctx *gin.Context) error {
	result, err := c.subCategoryService.GetAll(ctx)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "All subcategories retrieved successfully", result)
	return nil
}

func (c *SubCategoryController) GetByID(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	result, err := c.subCategoryService.GetByID(ctx, param.ID)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Subcategory retrieved successfully", result)
	return nil
}

func (c *SubCategoryController) GetByName(ctx *gin.Context) error {
	result, err := c.subCategoryService.GetByName(ctx, ctx.Param("name"))
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Subcategory retrieved successfully", result)
	return nil
}

func (c *SubCategoryController) Create(ctx *gin.Context) error {
	input := models.SubCategory{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	result, err := c.subCategoryService.Create(ctx, &input)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Subcategory created successfully", gin.H{"id": result})
	return nil
}

func (c *SubCategoryController) Update_Put(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	input := models.SubCategory{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.subCategoryService.Update_Put(ctx, param.ID, &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Subcategory updated successfully", nil)
	return nil
}

func (c *SubCategoryController) Delete(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	if err := c.subCategoryService.Delete(ctx, param.ID); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Subcategory deleted successfully", nil)
	return nil
}
