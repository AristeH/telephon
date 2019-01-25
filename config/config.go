package config

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	egui "github.com/alkresin/external"
)

// Config  Таблица конфигурации. Необходимые настройки для приложения.
type Config struct {
	DB                              *sql.DB
	PFontMain, PFontMenu, PFontText *egui.Font
	GuiInit                         string
}

// Parametrs  создание подключения к БД
var Parametrs Config

// InitDB  создание подключения к БД
func setDB(db *sql.DB) {
	Parametrs.DB = db
}

func isFileExists(sPath string) bool {
	if _, err := os.Stat(sPath); os.IsNotExist(err) {
		return false
	}
	return true
}

// SetParametrs установка параметров приложения
func SetParametrs() {

	type XFont struct {
		Family string `xml:"family,attr"`
		Height int    `xml:"height,attr"`
	}
	type DB struct {
		Path string `xml:"path,attr"`
		Name string `xml:"name,attr"`
	}
	type Ini struct {
		Guiserver string `xml:"guiserver"`
		Ipaddr    string `xml:"ip"`
		Port      int    `xml:"port"`
		Log       int    `xml:"log"`
		FontMain  XFont  `xml:"fontmain"`
		FontMenu  XFont  `xml:"fontmenu"`
		FontText  XFont  `xml:"fonttext"`
		DB        DB     `xml:"db"`
	}

	var pIni = &Ini{}

	// Check, is a current directory the same, where etutor's files are placed.
	// If no, try to change it to that one where executable is.
	sCurrDir, _ := os.Getwd()
	if !isFileExists(sCurrDir + "/app.ini") {
		ex, _ := os.Executable()
		sCurrDir := filepath.Dir(ex)
		if isFileExists(sCurrDir + "/app.ini") {
			os.Chdir(sCurrDir)
		}
	}

	getxml("app.ini", pIni)

	if pIni.Guiserver != "" {
		Parametrs.GuiInit = "guiserver=" + pIni.Guiserver + "\n"
	}
	if pIni.Port != 0 {
		Parametrs.GuiInit += fmt.Sprintf("port=%d\n", pIni.Port)
	}
	if pIni.Log == 1 {
		Parametrs.GuiInit += "log=1\n"
	} else if pIni.Log == 2 {
		Parametrs.GuiInit += "log=2\n"
	}

	if pIni.FontMain.Family != "" {
		Parametrs.PFontMain = egui.CreateFont(&egui.Font{Name: "fa",
			Family: pIni.FontMain.Family, Height: pIni.FontMain.Height})
	} else {
		Parametrs.PFontMain = egui.CreateFont(&egui.Font{Name: "fa", Family: "Courier New", Height: -19})
	}
	if pIni.FontMenu.Family != "" {
		Parametrs.PFontMenu = egui.CreateFont(&egui.Font{Name: "fm",
			Family: pIni.FontMenu.Family, Height: pIni.FontMenu.Height})
	} else {
		Parametrs.PFontMenu = Parametrs.PFontMain
	}
	if pIni.FontText.Family != "" {
		Parametrs.PFontText = egui.CreateFont(&egui.Font{Name: "ft",
			Family: pIni.FontText.Family, Height: pIni.FontText.Height})
	} else {
		Parametrs.PFontText = Parametrs.PFontMain
	}

	db, _ := sql.Open(pIni.DB.Name, pIni.DB.Path)
	Parametrs.DB = db
}

func getxml(sPath string, pXml interface{}) string {

	data, err := ioutil.ReadFile(sPath)
	if err != nil {
		return fmt.Sprintf("error reading file: %v", err)
	}

	err = xml.Unmarshal([]byte(data), pXml)
	if err != nil {
		return fmt.Sprintf("unmarshal error: %v", err)
	}
	return ""
}
