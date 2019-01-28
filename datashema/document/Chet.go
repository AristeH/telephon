package document

import "time"


// Chet Таблица Chet(СчетОператора). Счет за услуги выставленный за звонки телефона.
type Chet struct {
	ID           uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	IDTelephon   string
	Date         time.Time
	IDchetzvonki string
	Nomerscheta  string
	IDRegion     string
}
