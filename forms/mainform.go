package forms

import (
	"Telephon/obrabotki/loadzup"
	"Telephon/datashema/interval/personnel"

	egui "github.com/alkresin/external"
)

//MainForm основная форма приложения
func MainForm(wx, wy, ww, wh int, wtitle string) {
	egui.SetImagePath("img/")

	initStyle()
	pWindow := &egui.Widget{X: wx, Y: wy, W: ww, H: wh, Title: wtitle}
	egui.InitMainWindow(pWindow)
	paneltop := pWindow.AddWidget(&egui.Widget{Type: "paneltop", H: hp, AProps: map[string]string{"HStyle": "stTop"}})

	paneltop.AddWidget(&egui.Widget{Type: "button", X: 2, Y: 2, W: 100, H: 32, Title: "Сотрудники"})
	egui.PLastWidget.SetCallBackProc("onclick", fsett3, "fsett3")

	lpan := pWindow.AddWidget(&egui.Widget{Type: "panel", Name: "lpan", X: 0, Y: hp, W: hp, H: wh - hp,
		Anchor: egui.A_HORFIX + egui.A_BOTTOMREL + egui.A_TOPABS + egui.A_VERTFIX + egui.A_LEFTABS,
		AProps: map[string]string{"HStyle": "stTop"}})

	lpan.AddWidget(&egui.Widget{Type: "button", X: 2, Y: 2, W: 36, H: 32, Title: "ЗУП"})
	egui.PLastWidget.SetCallBackProc("onclick", obrabotkaLoadZUP, "obrabotkaLoadZUP")


	pWindow.Activate()
	egui.Exit()

}

func fsett3(p []string) string {
	personnel.Sotrform(100, 100, 800, 700, "Сотрудники")
	return ""
}

func obrabotkaLoadZUP(p []string) string {
	loadzup.OpenForm(100, 100, 400, 350, "Сотрудники")
	return ""
}
