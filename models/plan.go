package models

import (
	"time"
)

var EXPIRED_TIMES = [4]string{"day", "days", "month", "months"}

// time.Now() => Thu, 12 Mar 2015 21:17:33 -0500
// plan: {
//             Currency: "COP",
//           DeletedAt: nil,
//          Description: "15% de ahorro",
//         ExpiredRate: 1,
//         ExpiredTime: "month",
//                 Name: "Hurry Up!",
//     NumOfDiscounts: 15,
//                Price: 127500,
//               Status: true
// }
// the purchased plan ExpiredDate attribute = Thu, 12 Apr 2015 21:17:33 -0500
// 1 month after today

type Plan struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Status         bool   `json:"status"`
	Price          uint   `json:"price"`
	NumOfDiscounts uint   `json:"num_of_discounts"`
	Currency       string `json:"currency"`
	ExpiredRate    uint   `json:"expired_rate"` // (1 .. 31) days or (1 .. 12) months
	ExpiredTime    string `json:"expired_time"` // ( days or months )
}

type ClientPlan struct {
	Plan
	ExpiredDate        *time.Time `json:"expired_date"`
	NumOfDiscountsLeft uint       `json:"num_of_discounts_left"`
}
