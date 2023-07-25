package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/serenefiregroup/ffa_server/internal/handler"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/log"
	"github.com/serenefiregroup/ffa_server/pkg/profile"
	"github.com/serenefiregroup/ffa_server/pkg/sms"
)

func onStart(fn func() error) {
	if err := fn(); err != nil {
		panic(fmt.Sprintf("Error at onStart: %s\n", err))
	}
}

func init() {
	onStart(config.LoadConfig)
	onStart(log.LoadLog)
	onStart(db.InitDB)
	onStart(profile.Profile)
	onStart(sms.LoadConfig)
}

// Family Finance AI Server
func main() {
	gin.DefaultWriter = log.NewLoggerWriter()
	log.Info("Server start")
	handler.Run()
}
