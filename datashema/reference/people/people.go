package people

// People структура люди.
type People struct {
	ID   string `json:"УникальныйИдентификатор"`
	Code string `json:"Код"`
	Name string `json:"Наименование"`
}

// Sotrs список сотрудников
type Sotrs struct {
	Sotrs []People `json:"6_ФизическиеЛица"`
}

// CreateTable Возвращает строку создания таблицы
func (s People) CreateTable() string {
	return `
	CREATE TABLE PEOPLE (
		ID CHAR(36),
		NAME CHAR(100),
		CODE CHAR(20),
		CONSTRAINT PEOPLE_PK PRIMARY KEY (ID)
	);
	`
}
// Insert добавляет запись в таблицу
func (s People) Insert(р People) string {
	z := "UPDATE OR INSERT INTO People(ID, Name, Code)"+
		"VALUES ('" + р.ID + "','" + р.Name+ "','"+ р.Code  +"')" +
		"MATCHING(ID);"
	return z
		
} 
 


