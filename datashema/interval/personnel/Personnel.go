package personnel

//"strconv"
//"fmt"

// Personnel Таблица кадровая история.
type Personnel struct {
	IDpeople       string `json:"ФизическоеЛицо"`
	IDdepartment   string `json:"Подразделение"`
	IDpost         string `json:"Должность"`
	EmploymentType string `json:"ВидЗанятости"`
	DateStart      string `json:"ДатаНачала"`
	DateEnd        string `json:"ДатаОкончания"`
}

//CreateTable Возвращает строку создания таблицы
func (s Personnel) CreateTable() string {
	return `
	CREATE TABLE PERSONEL (
	ID BIGINT NOT NULL,
	IDPEOPLE CHAR(36),
	IDDEPARTMENT CHAR(36),
	IDPOST CHAR(36),
	EMPLOYMENTTYPE INTEGER,
	DATESTART DATE,
	DATEEND DATE,
	CONSTRAINT PERSONEL_PK PRIMARY KEY (ID)
);
	`
}

// Personnels кадровый список
type Personnels struct {
	Personnels []Personnel `json:"Кадры"`
}

//Insert Возвращает строку создания таблицы
func (s Personnel) Insert(pers Personnel) string {
	z := "INSERT INTO PERSONEL(IDpeople, IDdepartment, IDpost,  EmploymentType, DateStart, DateEnd )" +
		"VALUES ('" + pers.IDpeople + "','" +
		pers.IDdepartment + "','" +
		pers.IDpost + "'," +
		pers.EmploymentType +
		", CAST('" + pers.DateStart + "' AS DATE), CAST('" + pers.DateEnd + "' AS DATE));"
	//fmt.Println(z)
	return z

}
