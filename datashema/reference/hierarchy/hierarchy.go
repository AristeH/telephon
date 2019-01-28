package hierarchy

// Hierarchy Таблица иерархий, хранится в Json формате
type Hierarchy struct {
	ID   string
	Name string
	Text string	
	WeightLeft int
	WeightRight int
}

//CreateTable Возвращает строку создания таблицы
func (s Hierarchy) CreateTable() string {
	return `
CREATE TABLE Hierarchy
 (ID CHAR(36) , Name CHAR(100), Text BLOB, WeightLeft INTEGER, WeightRight INTEGER  );
 `
}