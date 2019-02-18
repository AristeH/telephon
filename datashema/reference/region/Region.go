package region
// Region Таблица Регионы.
type Region struct {
	ID     string `json:"УникальныйИдентификатор"`
	Name   string `json:"Наименование"`
	IDoblast string `json:"область"`
}
// CreateTable  Возвращает строку создания таблицы
func (s Region) CreateTable() string {
	return "CREATE TABLE Region(ID CHAR(36) , Name CHAR(50),  IDoblast CHAR(36),CONSTRAINT Region_PK PRIMARY KEY (ID) );"
}
