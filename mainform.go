package main

import egui "github.com/alkresin/external"

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
var hp = 40

func initStyle() {
	//стили используемые в окне
	egui.CreateStyle(&(egui.Style{Name: "stSplitter", Colors: []int32{clrBLUE}, BorderW: 2, BorderClr: clrBLACK}))
	egui.CreateStyle(&(egui.Style{Name: "stTop", Colors: []int32{clrLGRAY0, clrLGRAY5}, Orient: 5}))
	egui.CreateStyle(&(egui.Style{Name: "stPanel", Colors: []int32{clrLGRAY0, clrLGRAY2}, BorderW: 1, BorderClr: clrLGRAY5}))
	egui.CreateStyle(&(egui.Style{Name: "stBot", Colors: []int32{clrLGRAY1, clrLGRAY5}, Orient: 1, BorderW: 2, BorderClr: clrLGRAY5}))
}

/* func initPanelLeft(pWindow *egui.Widget, wx, wy, ww, wh int) *egui.Widget {
	panelleft := pWindow.AddWidget(&egui.Widget{Type: "panel", Name: "lpan", X: wx, Y: wy, W: ww, H: wh,
		Anchor: egui.A_TOPABS + egui.A_BOTTOMABS})
	panelleft.AddWidget(&egui.Widget{Type: "paneltop", H: hp,
	    Anchor: egui.A_TOPABS + egui.A_BOTTOMABS,
		AProps: map[string]string{"HStyle": "stTop"}})
	// Tree 
	pTree := panelleft.AddWidget(&egui.Widget{Type: "tree", Name: "tree",
		X: 1, Y: hp + 1, W: ww - 2, H: wh - hp - 1, Winstyle: egui.WS_VSCROLL,
		Anchor: egui.A_TOPREL + egui.A_BOTTOMREL + egui.A_RIGHTREL + egui.A_LEFTREL,
		AProps: map[string]string{"AImages": egui.ToString("folder.bmp", "folderopen.bmp")}})
	buildTree(pTree)
	return panelleft
}
 */
/* func initPanelRight(pWindow *egui.Widget, wx, wy, ww, wh int) *egui.Widget {
	var arr = [][]string{{"Наименование", "Name", "строка,250", "ФИО"}, {"Код", "Cod", "строка,50", "Код"}, {"ID", "ID", "строка,36", "GUID"}}
	panelright := pWindow.AddWidget(&egui.Widget{Type: "panel", Name: "rpan", X: wx, Y: wy, W: ww, H: wh,
		Anchor: egui.A_TOPABS + egui.A_BOTTOMABS + egui.A_RIGHTABS + egui.A_LEFTABS})

	panelright.AddWidget(&egui.Widget{Type: "paneltop", H: hp * 4,
		Anchor: egui.A_TOPABS + egui.A_BOTTOMABS,
		AProps: map[string]string{"HStyle": "stTop"}})

	pBrw := panelright.AddWidget(&egui.Widget{Type: "browse", Name: "brw", X: 1, Y: hp*4 + 1, W: ww - 10, H: wh - hp*8,
		Anchor: egui.A_TOPABS + egui.A_BOTTOMABS + egui.A_RIGHTABS + egui.A_LEFTABS})
	pBrw.SetParam("oStyleHead", egui.GetStyle("stPanel"))
	egui.BrwSetArray(pBrw, &arr)
	egui.BrwSetColumn(pBrw, 1, "Реквизит", 1, 0, false, 0)
	egui.BrwSetColumn(pBrw, 2, "Name", 1, 0, false, 0)
	egui.BrwSetColumn(pBrw, 3, "Тип", 1, 0, false, 20)
	egui.BrwSetColumn(pBrw, 4, "Описание", 1, 0, true, 0)

	// panelbot
	panelbot := panelright.AddWidget(&egui.Widget{Type: "panelbot", H: hp * 4})

	panelbot.AddWidget(&egui.Widget{Type: "combo", Name: "combLang", X: 3, Y: 1, W: 80, H: 24,
		AProps: map[string]string{"AItems": egui.ToString("Go")}})
	panelbot.AddWidget(&egui.Widget{Type: "cedit", X: 2, Y: 25, W: int(ww/2) - 1, H: 80,
	    Anchor: egui.A_TOPREL + egui.A_BOTTOMREL + egui.A_RIGHTREL + egui.A_LEFTREL})

	panelbot.AddWidget(&egui.Widget{Type: "combo", Name: "SQLLang", X: int(ww/2)+2, Y: 1, W: 80, H: 24,
	Anchor: egui.A_TOPREL + egui.A_BOTTOMREL + egui.A_RIGHTREL + egui.A_LEFTREL,
		AProps: map[string]string{"AItems": egui.ToString("sqlite", "firebird")}})

	panelbot.AddWidget(&egui.Widget{Type: "cedit", X: int(ww/2)+2, Y: 25, W: int(ww / 2)-10, H: 80,
	Anchor: egui.A_TOPREL + egui.A_BOTTOMREL + egui.A_RIGHTREL + egui.A_LEFTREL})
	return panelright
}
 */
func mainform(wx, wy, ww, wh int, wtitle string) {
	initStyle()
	pWindow := &egui.Widget{X: wx, Y: wy, W: ww, H: wh, Title: wtitle, AProps: map[string]string{"Icon": "main"}}
	egui.InitMainWindow(pWindow)
	pWindow.AddWidget(&egui.Widget{Type: "paneltop", H: hp, AProps: map[string]string{"HStyle": "stTop"}})
	// panelleft := initPanelLeft(pWindow, 0, hp, plw, ph)
	// panelright := initPanelRight(pWindow, plw+2, hp+2, ww-plw-4, ph)
	//pWindow.AddWidget(&egui.Widget{Type: "splitter", X: plw, Y: int((wh - plw - 2) / 2), W: 8, H: hp, Anchor: egui.A_VERTFIX,
		//AProps: map[string]string{"ALeft": egui.ToString(panelleft), "ARight": egui.ToString(panelright)}})
	pWindow.AddWidget(&egui.Widget{Type: "panelbot", H: hp, AProps: map[string]string{"HStyle": "stBot"}})
	pWindow.Activate()
	egui.Exit()
}

func buildTree(pTree *egui.Widget) {
	pTree.SetCallBackProc("onsize", nil, "{|o,x,y|o:Move(,,,y-72)}")
	egui.InsertNode(pTree, "", "n1", "Справочники", "", nil, nil, "")
	egui.InsertNode(pTree, "n1", "n1a", "Сотрудники", "", []string{"book.bmp"}, nil, "hwg_msginfo(\"Сотрудники\")")
	egui.InsertNode(pTree, "n1", "n2b", "Должности", "", []string{"book.bmp"}, nil, "hwg_msginfo(\"Должности\")")
	egui.InsertNode(pTree, "", "n2", "Документы", "", nil, nil, "")
	egui.InsertNode(pTree, "n2", "n2a", "Прием", "", nil, nil, "")
	egui.InsertNode(pTree, "", "n3", "Сведения", "", nil, nil, "")
	egui.InsertNode(pTree, "n3", "n3", "Документы", "", nil, nil, "")
}

func fldOnClick(p []string) string {
	return ""
}
