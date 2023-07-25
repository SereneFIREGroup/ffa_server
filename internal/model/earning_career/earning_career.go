package earning_career

import (
	"context"

	"github.com/opentracing/opentracing-go"
	dbPkg "github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"gorm.io/gorm"
)

type Earning struct {
	gorm.Model
	UserULID string `json:"user_ulid" gorm:"column:user_ulid,index:idx_user"`
	Amount   int64  `json:"amount" gorm:"column:amount"`
	Category string `json:"category" gorm:"column:category,size:64"`
	Desc     string `json:"desc" gorm:"column:desc,size:256"`
}

func NewBaseEarning(userULID string, amount int64, category, desc string) *Earning {
	return &Earning{
		UserULID: userULID,
		Amount:   amount,
		Category: category,
		Desc:     desc,
	}
}

// InsertEarningCareer insert earning_career to DB
func InsertEarningCareer(ctx context.Context, db *gorm.DB, career *Earning) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "InsertEarningCareer")
	defer span.Finish()

	err := db.Table(dbPkg.TableEarningCareer).Create(&career).Error
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

// ListEarningCareer get earning_career from DB
func ListEarningCareer(ctx context.Context, db *gorm.DB, userULID string, startDate, endDate int64) ([]*Earning, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListEarningCareer")
	defer span.Finish()

	careers := make([]*Earning, 0)
	err := db.Table(dbPkg.TableEarningCareer).
		Where("user_ulid = ? and created_at BETWEEN ? AND ?", userULID, startDate, endDate).
		Order("created_at asc").Find(&careers).Error
	if err != nil {
		return nil, errors.Trace(err)
	}
	return careers, nil
}
