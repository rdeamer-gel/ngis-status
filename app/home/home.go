package home

import (
	"github.com/rdeamer-gel/ngis-status/app/common"
	"html/template"
	"net/http"
	"os"
	"strings"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
	type Page struct {
		Title    string
		Active   string
		Version  string
		Versions map[string]string
	}

	oc_project := os.Getenv("OC_PROJECT")
	version := os.Getenv("OC_PROJECT")

	versions := make(map[string]string)

	for _, e := range os.Environ() {
		if strings.HasPrefix(strings.Split(e, "=")[0], "NGV_") {
			pair := strings.Split(e, "=")
			pair[0] = strings.Replace(pair[0], "NGV_NEW_", "", -1)
			versions[pair[0]] = pair[1]
		}
	}

	p := Page{
		Active:   "home",
		Title:    "Home",
		Version:  version,
		Versions: versions,
	}

	common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
