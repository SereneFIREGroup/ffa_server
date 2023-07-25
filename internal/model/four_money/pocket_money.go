package four_money

import (
	"context"

	"github.com/oklog/ulid/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/serenefiregroup/ffa_server/pkg/date"
	dbPkg "github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"gorm.io/gorm"
)

// PocketMoney Pocket Money 零花钱，活钱管理
type PocketMoney struct {
	Ulid       string `gorm:"type:varchar(26);primary_key;not null,column:ulid"`
	FamilyUlid string `gorm:"type:varchar(26);not null,column:family_ulid"`
	UserUlid   string `gorm:"type:varchar(26);not null,column:user_ulid"`
	Amount     int64  `gorm:"type:bigint;not null,column:amount"`
	Category   string `gorm:"type:varchar(32);not null,column:category"`
	Remark     string `gorm:"type:varchar(256);not null,column:remark"`
	CreateTime int64  `gorm:"type:bigint;not null,column:create_time"`
}

// NewBasePocketMoney create a new PocketMoney
func NewBasePocketMoney(familyUlid, userUlid string, amount int64, category, remark string) *PocketMoney {
	return &PocketMoney{
		Ulid:       ulid.Make().String(),
		FamilyUlid: familyUlid,
		UserUlid:   userUlid,
		Amount:     amount,
		Category:   category,
		Remark:     remark,
		CreateTime: date.TodayStart(), // set create time to today start
	}
}

// InsertPocketMoney insert PocketMoney to DB
func InsertPocketMoney(ctx context.Context, db *gorm.DB, pm *PocketMoney) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "InsertPocketMoney")
	defer span.Finish()

	err := db.Table(dbPkg.TablePocketMoney).Create(&pm).Error
	return errors.Sql(err)
}

// UpdatePocketMoney update PocketMoney to DB
func UpdatePocketMoney(ctx context.Context, db *gorm.DB, pm *PocketMoney) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "UpdatePocketMoney")
	defer span.Finish()

	err := db.Table(dbPkg.TablePocketMoney).Save(&pm).Error
	return errors.Sql(err)
}

// ListPocketMoney list PocketMoney from DB
func ListPocketMoney(ctx context.Context, db *gorm.DB, familyUlid, userUlid string, startTime, endTime int64) ([]*PocketMoney, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListPocketMoney")
	defer span.Finish()

	var pms []*PocketMoney
	err := db.Table(dbPkg.TablePocketMoney).
		Where("family_ulid = ? and user_ulid=? and create_time>=? and create_time<?", familyUlid, userUlid, startTime, endTime).
		Find(&pms).Error
	return pms, errors.Sql(err)
}
