package http

import (
	"net/http"

	"github.com/GiaBao0510/Ecommerce_golang/internal/models"
	"github.com/GiaBao0510/Ecommerce_golang/internal/service"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/apperrors"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductDetailController struct {
	productDetailService service.IProductDetailService
	logger               *zap.Logger
}

func NewProductDetailController(productDetailService service.IProductDetailService, logger *zap.Logger) *ProductDetailController {
	return &ProductDetailController{productDetailService: productDetailService, logger: logger}
}

func (c *ProductDetailController) GetAll(ctx *gin.Context) error {
	result, err := c.productDetailService.GetAll(ctx)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "All product details retrieved successfully", result)
	return nil
}

func (c *ProductDetailController) GetByID(ctx *gin.Context) error {
	result, err := c.productDetailService.GetByID(ctx, ctx.Param("id"))
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Product detail retrieved successfully", result)
	return nil
}

func (c *ProductDetailController) Create(ctx *gin.Context) error {
	input := models.ProductDetail{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.productDetailService.Create(ctx, &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Product detail created successfully", nil)
	return nil
}

func (c *ProductDetailController) Update_Put(ctx *gin.Context) error {
	input := models.ProductDetail{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.productDetailService.Update_Put(ctx, ctx.Param("id"), &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Product detail updated successfully", nil)
	return nil
}

func (c *ProductDetailController) UpdatePrice_Patch(ctx *gin.Context) error {
	var input struct {
		Price string `json:"price" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.productDetailService.UpdatePrice_Patch(ctx, ctx.Param("id"), input.Price); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Product price updated successfully", nil)
	return nil
}

func (c *ProductDetailController) UpdateLikeCount_Patch(ctx *gin.Context) error {
	var input struct {
		LikeCount int32 `json:"like_count"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.productDetailService.UpdateLikeCount_Patch(ctx, ctx.Param("id"), input.LikeCount); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Product like count updated successfully", nil)
	return nil
}

func (c *ProductDetailController) UpdateQuantity_Patch(ctx *gin.Context) error {
	var input struct {
		Quantity int32 `json:"quantity"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.productDetailService.UpdateQuantity_Patch(ctx, ctx.Param("id"), input.Quantity); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Product quantity updated successfully", nil)
	return nil
}

func (c *ProductDetailController) Delete(ctx *gin.Context) error {
	if err := c.productDetailService.Delete(ctx, ctx.Param("id")); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Product detail deleted successfully", nil)
	return nil
}
