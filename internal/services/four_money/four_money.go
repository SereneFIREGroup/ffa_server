package four_money

import (
	"context"

	familyModel "github.com/serenefiregroup/ffa_server/internal/model/family"
	fourMoneyModel "github.com/serenefiregroup/ffa_server/internal/model/four_money"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	jaegerUtils "github.com/serenefiregroup/ffa_server/pkg/jaeger"
)

var hub map[string]FourMoney

type FourMoney interface {
	ListCategory(ctx context.Context) (*Categories, error)
	Add(ctx context.Context, familyUlid, userUlid string, req *AddFourMoneyRequest) error
	Update(ctx context.Context, familyUlid, userUlid string, req *UpdateFourMoneyRequest) error
	List(ctx context.Context, familyUlid, userUlid string, req *ListFourMoneyRequest) (*ListFourMoneyResponse, error)
}

type UpdateFourMoneyRequest struct{}

type ListFourMoneyResponse struct{}

type ListFourMoneyRequest struct{}

func init() {
	hub = make(map[string]FourMoney)
	hub[fourMoneyModel.TypePocketMoney] = new(PocketMoney)
}

type AddFourMoneyRequest struct {
	Type     string `json:"type"`
	Amount   int64  `json:"amount"`
	Category string `json:"category"`
	Remark   string `json:"remark"`
}

func (req *AddFourMoneyRequest) Validate() error {
	if !fourMoneyModel.IsValidType(req.Type) {
		return errors.InvalidParameterError(errors.FourMoney, errors.Type, errors.InvalidParameter)
	}
	if req.Amount <= 0 {
		return errors.InvalidParameterError(req.Type, errors.Amount, errors.InvalidFormat)
	}
	if !fourMoneyModel.IsValidPocketMoneyCategory(req.Category) {
		return errors.InvalidParameterError(req.Type, errors.Category, errors.InvalidParameter)
	}
	return nil
}

func AddFourMoney(ctx context.Context, familyUlid, userUlid string, req *AddFourMoneyRequest) error {
	span, _ := jaegerUtils.WithSpan(ctx, "AddFourMoney")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return errors.Trace(err)
	}

	family, err := familyModel.GetFamily(ctx, db.DB, familyUlid)
	if err != nil {
		return errors.Trace(err)
	}
	if family == nil {
		return errors.NotFoundError(errors.Family)
	}

	user, err := userModel.GetUserByULID(ctx, db.DB, userUlid)
	if err != nil {
		return errors.Trace(err)
	}
	if user == nil {
		return errors.NotFoundError(errors.User)
	}

	//pm := fourMoneyModel.NewBasePocketMoney(familyUlid, userUlid, req.Amount, req.Category, req.Remark)
	//return InsertPocketMoney(spanCtx, db, pm)
	return nil
}
