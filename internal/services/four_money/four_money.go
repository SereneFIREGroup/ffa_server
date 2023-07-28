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
	ListCategory(ctx context.Context) ([]string, error)
	Add(ctx context.Context, familyID, userID string, req *AddFourMoneyRequest) error
	Update(ctx context.Context, familyID, userID string, req *UpdateFourMoneyRequest) error
	List(ctx context.Context, familyID, userID string, req *ListFourMoneyRequest) (*ListFourMoneyResponse, error)
}

type UpdateFourMoneyRequest struct{}

type ListFourMoneyResponse struct{}

type ListFourMoneyRequest struct{}

func init() {
	hub = make(map[string]FourMoney)
	hub[fourMoneyModel.TypePocketMoney] = new(PocketMoney)
}

func getFourMoneyService(t string) (FourMoney, error) {
	s, ok := hub[t]
	if !ok {
		return nil, errors.InvalidParameterError(errors.FourMoney, errors.Type, errors.InvalidParameter)
	}
	return s, nil
}

type AddFourMoneyRequest struct {
	Type     string `json:"type"`
	Amount   int64  `json:"amount"`
	Category string `json:"category"`
	Remark   string `json:"remark"`
}

func AddFourMoney(ctx context.Context, familyID, userID string, req *AddFourMoneyRequest) error {
	span, _ := jaegerUtils.WithSpan(ctx, "AddFourMoney")
	defer span.Finish()

	family, err := familyModel.GetFamily(ctx, db.DB, familyID)
	if err != nil {
		return errors.Trace(err)
	}
	if family == nil {
		return errors.NotFoundError(errors.Family)
	}

	user, err := userModel.GetUserByID(ctx, db.DB, userID)
	if err != nil {
		return errors.Trace(err)
	}
	if user == nil {
		return errors.NotFoundError(errors.User)
	}

	s, err := getFourMoneyService(req.Type)
	if err != nil {
		return errors.Trace(err)
	}
	err = s.Add(ctx, familyID, userID, req)
	return errors.Trace(err)
}
