package initialize

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vuongkhanh1537/tms-backend/global"
	"go.uber.org/zap"
)

func InitPostgres() {
	cfg := global.Config.Database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBname,
		cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		global.Logger.Error("Init Postgres Error", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("Init Postgres Successful")
	global.DB = db
}