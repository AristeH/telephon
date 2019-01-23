package oblast

// Oblast Таблица области. Содержит в себе области РФ
type Oblast struct {
	ID   string  
	Name string
}

//CreateTable Возвращает строку создания таблицы
func (s Oblast) CreateTable() string {
	return `
CREATE TABLE OBLAST
 (ID CHAR(36) , Name CHAR(100));
 `
}