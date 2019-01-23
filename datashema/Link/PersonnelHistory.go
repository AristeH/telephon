package link

import (
	"io/ioutil"

	"fmt"
	"encoding/json"
	"os"
	
)

// PersonnelHistory Таблица кадровая история.
type PersonnelHistory struct {
	IDpeople       string `IDpeople:"idpeople"`
	IDdepartment   string `IDdepartment:"iddepartment"`
	IDpost         string `IDpost:"idpost"`
	DateDoc     string `DateDoc:"iddepartment"`
	DateStart     string `DateStart:"DateStart"`
	DateEnd string `DateEnd:"DateEnd"`
}

//CreateTable Возвращает строку создания таблицы
func (s PersonnelHistory) CreateTable() string {
	return `
	    CREATE TABLE PERSONELLHISTORY(
			IDpeople CHAR(36),  
			IDdepartment CHAR(36),
			IDpost CHAR(36),
			DateDoc DATE,
			DateStart DATE,
			DateEnd DATE);
			`
}
//Insert Возвращает строку создания таблицы
func (s PersonnelHistory) Insert(rec PersonnelHistory ) string {
	return `
	    INSERT INTO  TABLE PERSONELLHISTORY(
			IDpeople CHAR(36),  
			IDdepartment CHAR(36),
			IDpost CHAR(36),
			DateDoc DATE,
			DateStart DATE,
			DateEnd DATE);
			`
} 

//Load Возвращает строку создания таблицы
func (s PersonnelHistory) Load(file string)  {
var f PersonnelHistory
    jsonFile, err := os.Open(file)
    defer jsonFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
	  byteValue, _ := ioutil.ReadAll(jsonFile)
	  json.Unmarshal(byteValue, &f)



} 
