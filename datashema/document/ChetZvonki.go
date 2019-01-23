package document

import "time"

// ChetZvonki Таблица Звонки. Все звонки телефона.
type ChetZvonki struct {
	Date      time.Time
	Number    string
	Amount    float64
	IDunit    uint64
	IDservice uint64
	IDregion  uint64
	Price     float64
	IDchet    uint64
}
