package md

// Columns определяет названия полей в таблицах, их типы данных и размер,
// связанные таблицы (для полей ссылочного типа), значения полей по умолчанию,
// флаги обязательности, уникальности и т.п.
// IDtable    - ссылка на таблицу
// Name       - название столбца в бд
// Note       - синоним
// Default    - значение по умолчанию
// Obligatory - обязательность
// Unique     - уникальность
type Columns struct {
	ID         string
	IDtable    Tables
	Name       string
	Note       string
	DataType   DataTypes
	Default    string
	Obligatory bool
	Unique     bool
}

func initcol(idtables string){


}