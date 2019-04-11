package home

import (
	"github.com/bobbydeveaux/ngis-status/app/common"
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

	version := os.Getenv("NGIS_VERSION")

	versions := make(map[string]string)

	for _, e := range os.Environ() {
		if strings.HasPrefix(strings.Split(e, "=")[0], "NGV_") {
			pair := strings.Split(e, "=")
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
