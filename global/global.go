package global

import (
	"database/sql"

	"github.com/vuongkhanh1537/tms-backend/pkg/logger"
	"github.com/vuongkhanh1537/tms-backend/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	DB *sql.DB
)
