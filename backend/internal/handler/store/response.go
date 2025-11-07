package store

type ListResponse struct {
	Stores []Response
}

type Response struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	RegularHolidays []string `json:"regularHolidays"`
	CategoryNames   []string `json:"categoryNames"`
	PaymentMethods  []string `json:"paymentMethods"`
	WebProfiles     []string `json:"webProfiles"`
}
