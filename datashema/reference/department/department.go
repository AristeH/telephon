package department


// Department структура Подразделения.
type Department struct {
	ID     string `json:"УникальныйИдентификатор"`
	Name   string `json:"Наименование"`
	Code   string `json:"Код"`
	Parent string `json:"Родитель"`
	WeithB  string `json:"Order"`
}

// Departments список подразделений
type Departments struct {
	Departments []Department `json:"Подразделения"`
}

// CreateTable Возвращает строку создания таблицы
func (s Department) CreateTable() string {
	return `
	    CREATE TABLE Department(
			ID CHAR(36),  
			Parent CHAR(36),
			Code CHAR(10),
			Name CHAR(150),
			WeithB INTEGER);
		`
}

// Insert добавляет запись в таблицу
func (s Department) Insert(р Department) string {
	z := "UPDATE OR INSERT INTO Department(ID, Parent, Code, Name, WeithB)" +
		"VALUES ('" + р.ID + "','" + р.Parent + "','" + р.Code + "','" + р.Name+ "','"+ р.WeithB + "')" +
		"MATCHING(ID);"
	return z

}

