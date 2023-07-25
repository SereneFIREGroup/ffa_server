package earning_career

// Category of earning
const (
	CategorySalary     = "salary"
	CategoryBonus      = "bonus"
	CategoryInvestment = "investment"
	CategoryOther      = "other"
)

const (
	DescLen = 256
)

var (
	// CategoryList is the list of earning category
	CategoryList = []string{
		CategorySalary,
		CategoryBonus,
		CategoryInvestment,
		CategoryOther,
	}
)
