package four_money

const (
	TypePocketMoney = "pocket_money"

	PocketMoneyCategoryWechat   = "wechat"
	PocketMoneyCategoryYueBao   = "yuebao"
	PocketMoneyCategoryBankCard = "bank_card"
	PocketMoneyCategoryOther    = "other"
)

var TypeMap = map[string]string{
	TypePocketMoney: "零花钱",
}

var PocketMoneyCategoryMap = map[string]string{
	PocketMoneyCategoryWechat:   "微信",
	PocketMoneyCategoryYueBao:   "余额宝",
	PocketMoneyCategoryBankCard: "银行卡",
	PocketMoneyCategoryOther:    "其他",
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
