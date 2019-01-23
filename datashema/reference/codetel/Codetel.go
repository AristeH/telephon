package codetel

// Codetel Таблица КодаТелефонов. Кода телефонов городов и сотовых операторов.
type Codetel struct {
	ID         string
	Code       string
	IDoperator string
	S          string
	Po         string
	IDregion   string
}

// CreateTable Возвращает строку создания таблицы
func (s Codetel) CreateTable() string {
	return "CREATE TABLE" +
	    		"Codetel(ID CHAR(36)," +
				"Name CHAR(50)," + 
			   " Code CHAR(50)," + 
				"IDoperator CHAR(36)," +
			   " IDregion CHAR(36), " +
			    "S CHAR(9)," +
			   " Po CHAR(9));"
}