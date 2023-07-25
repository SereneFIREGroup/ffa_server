package user

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/serenefiregroup/ffa_server/internal/model/constants"
	dbPkg "github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"github.com/serenefiregroup/ffa_server/pkg/hash"
	"gorm.io/gorm"
)

const (
	// NameLength is the max length of username
	NameLength = 128

	// StatusNormal normal user status
	StatusNormal = 1
	// StatusDisable disabled user status
	StatusDisable = 2
)

type User struct {
	FamilyULID string `json:"family_ulid" gorm:"column:family_ulid"`
	ULID       string `json:"ulid" gorm:"primarykey,column:ulid"`
	Name       string `json:"name" gorm:"column:name"`
	Avatar     string `json:"avatar" gorm:"column:avatar"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Password   string `json:"password" gorm:"column:password"`
	Status     uint8  `json:"status" gorm:"column:status"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
}

func NewBaseUser(familyULID, ulid, name, phone, password string) *User {
	return &User{
		FamilyULID: familyULID,
		ULID:       ulid,
		Name:       name,
		Avatar:     "",
		Phone:      phone,
		Password:   password,
		Status:     StatusNormal,
		CreateTime: time.Now().Unix(),
	}
}

func (u *User) SetAvatar(avatar string) {
	u.Avatar = avatar
}

func (u *User) Disable() {
	u.Status = StatusDisable
}

// BuildPasswordMD5 build password md5
func BuildPasswordMD5(password string) string {
	input := password + constants.ProductName
	return hash.MD5([]byte(input))
}

// ExistPhone check phone exist
func ExistPhone(db *gorm.DB, code string) (bool, error) {
	var count int64
	if err := db.Table(dbPkg.TableUser).Where("phone=?", code).Count(&count).Error; err != nil {
		return false, errors.Sql(err)
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// InsertUser insert user to DB
func InsertUser(db *gorm.DB, user *User) error {
	err := db.Table(dbPkg.TableUser).Create(&user).Error
	return errors.Sql(err)
}

// GetUserByPhone get user by phone
func GetUserByPhone(ctx context.Context, db *gorm.DB, phone string) (*User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "SignOut")
	defer span.Finish()

	var user User
	err := db.Table(dbPkg.TableUser).Where("phone=? and status=?", phone, StatusNormal).First(&user).Error
	if err != nil {
		return nil, errors.Sql(err)
	}
	return &user, nil
}

// GetUserByULID get user by ulid
func GetUserByULID(ctx context.Context, db *gorm.DB, ulid string) (*User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetUserByULID")
	defer span.Finish()

	var user User
	err := db.Table(dbPkg.TableUser).Where("ulid=? and status=?", ulid, StatusNormal).First(&user).Error
	if err != nil {
		return nil, errors.Sql(err)
	}
	return &user, nil
}

// GetUserByFamilyAndULID get user by family ulid and ulid
func GetUserByFamilyAndULID(ctx context.Context, db *gorm.DB, familyULID, userULID string) (*User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetUserByFamilyAndULID")
	defer span.Finish()

	var user *User
	err := db.Table(dbPkg.TableUser).Where("family_ulid=? and ulid=? and status=?", familyULID, userULID, StatusNormal).First(&user).Error
	if err != nil {
		return nil, errors.Sql(err)
	}
	return user, nil
}
