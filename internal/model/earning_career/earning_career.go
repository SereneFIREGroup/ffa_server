package earning_career

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	dbPkg "github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"gorm.io/gorm"
)

const (
	UnLimitTime = -1
)

type Earning struct {
	ID         string `json:"id" gorm:"primarykey,column:id"`
	UserID     string `json:"user_id" gorm:"column:user_id"`
	Amount     int64  `json:"amount" gorm:"column:amount"`
	Category   string `json:"category" gorm:"column:category"`
	Desc       string `json:"desc" gorm:"column:desc"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
}

func (e *Earning) GetYear() int {
	y, _, _ := time.Unix(e.CreateTime, 0).Date()
	return y
}

func NewBaseEarning(userID string, amount int64, category, desc string) *Earning {
	return &Earning{
		UserID:     userID,
		Amount:     amount,
		Category:   category,
		Desc:       desc,
		CreateTime: time.Now().Unix(),
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
func ListEarningCareer(ctx context.Context, db *gorm.DB, userID string, startDate, endDate int64) ([]*Earning, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListEarningCareer")
	defer span.Finish()

	whereQuery := "user_id = ? "
	whereArgs := []interface{}{userID}
	if startDate != UnLimitTime && endDate != UnLimitTime {
		whereQuery += "and created_at BETWEEN ? AND ?"
		whereArgs = append(whereArgs, startDate, endDate)
	}

	careers := make([]*Earning, 0)
	err := db.Table(dbPkg.TableEarningCareer).
		Where(whereQuery, whereArgs).Order("created_at asc").Find(&careers).Error
	if err != nil {
		return nil, errors.Trace(err)
	}
	return careers, nil
}
