package unit
// Unit Таблица ЕдиницыИзмерений.
type Unit struct {
	ID   uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name string
}