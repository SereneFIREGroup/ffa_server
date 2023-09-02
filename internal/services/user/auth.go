package user

import (
	"context"

	"github.com/gookit/goutil/strutil"
	"github.com/oklog/ulid/v2"
	"github.com/opentracing/opentracing-go"
	familyModel "github.com/serenefiregroup/ffa_server/internal/model/family"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	verifyModel "github.com/serenefiregroup/ffa_server/internal/model/verify"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	jaegerUtils "github.com/serenefiregroup/ffa_server/pkg/jaeger"
	"github.com/serenefiregroup/ffa_server/pkg/strings"
	"gorm.io/gorm"
)

type SignUpRequest struct {
	Phone      string `json:"phone"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	VerifyCode string `json:"verify_code"`
}

func (req *SignUpRequest) validate() error {
	req.Phone = strutil.Trim(req.Phone, " ")
	req.Name = strutil.Trim(req.Name, " ")

	if !strings.IsLenValidUTF8(req.Name, userModel.NameLength) {
		return errors.InvalidParameterError(errors.User, errors.UserName, errors.InvalidFormat)
	}
	if !strings.CheckPhoneFormat(req.Phone) {
		return errors.InvalidParameterError(errors.User, errors.Phone, errors.InvalidFormat)
	}
	return nil
}

type AuthPayload struct {
	FamilyID string `json:"family_id"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

func NewBaseAuthPayload(familyID, userID, name, avatar, token string) *AuthPayload {
	return &AuthPayload{
		FamilyID: familyID,
		ID:       userID,
		Name:     name,
		Avatar:   avatar,
		Token:    token,
	}
}

func SignUp(ctx context.Context, req *SignUpRequest) (*AuthPayload, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "SignUp")
	defer span.Finish()

	if err := req.validate(); err != nil {
		return nil, errors.Trace(err)
	}
	existPhone, err := userModel.ExistPhone(ctx, db.DB, req.Phone)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if existPhone {
		return nil, errors.InvalidParameterError(errors.User, errors.Phone, errors.AlreadyExist)
	}
	// check verify code
	vc, err := verifyModel.GetCode(db.DB, verifyModel.TypePhone, req.Phone, req.VerifyCode)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if vc == nil {
		return nil, errors.InvalidParameterError(errors.User, errors.VerifyCode, errors.InvalidFormat)
	}

	passwordMD5 := userModel.BuildPasswordMD5(req.Password)
	familyID, userID := ulid.Make().String(), ulid.Make().String()
	family := familyModel.NewBaseFamily(familyID, userID, req.Name)
	user := userModel.NewBaseUser(familyID, userID, req.Name, req.Phone, passwordMD5)
	session := userModel.NewBaseSession(userID)

	txErr := db.Transact(func(tx *gorm.DB) error {
		err := familyModel.InsertFamily(ctx, tx, family)
		if err != nil {
			return errors.Trace(err)
		}
		err = userModel.InsertUser(ctx, tx, user)
		if err != nil {
			return errors.Trace(err)
		}
		err = userModel.InsertSession(ctx, tx, session)
		return errors.Trace(err)
	})
	if txErr != nil {
		return nil, errors.Trace(txErr)
	}
	resp := NewBaseAuthPayload(familyID, userID, user.Name, user.Avatar, session.Token)
	return resp, nil
}

type SignInRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (req *SignInRequest) validate() error {
	req.Phone = strutil.Trim(req.Phone, " ")
	if !strings.CheckPhoneFormat(req.Phone) {
		return errors.InvalidParameterError(errors.User, errors.Phone, errors.InvalidFormat)
	}
	return nil
}

func SignIn(ctx context.Context, req *SignInRequest) (*AuthPayload, error) {
	span, _ := jaegerUtils.WithSpan(ctx, "SignIn")
	defer span.Finish()

	if err := req.validate(); err != nil {
		return nil, errors.Trace(err)
	}
	user, err := userModel.GetUserByPhone(ctx, db.DB, req.Phone)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if user == nil {
		return nil, errors.NotFoundError(errors.User)
	}
	passwordMD5 := userModel.BuildPasswordMD5(req.Password)
	if user.Password != passwordMD5 {
		return nil, errors.AuthFailureError(errors.IncorrectPassword)
	}

	session := userModel.NewBaseSession(user.ID)
	err = userModel.InsertSession(ctx, db.DB, session)
	if err != nil {
		return nil, errors.Trace(err)
	}
	resp := NewBaseAuthPayload(user.FamilyID, user.ID, user.Name, user.Avatar, session.Token)
	return resp, nil
}

func SignOut(ctx context.Context, userID string) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "SignOut")
	defer span.Finish()

	// found user
	user, err := userModel.GetUserByID(ctx, db.DB, userID)
	if err != nil {
		return errors.Trace(err)
	}
	if user == nil {
		return errors.NotFoundError(errors.User)
	}
	// delete session
	err = userModel.DeleteSession(ctx, db.DB, userID)
	return errors.Trace(err)
}

func Me(ctx context.Context, userID, token string) (*AuthPayload, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "Me")
	defer span.Finish()

	// found user
	user, err := userModel.GetUserByID(ctx, db.DB, userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if user == nil {
		return nil, errors.NotFoundError(errors.User)
	}
	// found session
	session, err := userModel.GetSession(ctx, db.DB, userID, token)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if session == nil {
		return nil, errors.AuthFailureError(errors.InvalidToken)
	}
	resp := NewBaseAuthPayload(user.FamilyID, user.ID, user.Name, user.Avatar, session.Token)
	return resp, nil
}
