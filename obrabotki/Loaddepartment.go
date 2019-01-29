package obrabotki

import (
	"Telephon/config"
	"Telephon/datashema/reference/department"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

//LoadDepartment загрузка должностей из файла
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
func LoadDepartment(file string) {
	var Departments department.Departments
	db := config.Parametrs.DB
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Departments)
	for i := 0; i < len(Departments.Departments); i++ {
		Departments.Departments[i].WeithB = strconv.Itoa(i)
		z := department.Department{}.Insert(Departments.Departments[i])
		_,err := db.Exec(z)
		if err != nil {
			fmt.Println(err)
		}
	}
}
