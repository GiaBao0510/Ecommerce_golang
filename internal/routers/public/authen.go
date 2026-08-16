package public

import (
	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthenRouter struct{}

func (r *AuthenRouter) InitAuthenRouter(Router *gin.RouterGroup, db *database.Queries, logger *zap.Logger) {
	authenController, err := Init
}