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
	return `
	CREATE TABLE CODETEL (
		ID CHAR(36),
		NAME CHAR(50),
		CODE CHAR(50),
		IDOPERATOR CHAR(36),
		IDREGION CHAR(36),
		S CHAR(9),
		PO CHAR(9),
		CONSTRAINT CODETEL_PK PRIMARY KEY (ID)
	);
	`
}
