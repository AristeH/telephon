package obrabotki

/* import (
	"github.com/jinzhu/gorm"
	"os"
	"fmt"
	"bufio"
	"time"
	"strings"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"strconv"
	ref "Telephon/datashema/reference"
	doc "Telephon/datashema/document"
	"golang.org/x/text/encoding/charmap"
)

// записать телефон и вернуть его ID
func recordtel(db gorm.DB, y string, Our bool) uint64 {
	codetel := ref.Codetel{}
	code, nomer := "", "";
	tel := ref.Telephon{}
	if len(y) == 10 && y[0:1] == "9" { //  Мобильный номер
		code = y[0:3]
		nomer = y[3:10]
		db.Where("Code = ? AND S <= ? AND Po<= ? ", code, nomer, nomer).Find(&codetel)
		if Our {
			db.FirstOrCreate(&tel, ref.Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion, Our: true})
		} else {
			db.FirstOrCreate(&tel, ref.Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion})
		}
	} else { // городской номер
		for i := 7; i > 4; i-- {
			db.Where("Code LIKE ?", y[0:i]).Find(&codetel)
			if codetel.ID > 0 {

				if Our {
					db.FirstOrCreate(&tel, ref.Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion, Our: true})
				} else {
					db.FirstOrCreate(&tel, ref.Telephon{Number: strings.Trim(y, "\""), Code: code, Nomer: nomer, IDregion: codetel.IDregion})
				}
				break
			}
		}

	}
	return tel.ID
}

// Preopbr a
func Preopbr(file string) {
	var tel string
	var nomerscheta uint64
	var line [] string
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=Telephone sslmode=disable password=1234")
	//db, err := gorm.Open("sqlite3", "C:\\SQLite\\gorm.db")
	if err != nil {
		panic("failed to connect database")
		return err 
	}
	// обновление таблиц базы
	db.AutoMigrate(&doc.ChetZvonki{}, &ref.Region{}, &doc.Chet{}, &ref.Unit{}, &ref.Service{}, &ref.Telephon{})
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
				reg := ref.Region{}
				db.FirstOrCreate(&reg, ref.Region{Name: strings.Trim(line[10], "\"")})
				un := ref.Unit{}
				db.FirstOrCreate(&un, ref.Unit{Name: strings.Trim(line[8], "\"")})
				ser := ref.Service{}
				db.FirstOrCreate(&ser, ref.Service{Name: strings.Trim(line[9], "\"")})
				k, _ := strconv.ParseFloat(strings.Trim(line[6], "\""), 10)
				j, _ := strconv.ParseFloat(strings.Trim(line[11], "\""), 10)
				nt := strings.Trim(line[5], "\"")
				if len(nt) > 2 {
					if nt[0] == '7' ||
						nt[0:1] == "89" {
						nt = nt[1:]
					}
				}
				recordtel(*db, nt, false)
				db.Create(&doc.ChetZvonki{d, nt, k, un.ID, ser.ID, reg.ID, j, nomerscheta})
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

				regID := recordtel(*db, y, true)

				tekschet := doc.Chet{}
				datascheta := time.Date(2017, 03, 01, 0, 0, 0, 0, Loc)
				db.FirstOrCreate(&tekschet, doc.Chet{IDtelephon: regID, Date: datascheta})
				if ud {
					db.Where("IDchet=? ", tekschet.ID).Delete(doc.ChetZvonki{})
				}
				fmt.Println(tekschet)

				nomerscheta = tekschet.ID
				fmt.Printf("%s\n", nomerscheta)
			}

		}
	}
}
 */