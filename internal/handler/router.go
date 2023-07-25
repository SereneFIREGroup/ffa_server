package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/log"
)

func Run() {
	port := config.String("port", "10001")
	portStr := fmt.Sprintf(":%s", port)

	debug := config.Bool("debug", true)
	gin.SetMode(gin.ReleaseMode)
	if debug {
		gin.SetMode(gin.DebugMode)
	}

	api := gin.New()
	api.Use(log.RecoveryHandleFunc())
	api.Use(log.LoggerHandleFunc())
	api.Use(PanicHandler)
	api.Use(Jaeger)

	log.Info("Server is running on port %s", port)

	auth := api.Group("/auth")
	auth.POST("/verify_phone", LeakBucket(), VerifyPhone)
	auth.POST("/sign_up", LeakBucket(), SignUp)
	auth.POST("/sign_in", SignIn)
	auth.GET("/me", Me)
	auth.POST("/sign_out", SignOut)

	family := api.Group("/family/:familyUUID")
	family.Use(CheckLogin())
	family.Use(CheckFamily())
	family.POST("/set_fire_gold", SetFIREGold)

	earning := api.Group("/earning")
	earning.Use(CheckLogin())
	earning.GET("/categories", ListEarningCategory)
	earning.POST("/add", AddEarning)
	earning.GET("/list", ListEarning)

	fourMoney := api.Group("/family/:familyUUID/four_money")
	fourMoney.Use(CheckLogin())
	family.Use(CheckFamily())
	fourMoney.GET("/categories", ListFourMoneyCategory)
	fourMoney.POST("/add", AddFourMoney)

	_ = api.Run(portStr)
}
