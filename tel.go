package main

import (
	"Telephon/config"

	_ "github.com/nakagami/firebirdsql"

	//link "Telephon/datashema/Link"
	ref "Telephon/datashema/reference"
	"fmt"

	egui "github.com/alkresin/external"
)

const (
	clrBLACK    = 0
	clrWHITE    = 0xffffff
	clrBLUE     = 0xff0000
	clrGREEN    = 32768
	clrLGRAY0   = 0xeeeeee
	clrLGRAY1   = 0xbbbbbb
	clrLGRAY2   = 0x999999
	clrLGRAY4   = 0x666666
	clrLGRAY5   = 0x444444
	esMULTILINE = 4
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
	//	ref.Division{}.Load("//FIRST/First-0бщая/1c/Сергей Щ/обработки зуп30/uu/podr.json", db)

	if egui.Init(app.GuiInit) != 0 {
		return
	}
	pWindow := &egui.Widget{X: 200, Y: 150, W: -800, H: -600, Title: "Телефоны",
		Font: app.PFontMain, AProps: map[string]string{"Icon": "etutor"}}
	egui.InitMainWindow(pWindow)
	pWindow.Activate()
	egui.Exit()
}
