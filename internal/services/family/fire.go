package family

import (
	"context"

	familyModel "github.com/serenefiregroup/ffa_server/internal/model/family"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	jaegerUtils "github.com/serenefiregroup/ffa_server/pkg/jaeger"
)

type SetFIREGoldRequest struct {
	Amount int64 `json:"amount"`
}

func (req *SetFIREGoldRequest) Validate() error {
	if req.Amount <= 0 {
		return errors.InvalidParameterError(errors.Family, errors.FIREGold, errors.InvalidFormat)
	}
	return nil
}

func SetFIREGold(ctx context.Context, familyULID, userULID string, req *SetFIREGoldRequest) error {
	span, _ := jaegerUtils.WithSpan(ctx, "SetFIREGold")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return errors.Trace(err)
	}

	family, err := familyModel.GetFamily(ctx, db.DB, familyULID)
	if err != nil {
		return errors.Trace(err)
	}
	if family == nil {
		return errors.NotFoundError(errors.Family)
	}

	user, err := userModel.GetUserByULID(ctx, db.DB, userULID)
	if err != nil {
		return errors.Trace(err)
	}
	if user == nil {
		return errors.NotFoundError(errors.User)
	}
	if family.Owner != userULID {
		return errors.AccessDeniedError(errors.Family, errors.InvalidOwner)
	}
	family.SetFIREGold(req.Amount)
	err = familyModel.UpdateFamily(ctx, db.DB, family)
	return errors.Trace(err)
}
