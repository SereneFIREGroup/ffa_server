package family

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	dbPkg "github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"gorm.io/gorm"
)

const (
	// NameLength is the max length of family name
	NameLength = 128

	// StatusNormal is the normal status of family
	StatusNormal = 1
	// StatusDeleted is the deleted status of family
	StatusDeleted = 2
)

type Family struct {
	ULID       string `json:"ulid" gorm:"primarykey,column:ulid"`
	Owner      string `json:"owner" gorm:"column:owner"`
	Name       string `json:"name" gorm:"column:name"`
	Status     int    `json:"status" gorm:"column:status"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	FIREGold   int64  `json:"fire_gold" gorm:"column:fire_gold"`
}

func NewBaseFamily(ulid, owner, name string) *Family {
	return &Family{
		ULID:       ulid,
		Owner:      owner,
		Name:       name,
		Status:     StatusNormal,
		CreateTime: time.Now().Unix(),
		FIREGold:   0,
	}
}

func (f *Family) SetFIREGold(amount int64) {
	f.FIREGold = amount
}

// InsertFamily insert family to DB
func InsertFamily(ctx context.Context, db *gorm.DB, family *Family) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "InsertFamily")
	defer span.Finish()

	err := db.Table(dbPkg.TableFamily).Create(&family).Error
	return errors.Sql(err)
}

// UpdateFamily update family to DB
func UpdateFamily(ctx context.Context, db *gorm.DB, family *Family) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "UpdateFamily")
	defer span.Finish()

	err := db.Table(dbPkg.TableFamily).Save(&family).Error
	return errors.Sql(err)
}

// GetFamily get family from DB
func GetFamily(ctx context.Context, db *gorm.DB, ulid string) (*Family, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetFamily")
	defer span.Finish()

	f := new(Family)
	err := db.Table(dbPkg.TableFamily).Where("ulid = ?", ulid).First(&f).Error
	if err != nil {
		return nil, errors.Sql(err)
	}
	return f, nil
}
