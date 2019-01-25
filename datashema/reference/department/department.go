package department

import (
	"Telephon/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Department структура Подразделения.
type Department struct {
	ID     string `json:"УникальныйИдентификатор"`
	Name   string `json:"Наименование"`
	Code   string `json:"Код"`
	Parent string `json:"Родитель"`
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
			Name CHAR(150));
		`
}

// Insert добавляет запись в таблицу
func (s Department) Insert(р Department) string {
	z := "UPDATE OR INSERT INTO Department(ID, Parent, Code, Name)" +
		"VALUES ('" + р.ID + "','" + р.Parent + "','" + р.Code + "','" + р.Name + "')" +
		"MATCHING(ID);"
	return z

}

//Load загрузка должностей из файла
// структура файла:
// "{
// "Подразделения": [
// 		{
// 			"Наименование": "Общее",
// 			"Код": "000000016",
// 			"УникальныйИдентификатор": "c1feb6af-477c-11e3-b5cf-00145edc4f8d",
// 			"Родитель": "00000000-0000-0000-0000-000000000000"
// 		},
// 		{
// 			"Наименование": "ОП \"Белогорск\"",
// 			"Код": "031400226",
// 			"УникальныйИдентификатор": "d2331076-032f-11e7-89ef-1c98ec28debf",
// 			"Родитель": "c1feb6af-477c-11e3-b5cf-00145edc4f8d"
// 		}
// ]
// }"
func (s Department) Load(file string) {
	var Departments Departments

	var t Node 
	var m map[string]Node

	db := config.Parametrs.DB
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Departments)
	for i := 0; i < len(Departments.Departments); i++ {
		z := Department{}.Insert(Departments.Departments[i])
		db.Exec(z)
		t.addNode(Departments.Departments[i], m)
	}
}

//Node  дерево
type Node struct {
	ID    string
	Name  string
	ParentID   string
	Children  []*Node
}


func (f *Node) addNode(dep Department) {
	f.Children = append(f.Children, &Node{ID : dep.ID,Name: dep.Name, ParentID:dep.Parent})
}