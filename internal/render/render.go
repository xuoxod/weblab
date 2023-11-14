package render

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet"
	"github.com/xuoxod/weblab/pkg/utils"
)

// may also use an absolute path:
var root, _ = os.Getwd()
var View = jet.NewHTMLSet(filepath.Join(root, "views"))

func InitViews() {
	View.SetDevelopmentMode(true)
	View.AddGlobal("appver", "0.0.3")
	View.AddGlobal("copyright", utils.CopyrightDate())
	View.AddGlobal("appname", "Awesome Web App")
	View.AddGlobal("appdate", fmt.Sprintf("%v", utils.DateTimeStamp()))
}

func Render(w http.ResponseWriter, tmpl string, vars jet.VarMap, data map[string]interface{}) error {
	view, err := View.GetTemplate(tmpl)

	if vars == nil {
		vars = make(jet.VarMap)
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = view.Execute(w, vars, data)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
