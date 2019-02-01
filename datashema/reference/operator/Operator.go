package operator

// Operator Таблица операторы. Сотовые операторы.
type Operator struct {
	ID   string
	Name string
}

// CreateTable Возвращает строку создания таблицы
func (s Operator) CreateTable() string {
	return `
	CREATE TABLE OPERATOR (
		ID CHAR(36),
		NAME CHAR(100),
		CONSTRAINT OPERATOR_PK PRIMARY KEY (ID)
	);
`
}
