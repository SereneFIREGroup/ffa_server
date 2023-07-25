package four_money

import (
	"context"

	fourMoneyModel "github.com/serenefiregroup/ffa_server/internal/model/four_money"
)

type Categories struct {
	PocketMoneyCategoryList []string
}

func ListFourMoneyCategory(spanCtx context.Context) (*Categories, error) {
	resp := new(Categories)
	resp.PocketMoneyCategoryList = fourMoneyModel.PocketMoneyCategoryList

	return resp, nil
}
