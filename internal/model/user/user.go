package user

import (
	"context"
	goErrors "errors"
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
	FamilyID   string `json:"family_id" gorm:"column:family_id"`
	ID         string `json:"id" gorm:"primarykey,column:id"`
	Name       string `json:"name" gorm:"column:name"`
	Avatar     string `json:"avatar" gorm:"column:avatar"`
	Phone      string `json:"phone" gorm:"column:phone"`
	Password   string `json:"password" gorm:"column:password"`
	Status     uint8  `json:"status" gorm:"column:status"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
}

func NewBaseUser(familyID, id, name, phone, password string) *User {
	return &User{
		FamilyID:   familyID,
		ID:         id,
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
func ExistPhone(ctx context.Context, db *gorm.DB, code string) (bool, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ExistPhone")
	defer span.Finish()

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
func InsertUser(ctx context.Context, db *gorm.DB, user *User) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "InsertUser")
	defer span.Finish()

	err := db.Table(dbPkg.TableUser).Create(&user).Error
	return errors.Sql(err)
}

// GetUserByPhone get user by phone
func GetUserByPhone(ctx context.Context, db *gorm.DB, phone string) (*User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetUserByPhone")
	defer span.Finish()

	var user User
	err := db.Table(dbPkg.TableUser).Where("phone=? and status=?", phone, StatusNormal).First(&user).Error
	if goErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Sql(err)
	}
	return &user, nil
}

// GetUserByID get user by ulid
func GetUserByID(ctx context.Context, db *gorm.DB, id string) (*User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetUserByID")
	defer span.Finish()

	var user User
	err := db.Table(dbPkg.TableUser).Where("id=? and status=?", id, StatusNormal).First(&user).Error
	if goErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Sql(err)
	}
	return &user, nil
}

// GetUserByFamilyAndID get user by family ulid and ulid
func GetUserByFamilyAndID(ctx context.Context, db *gorm.DB, familyID, userID string) (*User, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetUserByFamilyAndID")
	defer span.Finish()

	var user *User
	err := db.Table(dbPkg.TableUser).Where("family_id=? and id=? and status=?", familyID, userID, StatusNormal).First(&user).Error
	if goErrors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Sql(err)
	}
	return user, nil
}
