package operator
// Operator Таблица операторы. Сотовые операторы.
type Operator struct {
	ID   string 
	Name string
}

// CreateTable Возвращает строку создания таблицы
func (s Operator) CreateTable() string {
	return `
	 CREATE TABLE Operator
		(ID CHAR(36) , Name CHAR(100));
		`
}