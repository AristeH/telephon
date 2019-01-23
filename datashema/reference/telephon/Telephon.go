package telephon

// Telephon Таблица Телефоны. Содержит код и номер телефона.
type Telephon struct {
	ID     uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Code   string
	Nomer  string
	Number string
	IDRegion string
	Our    bool
}
