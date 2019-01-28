package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	d "database/sql"
)

// User структура люди.
type User struct {
	ID   string `json:"УникальныйИдентификатор"`
	Code string `json:"Код"`
	Name string `json:"Наименование"`
	Password string `json:"Пароль"`
}

// Users список сотрудников
type Users struct {
	users []User `json:"Пользователи"`
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
 

//Load загрузка физических лиц из файла
// структура файла:
// "{
//	Сотрудники": [
// 		{
// 			"Наименование": "Иванов иван иванович",
// 			"Код": "ЗЦК0000617",
// 			"УникальныйИдентификатор": "a54d0158-444d-11e7-80d1-1402ec43021b",
// 			"Родитель": "00000000-0000-0000-0000-000000000000"
// 		},
// 		{
// 			"Наименование": "Петров Петр Петрович",
// 			"Код": "ЗЦК0000991",
// 			"УникальныйИдентификатор": "c3c06ea7-66e0-11e7-8a50-1402ec43021b",
// 			"Родитель": "00000000-0000-0000-0000-000000000000"
// 		}
//  ]
// }"
func (s User) Load(file string,  db *d.DB ) {
	var sotrs Sotrs
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &sotrs)
	for i := 0; i < len(sotrs.Sotrs); i++ {
		z := User{}.Insert(sotrs.Sotrs[i])
		db.Exec(z)
	}
}
