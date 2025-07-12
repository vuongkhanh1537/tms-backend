package initialize

import (
	"github.com/vuongkhanh1537/tms-backend/global"
	"github.com/vuongkhanh1537/tms-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}