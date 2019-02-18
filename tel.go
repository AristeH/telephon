package main

import (
	//"Telephon/datashema/interval"
	"Telephon/config"
	"Telephon/forms"
	_ "github.com/nakagami/firebirdsql"
	//ref "Telephon/datashema/reference"
	"fmt"
	egui "github.com/alkresin/external"
)

func main() {
	var n int
	config.SetParametrs()
	app := config.Parametrs
	//ref.InitRef()
	//interval.InitInterval()
	defer config.Parametrs.DB.Close()
	app.DB.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)
	fmt.Println("Relations count=", n)
	if egui.Init(app.GuiInit) != 0 {
		return
	}
	forms.MainForm(100, 100, 800, 700, "Телефоны")
	egui.Exit()
}