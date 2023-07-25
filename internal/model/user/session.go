package user

import (
	"context"

	"github.com/opentracing/opentracing-go"
	dbPkg "github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"github.com/serenefiregroup/ffa_server/pkg/random"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	tokenLen = 64
)

type Session struct {
	UserULID string `json:"user_ulid" gorm:"primarykey,column:user_ulid"`
	Token    string `json:"token" gorm:"column:token"`
}

func NewBaseSession(userULID string) *Session {
	return &Session{
		UserULID: userULID,
		Token:    random.Alphanumeric(tokenLen),
	}
}

// InsertSession insert session to DB
func InsertSession(ctx context.Context, db *gorm.DB, s *Session) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "InsertSession")
	defer span.Finish()

	err := db.Table(dbPkg.TableSession).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_ulid"}},
		DoUpdates: clause.AssignmentColumns([]string{"token"}),
	}).Create(&s).Error
	if err != nil {
		return errors.Sql(err)
	}
	return errors.Sql(err)
}

// DeleteSession get session from DB
func DeleteSession(ctx context.Context, db *gorm.DB, userULID string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "DeleteSession")
	defer span.Finish()

	err := db.Table(dbPkg.TableSession).Delete(&Session{}, "user_ulid = ?", userULID).Error
	return errors.Sql(err)
}

// GetSession get session from DB
func GetSession(ctx context.Context, db *gorm.DB, userULID, token string) (*Session, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetSession")
	defer span.Finish()

	s := new(Session)
	err := db.Table(dbPkg.TableSession).Where("user_ulid = ? AND token = ?", userULID, token).First(&s).Error
	if err != nil {
		return nil, errors.Sql(err)
	}
	return s, nil
}
