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

type BrandController struct {
	brandService service.IBrandService
	logger       *zap.Logger
}

func NewBrandController(brandService service.IBrandService, logger *zap.Logger) *BrandController {
	return &BrandController{brandService: brandService, logger: logger}
}

func (c *BrandController) GetAll(ctx *gin.Context) error {
	result, err := c.brandService.GetAll(ctx)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "All brands retrieved successfully", result)
	return nil
}

func (c *BrandController) GetByID(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	result, err := c.brandService.GetByID(ctx, param.ID)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Brand retrieved successfully", result)
	return nil
}

func (c *BrandController) GetByName(ctx *gin.Context) error {
	result, err := c.brandService.GetByName(ctx, ctx.Param("name"))
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Brand retrieved successfully", result)
	return nil
}

func (c *BrandController) Create(ctx *gin.Context) error {
	input := models.Brand{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.brandService.Create(ctx, &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Brand created successfully", nil)
	return nil
}

func (c *BrandController) Update_Put(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	input := models.Brand{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.brandService.Update_Put(ctx, param.ID, &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Brand updated successfully", nil)
	return nil
}

func (c *BrandController) Delete(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	if err := c.brandService.Delete(ctx, param.ID); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Brand deleted successfully", nil)
	return nil
}
