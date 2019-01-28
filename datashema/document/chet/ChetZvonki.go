package document

import (
	"time"
)

// ChetZvonki Таблица Звонки. Все звонки телефона.
type ChetZvonki struct {
	Date      time.Time `json:"date,omitempty"`
	Number    string    `json:"number,omitempty"`
	Amount    float64   `json:"amount,omitempty"`
	IDUnit    string    `json:"unit,omitempty"`
	IDService string    `json:"service,omitempty"`
	IDRegion  string    `json:"region,omitempty"`
	Price     float64   `json:"price,omitempty"`
}
