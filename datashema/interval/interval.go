package interval

import (
	"Telephon/datashema/interval/personnel"
	"Telephon/config"

	"fmt"
	"strings"
)




// InitInterval создание не существующих справочников
func InitInterval() {
	db := config.Parametrs.DB
	elements := map[string]string{
		"PERSONEL":     personnel.Personnel{}.CreateTable(),

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
