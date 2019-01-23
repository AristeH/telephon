package md

// Tables содержит имена и описания всех таблиц базы.
// name имя таблицы в БД.
// tname представление таблицы в приложении 
// info описание таблицы 
type Tables struct {
	ID   string      
	Name string
	Tname string
	Info string
}
