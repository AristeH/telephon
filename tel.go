package main

import (
	//"Telephon/obrabotki"

	"Telephon/config"

	_ "github.com/nakagami/firebirdsql"

	//link "Telephon/datashema/Link"
	ref "Telephon/datashema/reference"
	"fmt"

	egui "github.com/alkresin/external"
)



func main() {
	var n int

	config.SetParametrs()
	app := config.Parametrs
	ref.InitRef()
	//link.InitLink()
	defer config.Parametrs.DB.Close()

	// Сведения
	//db.Exec(link.PersonnelHistory{}.CreateTable())

	app.DB.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)
	fmt.Println("Relations count=", n)
	//	_, err := db.Exec("insert into Post(id, name) values ('12','Основное')")
	//	fmt.Println(err)
	//	ref.People{}.Load("//FIRST/First-0бщая/1c/Сергей Щ/обработки зуп30/uu/sotr.json", db)
	//	ref.Post{}.Load("//FIRST/First-0бщая/1c/Сергей Щ/обработки зуп30/uu/dolg.json", db)
  // obrabotki.LoadDepartment("//FIRST/First-0бщая/1c/Сергей Щ/обработки зуп30/uu/podr.json")


	if egui.Init(app.GuiInit) != 0 {
		return
	}
	egui.SetImagePath("images/")
	mainform(100, 100, 800, 700, "Сотрудники")
}
