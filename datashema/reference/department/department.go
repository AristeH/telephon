package department

import (
	"strconv"
)

// Department структура Подразделения.
type Department struct {
	ID     string `json:"УникальныйИдентификатор"`
	Name   string `json:"Наименование"`
	Code   string `json:"Код"`
	Parent string `json:"Родитель"`
	WeithB int `json:"Уровень"`
}

// Departments список подразделений
type Departments struct {
	Departments []Department `json:"2_ПодразделенияОрганизаций"`
}

// CreateTable Возвращает строку создания таблицы
func (s Department) CreateTable() string {
	return `
CREATE TABLE DEPARTMENT (
	ID CHAR(36) NOT NULL,
	PARENT CHAR(36),
	CODE CHAR(10),
	NAME CHAR(150),
	WEITHB INTEGER,
	CONSTRAINT DEPARTMENT_PK PRIMARY KEY (ID)
);
`
}
// CreateTable Возвращает строку создания таблицы
func (s Department) CreateIndex() string {
	return `
CREATE INDEX DEPARTMENT_PARENT_IDX ON DEPARTMENT (PARENT,CODE,WEITHB,NAME);
);
`
}
// Insert добавляет запись в таблицу
func (s Department) Insert(р Department) string {
	z := "UPDATE OR INSERT INTO Department(ID, Parent, Code, Name, WeithB)" +
		"VALUES ('" + р.ID + "','" + р.Parent + "','" + р.Code + "','" + р.Name + "','" + strconv.Itoa(р.WeithB) + "')" +
		"MATCHING(ID);"
	return z

}
