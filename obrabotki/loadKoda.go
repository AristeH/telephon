package obrabotki
/* 
import (
	ref "Telephon/datashema/reference"
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)
//Loadcode a
func Loadcode(file string) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=telephone sslmode=disable password=1234")
	//db, err := gorm.Open("sqlite3", "C:\\SQLite\\gorm.db")
	db.AutoMigrate(&ref.Codetel{}, &ref.Region{}, &ref.Operator{}, &ref.Oblast{})
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
			r = line[5][0:Index]
			o = line[5][Index+1:]
		} else {
			r = line[5]
			o = line[5]
		}
		obl := ref.Oblast{}
		db.FirstOrCreate(&obl, ref.Oblast{Name: strings.TrimSpace(o)})
		reg := ref.Region{}
		db.FirstOrCreate(&reg, ref.Region{Name: strings.TrimSpace(r), Oblast: *obl})
		oper := ref.Operator{}
		db.FirstOrCreate(&oper, ref.Operator{Name: strings.TrimSpace(line[4])})
		cd := ref.Codetel{}
		db.FirstOrCreate(&cd, ref.Codetel{Code: line[0], S: strings.TrimSpace(strings.Trim(line[1], "/'")), Po: strings.TrimSpace(strings.Trim(line[2], "/'")), IDregion: reg.ID, IDoperator: oper.ID})
	}
} */
