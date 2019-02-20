package main

import (
//	"Telephon/obrabotki/loadzup"
	"Telephon/datashema/interval"
	"Telephon/config"
	"Telephon/forms"
	_ "github.com/nakagami/firebirdsql"
	ref "Telephon/datashema/reference"
	"fmt"
	egui "github.com/alkresin/external"
)

func main() {
	var n int
	config.SetParametrs()

	ref.InitRef()
	interval.InitInterval()

	config.Parametrs.DB.QueryRow("SELECT Count(*) FROM rdb$relations").Scan(&n)

	fmt.Println("Relations count=", n)
	if egui.Init(config.Parametrs.GuiInit) != 0 {
		return
	}
	//loadzup.LoadZUP()
	forms.MainForm(100, 100, 800, 700, "Телефоны")
	egui.Exit()
}