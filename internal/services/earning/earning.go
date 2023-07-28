package earning

import (
	"context"
	"math"

	"github.com/opentracing/opentracing-go"
	earningCareerModel "github.com/serenefiregroup/ffa_server/internal/model/earning_career"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	datePkg "github.com/serenefiregroup/ffa_server/pkg/date"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"github.com/serenefiregroup/ffa_server/pkg/slice"
	stringsUtils "github.com/serenefiregroup/ffa_server/pkg/strings"
)

type ListEarningCategoryResp struct {
	Categories []string `json:"category"`
}

func ListEarningCategory(ctx context.Context) (*ListEarningCategoryResp, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListEarningCategory")
	defer span.Finish()

	return &ListEarningCategoryResp{Categories: earningCareerModel.CategoryList}, nil
}

type AddEarningRequest struct {
	Amount   int64  `json:"amount"`
	Date     int64  `json:"date"`
	Category string `json:"category"`
	Desc     string `json:"desc"`
}

func (req *AddEarningRequest) Valid() error {
	if req.Amount <= 0 {
		return errors.InvalidParameterError(errors.Earning, errors.Amount, errors.InvalidFormat)
	}
	if !slice.Exist(earningCareerModel.CategoryList, req.Category) {
		return errors.InvalidParameterError(errors.Earning, errors.Category, errors.InvalidEnum)
	}
	if !stringsUtils.IsLenValidUTF8(req.Desc, earningCareerModel.DescLen) {
		return errors.InvalidParameterError(errors.Earning, errors.Desc, errors.InvalidFormat)
	}
	return nil
}

func AddEarning(ctx context.Context, userID string, req *AddEarningRequest) error {
	span, _ := opentracing.StartSpanFromContext(ctx, "AddEarning")
	defer span.Finish()

	if err := req.Valid(); err != nil {
		return errors.Trace(err)
	}

	user, err := userModel.GetUserByID(ctx, db.DB, userID)
	if err != nil {
		return errors.Trace(err)
	}
	if user == nil {
		return errors.NotFoundError(errors.User)
	}
	earning := earningCareerModel.NewBaseEarning(userID, req.Amount, req.Category, req.Desc)
	err = earningCareerModel.InsertEarningCareer(ctx, db.DB, earning)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

type ListEarningRequest struct {
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
}

func (req *ListEarningRequest) Valid() error {
	if req.StartDate <= 0 {
		return errors.InvalidParameterError(errors.Earning, errors.StartDate, errors.InvalidFormat)
	}
	if req.EndDate <= 0 {
		return errors.InvalidParameterError(errors.Earning, errors.EndDate, errors.InvalidFormat)
	}
	if req.StartDate > req.EndDate {
		return errors.InvalidParameterError(errors.Earning, errors.StartDate, errors.InvalidFormat)
	}
	return nil
}

type ListEarningResp struct {
	Total    int64             `json:"total"`
	Earnings []*EarningPayload `json:"earnings"`
}

func ListEarning(ctx context.Context, userID string, req *ListEarningRequest) (*ListEarningResp, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListEarning")
	defer span.Finish()

	if err := req.Valid(); err != nil {
		return nil, errors.Trace(err)
	}

	user, err := userModel.GetUserByID(ctx, db.DB, userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if user == nil {
		return nil, errors.NotFoundError(errors.User)
	}

	startDateStamp := datePkg.DateStart(req.StartDate)
	endDateStamp := datePkg.DateEnd(req.EndDate)
	earnings, err := earningCareerModel.ListEarningCareer(ctx, db.DB, userID, startDateStamp, endDateStamp)
	if err != nil {
		return nil, errors.Trace(err)
	}
	total := 0
	resp := new(ListEarningResp)
	resp.Earnings = make([]*EarningPayload, len(earnings), len(earnings))
	for i, earning := range earnings {
		total += int(earning.Amount)
		resp.Earnings[i] = Model2Payload(earning)
	}
	return resp, nil
}

type AggrEarningResp struct {
	Total        int64                   `json:"total"`
	YearEarnings []*YearlyEarningPayload `json:"year_earnings"`
}

func AggregationEarning(ctx context.Context, userID string) (*AggrEarningResp, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListEarning")
	defer span.Finish()

	user, err := userModel.GetUserByID(ctx, db.DB, userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if user == nil {
		return nil, errors.NotFoundError(errors.User)
	}

	earnings, err := earningCareerModel.ListEarningCareer(ctx, db.DB, userID, earningCareerModel.UnLimitTime, earningCareerModel.UnLimitTime)
	if err != nil {
		return nil, errors.Trace(err)
	}
	total := 0
	resp := new(AggrEarningResp)
	resp.YearEarnings = make([]*YearlyEarningPayload, 0, len(earnings))
	preYear := math.MinInt
	for _, earning := range earnings {
		total += int(earning.Amount)
		year := earning.GetYear()
		if year != preYear {
			resp.YearEarnings = append(resp.YearEarnings, &YearlyEarningPayload{
				Year:   year,
				Amount: 0,
			})
		}
		resp.YearEarnings[len(resp.YearEarnings)-1].Amount += earning.Amount
	}
	return resp, nil
}
