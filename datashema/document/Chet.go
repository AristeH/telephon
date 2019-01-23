package document

import "time"
import "Telephon/datashema/reference"

// Chet Таблица Chet(СчетОператора). Счет за услуги выставленный за звонки телефона.
type Chet struct {
	ID           uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Telephon     reference.Telephon
	Date         time.Time
	IDchetzvonki uint64
	Nomerscheta  string
	Region       reference.Region
}
