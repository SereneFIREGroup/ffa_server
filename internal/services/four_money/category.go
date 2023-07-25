package four_money

import (
	"context"

	"github.com/opentracing/opentracing-go"
	fourMoneyModel "github.com/serenefiregroup/ffa_server/internal/model/four_money"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
)

type Categories struct {
	PocketMoneyCategoryList []string
}

// ListFourMoneyCategory list four money category
func ListFourMoneyCategory(ctx context.Context) (*Categories, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListFourMoneyCategory")
	defer span.Finish()

	pms, _ := getFourMoneyService(fourMoneyModel.TypePocketMoney)
	pmCategories, err := pms.ListCategory(ctx)
	if err != nil {
		return nil, errors.Trace(err)
	}

	resp := new(Categories)
	resp.PocketMoneyCategoryList = pmCategories
	return resp, nil
}
