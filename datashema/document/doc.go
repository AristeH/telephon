package document

import (
	"time"
)

// Таблица СчетОператора. Счет за услуги выставленный за звонки телефона.
type Doc struct {
	Date        time.Time    `json:"date,omitempty"`
	Nomer string       `json:"nomerscheta,omitempty"`
}