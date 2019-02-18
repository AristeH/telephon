package link

// import (

// 	"time"
// 	d "database/sql"
// 	"fmt"
// 	"strings"


// )

// type link struct {
// 	DateEnd time.Time

// 	Name string
// }

// func (s link) CreateTable(name string) string {
// 	return "CREATE TABLE " + name +
// 		"(ID CHAR(36) , Name CHAR(50));"
// }

// func (s link) GetFields(name string) string {
// 	return "'DateEnd','DateStart'"
// }
// // InitLink создание не существующих справочников 
// func InitLink(db *d.DB) {
// 	elements := map[string]string{
// 		"PERSONEL":   Personnel{}.CreateTable(),
// 	}
// 	var (
// 		name string
// 	)

// 	rows, err := db.Query("SELECT  RDB$RELATION_NAME  FROM RDB$RELATIONS  WHERE RDB$RUNTIME  is not NULL;")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for rows.Next() {
// 		if err := rows.Scan(&name); err != nil {
// 			fmt.Println(err)
// 		}
// 		_, ok := elements[strings.TrimSpace(name)]
// 		if ok {
// 			delete(elements, strings.TrimSpace(name))
// 		}

// 	}
// 	for country := range elements {
// 		fmt.Println("Создана таблица ", country, "is", elements[country])
// 		_, err = db.Exec(elements[country])
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	}


// }