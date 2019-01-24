package bdTable

import (
	"bufio"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_"github.com/jinzhu/gorm/dialects/postgres"
"time"
	"strconv"
	"github.com/jinzhu/gorm"
	"os"
	"fmt"
	"strings"
	"io"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"golang.org/x/text/transform"
	//	"encoding/csv"

)
// Таблица области. Содержит в себе области РФ
type Oblast struct {
	ID   uint64  `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name string
}

// Таблица операторы. Сотовые операторы.
type Operator struct {
	ID   uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name string
}



// Таблица СчетОператора. Счет за услуги выставленный за звонки телефона.
type Chet struct {
	ID          uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	IDtelephon  uint64
	Date        time.Time
	Zvonki      [] Zvonki `gorm:"ForeignKey:IDchet;AssociationForeignKey:Refer"`
	Nomerscheta string
	IDregion    uint64
}

// Таблица Звонки. Все звонки телефона.
type Zvonki struct {
	Date      time.Time
	Number    string
	Amount    float64
	IDunit    uint64
	IDservice uint64
	IDregion  uint64
	Price     float64
	IDchet    uint64
}

// Таблица КодаТелефонов. Кода телефонов городов и сотовых операторов.
type Codetel struct {
	ID         uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Code       string
	IDoperator uint64
	S          string
	Po         string
	IDregion   uint64
}

// Таблица Регионы.
type Region struct {
	ID       uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name     string
	IDoblast uint64

}

// Таблица ЕдиницыИхмерений.
type Unit struct {
	ID   uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name string
}

// Таблица Услуги.
type Service struct {
	ID   uint64`sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name string
}

// Таблица Телефоны. Содержит код и номер телефона.
type Telephon struct {
	ID       uint64`sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Code     string
	Nomer    string
	Number   string
	IDregion uint64
	Our      bool
}

func loadcode(file string) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=telephone sslmode=disable password=1234")
	//db, err := gorm.Open("sqlite3", "C:\\SQLite\\gorm.db")
	db.AutoMigrate(&Codetel{}, &Region{}, &Operator{}, &Oblast{})
	if err != nil {
		panic("failed to connect database")
		return
	}
	f, err := os.Open(file)
	defer f.Close()

	input := bufio.NewReader(f)

	for {

		stex, _ := input.ReadString('\n')
		sr := strings.NewReader(stex)
		str := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
		buf, _ := ioutil.ReadAll(str)
		s := string(buf)
		line := strings.Split(s, ";")

		r := ""
		o := ""
		Index := strings.LastIndex(line[5], "|")
		fmt.Println(Index)
		fmt.Println(line[5])
		if Index > 0 {
			r = line[5] [0:Index]
			o = line[5] [Index+1:]
		} else {
			r = line[5]
			o = line[5]
		}
		obl := Oblast{}
		db.FirstOrCreate(&obl, Oblast{Name: strings.TrimSpace(o)})
		reg := Region{}
		db.FirstOrCreate(&reg, Region{Name: strings.TrimSpace(r), IDoblast: obl.ID})
		oper := Operator{}
		db.FirstOrCreate(&oper, Operator{Name: strings.TrimSpace(line[4])})
		cd := Codetel{}
		db.FirstOrCreate(&cd, Codetel{Code: line[0], S: strings.TrimSpace(strings.Trim(line[1], "/'")), Po: strings.TrimSpace(strings.Trim(line[2], "/'")), IDregion: reg.ID, IDoperator: oper.ID})

		//fmt.Println(k)
	}
}

// записать телефон и вернуть его ID
func Recordtel(db gorm.DB, y string, Our bool) uint64 {
	codetel := Codetel{}
	code, nomer := "", "";
	tel := Telephon{}
	if len(y) == 10 && y[0:1] == "9" { //  Мобильный номер
		code = y[0:3]
		nomer = y[3:10]
		db.Where("Code = ? AND S <= ? AND Po<= ? ", code, nomer, nomer).Find(&codetel)
		if Our {
			db.FirstOrCreate(&tel, Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion, Our: true})
		} else {
			db.FirstOrCreate(&tel, Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion})
		}
	} else { // городской номер
		for i := 7; i > 4; i-- {
			db.Where("Code LIKE ?", y[0:i]).Find(&codetel)
			if codetel.ID > 0 {

				if Our {
					db.FirstOrCreate(&tel, Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion, Our: true})
				} else {
					db.FirstOrCreate(&tel, Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion})
				}
				break
			}
		}

	}
	return tel.ID
}

func preopbr(file string) {
	var tel string
	var nomerscheta uint64
	var line [] string
	//	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=Telephone sslmode=disable password=1234")
	db, err := gorm.Open("sqlite3", "C:\\SQLite\\gorm.db")
	if err != nil {
		panic("failed to connect database")
		return
	}
	// обновление таблиц базы
	db.AutoMigrate(&Zvonki{}, &Region{}, &Chet{}, &Unit{}, &Service{}, &Telephon{})
	f, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ошибка: %v\n", err)
	}
	defer f.Close()

	input := bufio.NewReader(f)
	ff := "Детализация"
	k := 0
	tel = ""
	ud := true
	Loc, _ := time.LoadLocation("Local")
	for {
		k++
		stex, err := input.ReadString('\n')
		sr := strings.NewReader(stex)
		str := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
		buf, _ := ioutil.ReadAll(str)
		s := string(buf)
		line = strings.Split(s, ";")
		fmt.Printf("%s\n", line)
		if err == io.EOF {
			break
		}
		if err == nil {
		}
		if len(s) > 10 {
			if s[1] == '0' ||
				s[1] == '1' ||
				s[1] == '2' ||
				s[1] == '3' {
				Dt := strings.Trim(line[0], "\"") + " " + strings.Trim(line[3], "\"")
				d, _ := time.ParseInLocation("02.01.06 15:04:05", Dt, Loc)
				//fmt.Println(d)
				reg := Region{}
				db.FirstOrCreate(&reg, Region{Name: strings.Trim(line[10], "\"")})
				un := Unit{}
				db.FirstOrCreate(&un, Unit{Name: strings.Trim(line[8], "\"")})
				ser := Service{}
				db.FirstOrCreate(&ser, Service{Name: strings.Trim(line[9], "\"")})
				k, _ := strconv.ParseFloat(strings.Trim(line[6], "\""), 10)
				j, _ := strconv.ParseFloat(strings.Trim(line[11], "\""), 10)
				nt := strings.Trim(line[5], "\"")
				if len(nt) > 2 {
					if nt[0] == '7' ||
						nt[0:1] == "89" {
						nt = nt[1:]
					}
				}
				Recordtel(*db, nt, false)
				db.Create(&Zvonki{d, nt, k, un.ID, ser.ID, reg.ID, j, nomerscheta})
			}
			if strings.Contains(s, ff) {
				if tel != s {
					tel = s
					ud = true
				} else {
					ud = false
				}
				Index := strings.Index(line[0], "+")
				fmt.Println(Index)
				y := line[0] [Index+1:len(line[0])-1]
				if y[0] == '7' ||
					y[0:1] == "89" {
					y = y[1:]
				}

				regID := Recordtel(*db, y, true)

				tekschet := Chet{}
				datascheta := time.Date(2017, 03, 01, 0, 0, 0, 0, Loc)
				db.FirstOrCreate(&tekschet, Chet{IDtelephon: regID, Date: datascheta})
				if ud {
					db.Where("IDchet=? ", tekschet.ID).Delete(Zvonki{})
				}
				fmt.Println(tekschet)

				nomerscheta = tekschet.ID
				fmt.Printf("%s\n", nomerscheta)
			}

		}
	}
}
