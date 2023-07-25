package user

import (
	"context"
	gostrings "strings"

	"github.com/opentracing/opentracing-go"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	verifyModel "github.com/serenefiregroup/ffa_server/internal/model/verify"
	concUtils "github.com/serenefiregroup/ffa_server/pkg/conc_utils"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"github.com/serenefiregroup/ffa_server/pkg/log"
	"github.com/serenefiregroup/ffa_server/pkg/random"
	"github.com/serenefiregroup/ffa_server/pkg/sms"
	"github.com/serenefiregroup/ffa_server/pkg/strings"
)

type VerifyPhoneRequest struct {
	Phone string `json:"phone"`
}

func (req *VerifyPhoneRequest) validate() error {
	req.Phone = gostrings.TrimSuffix(gostrings.TrimPrefix(req.Phone, " "), " ")
	if !strings.CheckPhoneFormat(req.Phone) {
		return errors.InvalidParameterError(errors.User, errors.Phone, errors.InvalidFormat)
	}
	return nil
}

func VerifyPhone(ctx context.Context, req *VerifyPhoneRequest, ip string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "VerifyPhone")
	defer span.Finish()

	if err := req.validate(); err != nil {
		return errors.Trace(err)
	}
	existPhone, err := userModel.ExistPhone(db.DB, req.Phone)
	if err != nil {
		return errors.Trace(err)
	}
	if existPhone {
		return errors.InvalidParameterError(errors.User, errors.Phone, errors.AlreadyExist)
	}
	// clear old verify code
	err = verifyModel.ClearVerifyCode(db.DB, req.Phone)
	if err != nil {
		return errors.Trace(err)
	}
	code := random.StringNumber(3, "")
	verifyCode := verifyModel.NewPhoneVerifyCode(req.Phone, code, ip)
	err = verifyModel.CreateVerifyCode(db.DB, verifyCode)
	if err != nil {
		return errors.Trace(err)
	}
	concUtils.GoSafe(func() {
		ok, sendErr := sms.PostSignUpVerifySMS(req.Phone, code)
		if sendErr != nil || !ok {
			log.Warn("[VerifyPhone.PostSignUpVerifySMS] phone: %s, sendErr: %v, ok: %v", req.Phone, sendErr, ok)
		}
	})
	return nil
}
