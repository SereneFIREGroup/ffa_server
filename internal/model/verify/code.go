package verify

import (
	"time"

	dbPkg "github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"gorm.io/gorm"
)

const (
	TypeEmail = iota + 1
	TypePhone

	ExpireTime = 60 * 5

	StatusNormal  = 1
	StatusDisable = 2
)

type VerifyCode struct {
	VerifyType int    `json:"verify_type" gorm:"column:verify_type"`
	Number     string `json:"number" gorm:"column:number"`
	Code       string `json:"code" gorm:"column:code"`
	CreateTime int64  `json:"create_time" gorm:"column:create_time"`
	IP         string `json:"ip" gorm:"column:ip"`
	Status     int8   `json:"status" gorm:"column:status"`
}

func NewEmailVerifyCode(number, code, ip string) *VerifyCode {
	return NewVerifyCode(TypeEmail, number, code, ip)
}

func NewPhoneVerifyCode(number, code, ip string) *VerifyCode {
	return NewVerifyCode(TypePhone, number, code, ip)
}

func NewVerifyCode(verifyType int, number, code, ip string) *VerifyCode {
	return &VerifyCode{
		VerifyType: verifyType,
		Number:     number,
		Code:       code,
		CreateTime: time.Now().Unix(),
		IP:         ip,
		Status:     StatusNormal,
	}
}

// IsExpired check verify code is expired
func (v *VerifyCode) IsExpired() bool {
	return time.Now().Unix()-v.CreateTime > ExpireTime
}

// CreateVerifyCode create verify code
func CreateVerifyCode(db *gorm.DB, code *VerifyCode) error {
	return errors.Sql(db.Table(dbPkg.TableVerifyCode).Create(code).Error)
}

// GetCode get verify code
func GetCode(db *gorm.DB, verifyType int, number, code string) (*VerifyCode, error) {
	var vc *VerifyCode
	if err := db.Table(dbPkg.TableVerifyCode).
		Where("verify_type=? AND number=? and code=? and status=?", verifyType, number, code, StatusNormal).
		Order("create_time DESC").First(&vc).Error; err != nil {
		return nil, errors.Sql(err)
	}
	return vc, nil
}

// ClearVerifyCode clear verify code
func ClearVerifyCode(db *gorm.DB, number string) error {
	return errors.Sql(db.Table(dbPkg.TableVerifyCode).Where("number=?", number).Update("status", StatusDisable).Error)
}
