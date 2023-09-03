package family

import (
	"context"

	familyModel "github.com/serenefiregroup/ffa_server/internal/model/family"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	jaegerUtils "github.com/serenefiregroup/ffa_server/pkg/jaeger"
)

type InfoResp struct {
	ID          string `json:"id"`
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	FireGoal    int64  `json:"fire_goal"`
	TotalAssets int64  `json:"total_assets"`
}

func Info(ctx context.Context, familyID, userID string) (*InfoResp, error) {
	span, _ := jaegerUtils.WithSpan(ctx, "GetFamilyInfo")
	defer span.Finish()

	user, err := userModel.GetUserByID(ctx, db.DB, userID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if user == nil {
		return nil, errors.NotFoundError(errors.User)
	}
	family, err := familyModel.GetFamily(ctx, db.DB, familyID)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if family == nil {
		return nil, errors.NotFoundError(errors.Family)
	}
	resp := new(InfoResp)
	resp.ID = family.ID
	resp.Owner = family.Owner
	resp.Name = family.Name
	resp.FireGoal = family.FIREGoal
	resp.TotalAssets = family.FIREGoal // TODO: 计算整个家庭的总资产

	return resp, nil
}
