package post

// Post структура должности.
type Post struct {
	ID   string `json:"УникальныйИдентификатор"`
	Name string `json:"Наименование"`
}

// Posts список должностей
type Posts struct {
	Posts []Post `json:"1_ДолжностиОрганизаций"`
}

// CreateTable Возвращает строку создания таблицы
func (s Post) CreateTable() string {
	return `
	CREATE TABLE POST (
		ID CHAR(36),
		NAME CHAR(100),
		CONSTRAINT POST_PK PRIMARY KEY (ID)
	);
		`
}

// Insert добавляет запись в таблицу
func (s Post) Insert(р Post) string {
	z := "UPDATE OR INSERT INTO Post(ID, Name) "+
		"VALUES ('" + р.ID + "','" + р.Name  +"') " +
		"MATCHING(ID);"
	return z
		
} 

