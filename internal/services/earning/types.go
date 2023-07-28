package earning

import (
	earningCareerModel "github.com/serenefiregroup/ffa_server/internal/model/earning_career"
)

type EarningPayload struct {
	Amount   int64  `json:"amount"`
	Date     int64  `json:"date"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
}

type YearlyEarningPayload struct {
	Year   int   `json:"year"`
	Amount int64 `json:"amount"`
}

// Model2Payload converts model to payload
func Model2Payload(earning *earningCareerModel.Earning) *EarningPayload {
	return &EarningPayload{
		Amount:   earning.Amount,
		Date:     earning.CreateTime,
		Category: earning.Category,
		Desc:     earning.Desc,
	}
}

// ModelList2PayloadList converts model list to payload list
func ModelList2PayloadList(earnings []*earningCareerModel.Earning) []*EarningPayload {
	var ret []*EarningPayload
	for _, earning := range earnings {
		ret = append(ret, Model2Payload(earning))
	}
	return ret
}

// CalcTotalAmount calculates total amount of earnings
func CalcTotalAmount(earnings []*earningCareerModel.Earning) int64 {
	var totalAmount int64
	for _, earning := range earnings {
		totalAmount += earning.Amount
	}
	return totalAmount
}
