package store

type Store struct {
	ID              int64
	Name            string
	RegularHolidays []string
	CategoryNames   []string
	PaymentMethods  []string
	WebProfiles     []string
}
