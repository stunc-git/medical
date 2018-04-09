package system

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/denisbakhtin/medical/helpers"
)

var tmpl *template.Template

func loadTemplates() {
	tmpl = template.New("").Funcs(template.FuncMap{
		"isActive":        helpers.IsActive,
		"stringInSlice":   helpers.StringInSlice,
		"dateTime":        helpers.DateTime,
		"date":            helpers.Date,
		"mainMenu":        helpers.MainMenu,
		"scrollMenu":      helpers.ScrollMenu,
		"oddEvenClass":    helpers.OddEvenClass,
		"truncate":        helpers.Truncate,
		"sellingPreface":  helpers.SellingPreface,
		"promoTill":       helpers.PromoTill,
		"cityList":        helpers.CityList,
		"eqRU":            helpers.EqRU,
		"allReviews":      helpers.AllReviews,
		"isFirstInTheRow": helpers.IsFirstInTheRow,
		"isLastInTheRow":  helpers.IsLastInTheRow,
		"isLast":          helpers.IsLast,
	})

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && strings.HasSuffix(f.Name(), ".tmpl") {
			var err error
			tmpl, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	}

	err := filepath.Walk("views", fn)
	if err != nil {
		log.Panic(err)
	}
}

//GetTemplates exports loaded templates
func GetTemplates() *template.Template {
	return tmpl
}
