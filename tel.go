package main

import (
	"Telephon/config"
	//link "Telephon/datashema/Link"
	ref "Telephon/datashema/reference"

	"database/sql"
	"fmt"

	_ "github.com/nakagami/firebirdsql"
)

func main() {
	var n int

	db, _ := sql.Open("firebirdsql", "SYSDBA:masterkey@127.0.0.1:3050/d:/fr/FIRST.FDB")
	config.Parametrs.DB = db
	ref.InitRef()
	//link.InitLink()
	defer db.Close()

	// Сведения
	//db.Exec(link.PersonnelHistory{}.CreateTable())

	config.Parametrs.DB.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)
	fmt.Println("Relations count=", n)
//	_, err := db.Exec("insert into Post(id, name) values ('12','Основное')")
//	fmt.Println(err)
//	ref.People{}.Load("//FIRST/First-0бщая/1c/Сергей Щ/обработки зуп30/uu/sotr.json", db)
//	ref.Post{}.Load("//FIRST/First-0бщая/1c/Сергей Щ/обработки зуп30/uu/dolg.json", db)
//	ref.Division{}.Load("//FIRST/First-0бщая/1c/Сергей Щ/обработки зуп30/uu/podr.json", db)
}
