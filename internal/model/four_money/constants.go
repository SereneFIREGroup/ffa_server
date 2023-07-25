package four_money

const (
	TypePocketMoney = "pocket_money"

	PocketMoneyCategoryWechat   = "wechat"
	PocketMoneyCategoryYueBao   = "yuebao"
	PocketMoneyCategoryBankCard = "bank_card"
	PocketMoneyCategoryOther    = "other"
)

var TypeMap = map[string]struct{}{
	TypePocketMoney: {},
}

var PocketMoneyCategoryMap = map[string]struct{}{
	PocketMoneyCategoryWechat:   {},
	PocketMoneyCategoryYueBao:   {},
	PocketMoneyCategoryBankCard: {},
	PocketMoneyCategoryOther:    {},
}

// PocketMoneyCategoryList PocketMoneyCategoryList
var PocketMoneyCategoryList = []string{
	PocketMoneyCategoryWechat,
	PocketMoneyCategoryYueBao,
	PocketMoneyCategoryBankCard,
	PocketMoneyCategoryOther,
}

// IsValidType check if type is valid
func IsValidType(t string) bool {
	_, ok := TypeMap[t]
	return ok
}

// IsValidPocketMoneyCategory check if category is valid
func IsValidPocketMoneyCategory(category string) bool {
	_, ok := PocketMoneyCategoryMap[category]
	return ok
}
