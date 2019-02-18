package loadzup

import (
	// "Telephon/config"
	// "fmt"
	// "strconv"
	// "strings"

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
	clrLGRAY6   = 0xdddddd
	esMULTILINE = 4
)

//OpenForm форма сотрудника
func OpenForm(wx, wy, ww, wh int, wtitle string) {
	egui.BeginPacket()
	egui.SetImagePath("img/")
	egui.CreateStyle(&(egui.Style{Name: "stPanel", Colors: []int32{clrLGRAY0, clrLGRAY2}, BorderW: 1, BorderClr: clrLGRAY6}))
	pDlg := &egui.Widget{X: wx, Y: wy, W: ww, H: wh, Title: wtitle}
	egui.InitDialog(pDlg)
	pDlg.AddWidget(&egui.Widget{Type: "label", X: 20, Y: 10, W: 160, H: 24, Title: "Выберите файл:"})
	pDlg.AddWidget(&egui.Widget{Type: "edit", Name: "edi1", X: 20, Y: 32, W: 160, H: 24,
		AProps: map[string]string{"Picture": "@!R /XXX:XXX/"}})
	pDlg.Activate()
	egui.EndPacket()

}
