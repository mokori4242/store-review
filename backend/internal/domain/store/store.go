package store

type Store struct {
	ID              int64
	Name            string
	RegularHolidays []string
	CategoryNames   []string
	PaymentMethods  []string
	WebProfiles     []string
}

func NewStore(name string, regularHolidays, categoryNames, paymentMethods, webProfiles []string) *Store {
	return &Store{
		Name:            name,
		RegularHolidays: regularHolidays,
		CategoryNames:   categoryNames,
		PaymentMethods:  paymentMethods,
		WebProfiles:     webProfiles,
	}
}
