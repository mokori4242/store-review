package store

import "store-review/internal/domain/store"

type response struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	RegularHolidays []string `json:"regularHolidays"`
	CategoryNames   []string `json:"categoryNames"`
	PaymentMethods  []string `json:"paymentMethods"`
	WebProfiles     []string `json:"webProfiles"`
}

func newResponse(s *store.Store) response {
	return response{
		ID:              s.ID,
		Name:            s.Name,
		RegularHolidays: parseHolidays(s.RegularHolidays),
		CategoryNames:   s.CategoryNames,
		PaymentMethods:  parsePaymentMethods(s.PaymentMethods),
		WebProfiles:     s.WebProfiles,
	}
}

func parseHolidays(hs []string) []string {
	var parsedHolidays []string
	for _, h := range hs {
		switch h {
		case "0":
			parsedHolidays = append(parsedHolidays, "月曜")
		case "1":
			parsedHolidays = append(parsedHolidays, "火曜")
		case "2":
			parsedHolidays = append(parsedHolidays, "水曜")
		case "3":
			parsedHolidays = append(parsedHolidays, "木曜")
		case "4":
			parsedHolidays = append(parsedHolidays, "金曜")
		case "5":
			parsedHolidays = append(parsedHolidays, "土曜")
		case "6":
			parsedHolidays = append(parsedHolidays, "日曜")
		}
	}
	return parsedHolidays
}

func parsePaymentMethods(ps []string) []string {
	var parsedPaymentMethods []string
	for _, p := range ps {
		switch p {
		case "0":
			parsedPaymentMethods = append(parsedPaymentMethods, "PayPay")
		case "1":
			parsedPaymentMethods = append(parsedPaymentMethods, "現金")
		case "2":
			parsedPaymentMethods = append(parsedPaymentMethods, "RakutenPay")
		}
	}
	return parsedPaymentMethods
}
