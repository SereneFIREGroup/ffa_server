package model

import (
	earningCareerModel "github.com/serenefiregroup/ffa_server/internal/model/earning_career"
	familyModel "github.com/serenefiregroup/ffa_server/internal/model/family"
	"github.com/serenefiregroup/ffa_server/internal/model/four_money"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	verifyModel "github.com/serenefiregroup/ffa_server/internal/model/verify"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/log"
)

func init() {
	debug := config.Bool("debug", true)
	if !debug {
		return
	}
	err := db.DB.AutoMigrate(familyModel.Family{})
	if err != nil {
		log.Error("AutoMigrate family error: %s", err)
	}
	err = db.DB.AutoMigrate(userModel.User{})
	if err != nil {
		log.Error("AutoMigrate user error: %s", err)
	}
	err = db.DB.AutoMigrate(userModel.Session{})
	if err != nil {
		log.Error("AutoMigrate session error: %s", err)
	}
	err = db.DB.AutoMigrate(verifyModel.VerifyCode{})
	if err != nil {
		log.Error("AutoMigrate verify code error: %s", err)
	}
	err = db.DB.AutoMigrate(earningCareerModel.Earning{})
	if err != nil {
		log.Error("AutoMigrate earning error: %s", err)
	}
	err = db.DB.AutoMigrate(four_money.PocketMoney{})
	if err != nil {
		log.Error("AutoMigrate pocket money error: %s", err)
	}
}
