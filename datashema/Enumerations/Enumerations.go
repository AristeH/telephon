package enumeration

import (
	"strconv"
)

// Enumeration структура Перечисления.
type Enumeration struct {
	ID     string `json:"УникальныйИдентификатор"`
	Name   string `json:"Наименование"`
	Code   string `json:"Код"`
	Parent string `json:"Родитель"`
	WeithB int `json:"Уровень"`
}

// Enumerations список Перечислений
type Enumerations struct {
	Enumerations []Enumeration `json:"Перечисления"`
}

// CreateTable Возвращает строку создания таблицы
func (s Enumeration) CreateTable() string {
	return `
CREATE TABLE Enumeration (
	ID CHAR(36) NOT NULL,
	TYPE CHAR(36),
	CODE CHAR(10),
	NAME CHAR(150),
	CONSTRAINT Enumeration_PK PRIMARY KEY (ID)
);
`
}
// CreateIndex Возвращает строку создания таблицы
func (s Enumeration) CreateIndex() string {
	return `
CREATE INDEX Enumeration_PARENT_IDX ON Enumeration (PARENT,CODE,WEITHB,NAME);
);
`
}
// Insert добавляет запись в таблицу
func (s Enumeration) Insert(р Enumeration) string {
	z := "UPDATE OR INSERT INTO Department(ID, Parent, Code, Name, WeithB)" +
		"VALUES ('" + р.ID + "','" + р.Parent + "','" + р.Code + "','" + р.Name + "','" + strconv.Itoa(р.WeithB) + "')" +
		"MATCHING(ID);"
	return z

}
