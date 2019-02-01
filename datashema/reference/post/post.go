package post
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	d "database/sql"
)
// Post структура должности.
type Post struct {
	ID   string `json:"УникальныйИдентификатор"`
	Name string `json:"Наименование"`
}

// Posts список должностей
type Posts struct {
	Posts []Post `json:"Должности"`
}

// CreateTable Возвращает строку создания таблицы
func (s Post) CreateTable() string {
	return `
	CREATE TABLE POST (
		ID CHAR(36),
		NAME CHAR(100)
	);
		`
}

// Insert добавляет запись в таблицу
func (s Post) Insert(р Post) string {
	z := "UPDATE OR INSERT INTO Post(ID, Name)"+
		"VALUES ('" + р.ID + "','" + р.Name  +"')" +
		"MATCHING(ID);"
	return z
		
} 

//Load загрузка должностей из файла
// структура файла:
// "{
//	"Должности": [
// 		{
// 			"Наименование": ""Дорожный рабочий 3 разряд",
// 			"УникальныйИдентификатор": "a54d0158-444d-11e7-80d1-1402ec43021b"
// 		},
// 		{
// 			"Наименование": ""Дорожный рабочий 2 разряд",
// 			"УникальныйИдентификатор": "c3c06ea7-66e0-11e7-8a50-1402ec43021b"
// 		}
//  ]
// }"
func (s Post) Load(file string,  db *d.DB ) {
	var posts Posts
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &posts)
	for i := 0; i < len(posts.Posts); i++ {
		z := Post{}.Insert(posts.Posts[i])
		db.Exec(z)
	}
}