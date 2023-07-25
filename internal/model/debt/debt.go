package debt

// debt 债务，欠款

type Debt struct {
	ULID       string `json:"ulid" gorm:"primarykey,column:ulid"`
	FamilyULID string
	UserULID   string
	Debtor     string
	Amount     int64
}
