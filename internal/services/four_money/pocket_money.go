package four_money

import "context"

type PocketMoney struct{}

func (p *PocketMoney) ListCategory(ctx context.Context) (*Categories, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PocketMoney) Add(ctx context.Context, familyUlid, userUlid string, req *AddFourMoneyRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p *PocketMoney) Update(ctx context.Context, familyUlid, userUlid string, req *UpdateFourMoneyRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p *PocketMoney) List(ctx context.Context, familyUlid, userUlid string, req *ListFourMoneyRequest) (*ListFourMoneyResponse, error) {
	//TODO implement me
	panic("implement me")
}
