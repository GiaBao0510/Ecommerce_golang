package global

import (
	"database/sql"

	"github.com/GiaBao0510/Ecommerce_golang/internal/database"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/logger"
	"github.com/GiaBao0510/Ecommerce_golang/pkg/setting"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.AppLoggers
	PostgreSQL *sql.DB
	Redis *redis.Client
	DB *database.Queries
)