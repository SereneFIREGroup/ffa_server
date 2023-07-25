package four_money

import (
	"context"

	"github.com/opentracing/opentracing-go"
	fourMoneyModel "github.com/serenefiregroup/ffa_server/internal/model/four_money"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
)

type PocketMoney struct{}

func (p *PocketMoney) ListCategory(ctx context.Context) ([]string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "PocketMoney.ListCategory")
	defer span.Finish()

	return fourMoneyModel.PocketMoneyCategoryList, nil
}

func validatePocketMoney(req *AddFourMoneyRequest) error {
	if req.Amount <= 0 {
		return errors.InvalidParameterError(req.Type, errors.Amount, errors.InvalidFormat)
	}
	if !fourMoneyModel.IsValidPocketMoneyCategory(req.Category) {
		return errors.InvalidParameterError(req.Type, errors.Category, errors.InvalidParameter)
	}
	return nil
}

func (p *PocketMoney) Add(ctx context.Context, familyUlid, userUlid string, req *AddFourMoneyRequest) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "PocketMoney.Add")
	defer span.Finish()

	err := validatePocketMoney(req)
	if err != nil {
		return errors.Trace(err)
	}
	pm := fourMoneyModel.NewBasePocketMoney(familyUlid, userUlid, req.Amount, req.Category, req.Remark)
	err = fourMoneyModel.InsertPocketMoney(ctx, db.DB, pm)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (p *PocketMoney) Update(ctx context.Context, familyUlid, userUlid string, req *UpdateFourMoneyRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p *PocketMoney) List(ctx context.Context, familyUlid, userUlid string, req *ListFourMoneyRequest) (*ListFourMoneyResponse, error) {
	//TODO implement me
	panic("implement me")
}
