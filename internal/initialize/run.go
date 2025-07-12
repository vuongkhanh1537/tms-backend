package initialize

import (
	"fmt"

	"github.com/vuongkhanh1537/tms-backend/global"
)

func Run() {
	LoadConfig()

	r := InitRouter()

	addr := fmt.Sprintf(":%d", global.Config.Server.Port)
	r.Run(addr)
}