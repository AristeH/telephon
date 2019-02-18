package loadzup

import (
	"Telephon/datashema/reference/people"
	"Telephon/datashema/reference/post"
	"strconv"

	//"time"
	"Telephon/config"
	"Telephon/datashema/interval/personnel"
	"Telephon/datashema/reference/department"
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	//"strings"
)

//Conf общие сведения об обмене
type Conf struct {
	Dn   string `json:"ДатаНачала"`
	Dk   string `json:"ДатаОкончания"`
	Conf string `json:"Конфигурация"`
}

var g Conf
var gen int

func getEmploymentType(file string) string {
	switch file {
	case "Основное место работы":
		return "1"
	case "Внешнее совместительство":
		return "2"
	case "Внутреннее совместительство":
		return "3"
	}
	return "0"
}

//Load загрузка должностей из файла
func loadPost(file string) {
	var posts post.Posts
	db := config.Parametrs.DB
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &posts)
	for i := 0; i < len(posts.Posts); i++ {
		z := post.Post{}.Insert(posts.Posts[i])
		db.Exec(z)
	}
}

//LoadDepartment загрузка подразделений из файла
func loadDepartment(file string) {
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
		z := department.Department{}.Insert(Departments.Departments[i])
		_, err := db.Exec(z)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//Load загрузка физических лиц из файла
func loadPeople(file string) {
	var sotrs people.Sotrs
	db := config.Parametrs.DB
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &sotrs)
	for i := 0; i < len(sotrs.Sotrs); i++ {
		z := people.People{}.Insert(sotrs.Sotrs[i])
		db.Exec(z)
	}
}

func loadpersonnelEmployment(file string) {
	var Personnels personnel.Personnels
	db := config.Parametrs.DB

	z := "DELETE FROM PERSONEL WHERE  DateStart >= '" + g.Dn + "'  and  DateStart <=  '" + g.Dk + "';"
	_, err := db.Exec(z)
	if err != nil {
		fmt.Println(err)
	}

	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Personnels)
	for i := 0; i < len(Personnels.Personnels); i++ {
		pers := Personnels.Personnels[i]
		pers.EmploymentType = getEmploymentType(pers.EmploymentType)
		gen = gen + 1
		z = "INSERT INTO PERSONEL(ID, IDpeople, IDdepartment, IDpost,  EmploymentType, DateStart, DateEnd )" +
			"VALUES (" + strconv.Itoa(gen) + ",'" + pers.IDpeople + "','" +
			pers.IDdepartment + "','" +
			pers.IDpost + "'," +
			pers.EmploymentType +
			", CAST('" + pers.DateStart + "' AS DATE), CAST('" + pers.DateEnd + "' AS DATE));"
		_, err := db.Exec(z)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func loadpersonnelMovement(file string) {
	var Personnels personnel.Personnels
	db := config.Parametrs.DB
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Personnels)
	for i := 0; i < len(Personnels.Personnels); i++ {
		pers := Personnels.Personnels[i]
		pers.EmploymentType = getEmploymentType(pers.EmploymentType)
		z := "UPDATE PERSONEL " +
			"SET  DATEEND ='" + pers.DateStart +
			"' WHERE  IDPEOPLE = '" + pers.IDpeople +
			"' and EMPLOYMENTTYPE = '" + pers.EmploymentType +
			"' and DATEEND = '3000-01-01'  AND DATESTART = (SELECT  MAX(DATESTART) " +
			"FROM PERSONEL" +
			" WHERE DATEEND = '3000-01-01' and IDPEOPLE = '" + pers.IDpeople + "' and DATESTART < '" + pers.DateStart +
			"' GROUP BY IDPEOPLE);"
		_, err := db.Exec(z)
		//fmt.Println(z)
		if err != nil {
			fmt.Println(err)
		}

		gen = gen + 1
		z = "INSERT INTO PERSONEL(ID, IDpeople, IDdepartment, IDpost,  EmploymentType, DateStart, DateEnd )" +
			"VALUES (" + strconv.Itoa(gen) + ",'" + pers.IDpeople + "','" +
			pers.IDdepartment + "','" +
			pers.IDpost + "'," +
			pers.EmploymentType +
			", CAST('" + pers.DateStart + "' AS DATE), CAST('" + pers.DateEnd + "' AS DATE));"
		//fmt.Println(z)
		_, err = db.Exec(z)
		if err != nil {
			fmt.Println(err)
		}

	}
}
func loadpersonnelDismissal(file string) {
	var Personnels personnel.Personnels
	db := config.Parametrs.DB
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Personnels)
	for i := 0; i < len(Personnels.Personnels); i++ {
		pers := Personnels.Personnels[i]
		pers.EmploymentType = getEmploymentType(pers.EmploymentType)
		z := "UPDATE PERSONEL " +
			"SET  DATEEND ='" + pers.DateEnd +
			"' WHERE  IDPEOPLE = '" + pers.IDpeople +
			"' and EMPLOYMENTTYPE = '" + pers.EmploymentType +
			"' and DATEEND = '3000-01-01'  AND DATESTART = (SELECT  MAX(DATESTART) " +
			"FROM PERSONEL" +
			" WHERE DATEEND = '3000-01-01' and IDPEOPLE = '" + pers.IDpeople + "' and DATESTART < '" + pers.DateEnd +
			"' GROUP BY IDPEOPLE);"
		_, err := db.Exec(z)
		//fmt.Println(z)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		defer fileReader.Close()
		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()
		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}
	return nil
}

func loadconf(file string) {
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &g)
}

//LoadZUP загрузка сведения ЗУП из файла
func LoadZUP() {
	path := config.Parametrs.Path
	unzip(path+"нн.zip", path+"zup/")
	loadconf(path + "zup/Конфигурация.json")
	config.Parametrs.DB.QueryRow("SELECT max(ID) FROM PERSONEL").Scan(&gen)
	filepath.Walk(path+"zup/", func(path1 string, info os.FileInfo, err error) error {
		dir, file := filepath.Split(path1)
		fmt.Println("Dir:", dir)
		fmt.Println("File:", file)
		switch file {
		case "1_ДолжностиОрганизаций.json":
			//loadPost(path1)
		case "2_ПодразделенияОрганизаций.json":
			//loadDepartment(path1)
		case "3_ДополнительныеНачисленияОрганизаций.json":
			//loadDepartment(path1)
		case "4_ОсновныеНачисленияОрганизаций.json":
			//loadDepartment(path1)
		case "5_УдержанияОрганизаций.json":
			//loadDepartment(path1)
		case "6_ФизическиеЛица.json":
			//loadPeople(path1)
		case "7_ПриемНаРаботуВОрганизацию.json":
			//loadpersonnelEmployment(path1)
		case "8_КадровоеПеремещениеОрганизаций.json":
			//loadpersonnelMovement(path1)
		case "9_УвольнениеИзОрганизаций.json":
			loadpersonnelDismissal(path1)
		}
		return nil
	})

}
