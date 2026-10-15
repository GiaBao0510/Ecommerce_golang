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

type ShopOwnerController struct {
	shopOwnerService service.IShopOwnerService
	logger           *zap.Logger
}

func NewShopOwnerController(shopOwnerService service.IShopOwnerService, logger *zap.Logger) *ShopOwnerController {
	return &ShopOwnerController{shopOwnerService: shopOwnerService, logger: logger}
}

func (c *ShopOwnerController) GetAll(ctx *gin.Context) error {
	result, err := c.shopOwnerService.GetAll(ctx)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "All shop owners retrieved successfully", result)
	return nil
}

func (c *ShopOwnerController) GetByID(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	result, err := c.shopOwnerService.GetByID(ctx, param.ID)
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner retrieved successfully", result)
	return nil
}

func (c *ShopOwnerController) GetByName(ctx *gin.Context) error {
	result, err := c.shopOwnerService.GetByName(ctx, ctx.Param("name"))
	if err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner retrieved successfully", result)
	return nil
}

func (c *ShopOwnerController) Create(ctx *gin.Context) error {
	input := models.ShopOwner{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.shopOwnerService.Create(ctx, &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner created successfully", nil)
	return nil
}

func (c *ShopOwnerController) Update_Put(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	input := models.ShopOwner{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	if err := c.shopOwnerService.Update_Put(ctx, param.ID, &input); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner updated successfully", nil)
	return nil
}

func (c *ShopOwnerController) UpdateName_Patch(ctx *gin.Context) error {
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	if err := c.shopOwnerService.UpdateName_Patch(ctx, param.ID, input.Name); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner name updated successfully", nil)
	return nil
}

func (c *ShopOwnerController) UpdateDescription_Patch(ctx *gin.Context) error {
	var input struct {
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	if err := c.shopOwnerService.UpdateDescription_Patch(ctx, param.ID, input.Description); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner description updated successfully", nil)
	return nil
}

func (c *ShopOwnerController) UpdateAvatar_Patch(ctx *gin.Context) error {
	var input struct {
		AvatarURL string `json:"avatar_url"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		return apperrors.NewBadRequestError("Invalid input data: " + err.Error())
	}
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	if err := c.shopOwnerService.UpdateAvatar_Patch(ctx, param.ID, input.AvatarURL); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner avatar updated successfully", nil)
	return nil
}

func (c *ShopOwnerController) Delete(ctx *gin.Context) error {
	var param dto.ID_Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		return HandleValidationError(err)
	}
	if err := c.shopOwnerService.Delete(ctx, param.ID); err != nil {
		return err
	}
	response.Success_Response(ctx, http.StatusOK, "Shop owner deleted successfully", nil)
	return nil
}
