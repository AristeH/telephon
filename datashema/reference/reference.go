package reference

import (
	"Telephon/config"
	"Telephon/datashema/reference/codetel"
	"Telephon/datashema/reference/department"
	"Telephon/datashema/reference/hierarchy"
	"Telephon/datashema/reference/oblast"
	"Telephon/datashema/reference/operator"
	"Telephon/datashema/reference/people"
	"Telephon/datashema/reference/post"

	"fmt"
	"strings"
)

type reference struct {
	ID   string `json:"УникальныйИдентификатор"`
	Cod  string `Cod:"Код"`
	Name string `json:"Наименование"`
}

func (s reference) CreateTable(name string) string {
	return "CREATE TABLE " + name +
		"(ID CHAR(36) , Name CHAR(50));"
}

func (s reference) GetFields(name string) string {
	return "'ID','Name'"
}

// InitRef создание не существующих справочников
func InitRef() {
	db := config.Parametrs.DB
	err1 := db.Ping()
	if err1 != nil {
		fmt.Println(err1)
	}
	elements := map[string]string{
		"OBLAST":     oblast.Oblast{}.CreateTable(),
		"PEOPLE":     people.People{}.CreateTable(),
		"DEPARTMENT": department.Department{}.CreateTable(),
		"POST":       post.Post{}.CreateTable(),
		"HIERARCHY":  hierarchy.Hierarchy{}.CreateTable(),
		"CODETEL":    codetel.Codetel{}.CreateTable(),
		"OPERATOR":   operator.Operator{}.CreateTable(),
	}
	var (
		name string
	)

	rows, err := db.Query("SELECT  RDB$RELATION_NAME  FROM RDB$RELATIONS where (RDB$SYSTEM_FLAG = 0) AND (RDB$RELATION_TYPE = 0)	order by RDB$RELATION_NAME;")
	n := 0
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			n = n + 1
			if err := rows.Scan(&name); err != nil {
				fmt.Println(err)
			}
			_, ok := elements[strings.TrimSpace(name)]
			if ok {
				delete(elements, strings.TrimSpace(name))
			}

		}
	}
	fmt.Println(n)
	for country := range elements {
		fmt.Println("Создаем таблицу ", country, " is ", elements[country])
		_, err = db.Exec(elements[country])
		if err != nil {
			fmt.Println(err)
		}
	}

}
