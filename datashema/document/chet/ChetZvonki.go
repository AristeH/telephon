package document

import (
	ref "Telephon/datashema/reference"
	"time"
)

// Таблица Звонки. Все звонки телефона.
type ChetZvonki struct {
	Date    time.Time   `json:"date,omitempty"`
	Number  string      `json:"number,omitempty"`
	Amount  float64     `json:"amount,omitempty"`
	Unit    ref.Unit    `json:"unit,omitempty"`
	Service ref.Service `json:"service,omitempty"`
	Region  ref.Region  `json:"region,omitempty"`
	Price   float64     `json:"price,omitempty"`
}
