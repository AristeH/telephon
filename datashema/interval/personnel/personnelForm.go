package personnel

import (
	"Telephon/config"
	"fmt"
	"strconv"
	"strings"

	egui "github.com/alkresin/external"
)
var pTree *egui.Widget
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
	clrLGRAY6   = 0xdddddd
	esMULTILINE = 4
)

//Personnelform форма сотрудника
func Personnelform(wx, wy, ww, wh int, wtitle string) {
	egui.SetDateFormat("DD.MM.YYYY")
	egui.BeginPacket()
	egui.SetImagePath("img/")
	egui.CreateStyle(&(egui.Style{Name: "stPanel", Colors: []int32{clrLGRAY0, clrLGRAY2}, BorderW: 1, BorderClr: clrLGRAY5}))
	pDlg := &egui.Widget{X: wx, Y: wy, W: ww, H: wh, Title: wtitle}
	egui.InitDialog(pDlg)
	pDlg.SetCallBackProc("oninit", buildTree, "buildTree")
	// дерево слева
	pTree = pDlg.AddWidget(&egui.Widget{Type: "tree", Name: "tree",
		X: 5, Y: 5, W: ww / 3, H: wh - 10, Winstyle: egui.WS_VSCROLL,
		Anchor: egui.A_TOPREL + egui.A_BOTTOMREL + egui.A_RIGHTREL + egui.A_LEFTREL,
		AProps: map[string]string{"AImages": egui.ToString("folder.bmp", "folderopen.bmp")}})
	pTree.SetCallBackProc("onsize", nil, "{|o,x,y|o:Move(,,,y-2)}")

	//табличка справа
	pBrw := pDlg.AddWidget(&egui.Widget{Type: "browse", Name: "brw", X: ww/3 + 4, Y: 1, W: ww - ww/3 - 21, H: wh - 41,
		Anchor: egui.A_TOPREL + egui.A_BOTTOMABS + egui.A_RIGHTABS + egui.A_LEFTREL})
	buildBrowse(pBrw)
	pBrw.SetParam("oStyleHead", egui.GetStyle("stPanel"))
	//сплиттер между деревом и табличкой
	pDlg.AddWidget(&egui.Widget{Type: "splitter", X: ww/3 + 1, Y: int((wh - 2) / 2), W: 8, H: 60, Anchor: egui.A_VERTFIX + egui.A_TOPREL + egui.A_BOTTOMREL + egui.A_RIGHTREL + egui.A_LEFTREL,
		AProps: map[string]string{"ALeft": egui.ToString(pTree), "ARight": egui.ToString(pBrw)}})

	pDlg.Activate()
	egui.EndPacket()

}

func buildTree([]string) string{
	db := config.Parametrs.DB
	t := `
	     SELECT ID, PARENT, NAME, WEITHB
         FROM DEPARTMENT  order by WEITHB, PARENT
	     `
	rows, err := db.Query(t)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}

	id := ""
	parent := ""
	name := ""
	level := 0
	for rows.Next() {
		if err := rows.Scan(&id, &parent, &name, &level); err != nil {

			fmt.Println(err)
		}
		if strings.TrimSpace(parent) == "00000000-0000-0000-0000-000000000000" {
			parent = ""
		}
		// if level == 2 {
		// 	break
		// }

		egui.InsertNode(pTree, strings.TrimSpace(parent), strings.TrimSpace(id), strings.TrimSpace(name), "", nil, nil, "")
	}
	return ""
}

func buildBrowse(pBrw *egui.Widget) {
	var arr [][]string
	var cbColor egui.CodeBlock = `private oBrw := Widg("main.brw")
if oBrw != Nil
  if oBrw:nPaintRow%2 == 0
    return {0,15658734,14116608,13421772}
  endif
endif
return {0,16777215,14116608,13421772}`

	pBrw.SetParam("bColorSel", clrLGRAY0)
	pBrw.SetParam("htbColor", clrLGRAY6)
	pBrw.SetParam("tColorSel", 0)
	pBrw.SetParam("httColor", 0)
	pBrw.SetParam("lInFocus", true)

	db := config.Parametrs.DB
	t := `
SELECT
	a.ID,
	b.CODE,	
	b.NAME,
	c.NAME,
	a.DATESTART
FROM
	PERSONEL a,
	PEOPLE b,
	POST c
WHERE
	a.IDPEOPLE = b.ID
	AND c.ID = a.IDPOST
	AND a.EMPLOYMENTTYPE = '1'
	AND a.DATEEND = '3000-01-01' 
ORDER BY b.NAME;
	     `
	rows, _ := db.Query(t)
	defer rows.Close()
	id := 0
	post := ""
	people := ""
	tn := ""
	datestart := ""
	for rows.Next() {
		if err := rows.Scan(&id, &tn, &people, &post, &datestart); err != nil {
			fmt.Println(err)
		}
		arr = append(arr, []string{strconv.Itoa(id), tn, people, post, datestart})
	}

	egui.BrwSetArray(pBrw, &arr)

	egui.BrwSetColumn(pBrw, 1, "ID", 1, 0, false, 0)
	egui.BrwSetColumn(pBrw, 2, "Таб. номер", 1, 0, false, 0)
	egui.BrwSetColumn(pBrw, 3, "Сотрудник", 1, 0, false, 0)
	egui.BrwSetColumn(pBrw, 4, "Должность", 1, 0, false, 0)
	egui.BrwSetColumn(pBrw, 5, "Дата приема", 1, 0, false, 14)

	egui.BrwSetColumnEx(pBrw, 1, "bColorBlock", cbColor)
	egui.BrwSetColumnEx(pBrw, 2, "bColorBlock", cbColor)
	egui.BrwSetColumnEx(pBrw, 3, "bColorBlock", cbColor)
	egui.BrwSetColumnEx(pBrw, 4, "bColorBlock", cbColor)
	egui.BrwSetColumnEx(pBrw, 5, "bColorBlock", cbColor)
}

func fldOnClick(p []string) string {
	return ""
}
