package user



// User структура люди.
type User struct {
	ID   string `json:"УникальныйИдентификатор"`
	Code string `json:"Код"`
	Name string `json:"Наименование"`
	Password string `json:"Пароль"`
}

// Users список сотрудников
type Users struct {
	users []User 
}

// CreateTable Возвращает строку создания таблицы
func (s User) CreateTable() string {
	return `
	 CREATE TABLE People
		(ID CHAR(36) , Name CHAR(100), Code  CHAR(10));
		`
}
// Insert добавляет запись в таблицу
func (s User) Insert(р User) string {
	z := "UPDATE OR INSERT INTO People(ID, Name, Code)"+
		"VALUES ('" + р.ID + "','" + р.Name+ "','"+ р.Code  +"')" +
		"MATCHING(ID);"
	return z
		
} 
 


