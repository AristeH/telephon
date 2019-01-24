package main

import egui "github.com/alkresin/external"
var hp = 40


func initStyle() {
	//стили используемые в окне
	egui.CreateStyle(&(egui.Style{Name: "stSplitter", Colors: []int32{clrBLUE}, BorderW: 2, BorderClr: clrBLACK}))
	egui.CreateStyle(&(egui.Style{Name: "stTop", Colors: []int32{clrLGRAY0, clrLGRAY5}, Orient: 5}))
	egui.CreateStyle(&(egui.Style{Name: "stPanel", Colors: []int32{clrLGRAY0, clrLGRAY2}, BorderW: 1, BorderClr: clrLGRAY5}))
	egui.CreateStyle(&(egui.Style{Name: "stBot", Colors: []int32{clrLGRAY1, clrLGRAY5}, Orient: 1, BorderW: 2, BorderClr: clrLGRAY5}))
}

func mainform(wx, wy, ww, wh int, wtitle string) {

	initStyle()
	pWindow := &egui.Widget{X: wx, Y: wy, W: ww, H: wh, Title: wtitle, AProps: map[string]string{"Icon": "main"}}
	egui.InitMainWindow(pWindow)


	pWindow.AddWidget(&egui.Widget{Type: "panelbot", H: hp, AProps: map[string]string{"HStyle": "stBot"}})
	pWindow.Activate()
	egui.Exit()
}