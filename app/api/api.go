package api

import (
	"github.com/rdeamer-gel/ngis-status/app/common"
	"net/http"
	"os"
	"strings"
	"io"

	"encoding/json"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {


	versions := make(map[string]string)

	versions["NGIS_VERSION"] = os.Getenv("NGIS_VERSION")
	for _, e := range os.Environ() {
		if strings.HasPrefix(strings.Split(e, "=")[0], "NGV_") {
			pair := strings.Split(e, "=")
			pair[0] = strings.Replace(pair[0], "NGV_NEW_", "", -1)
			versions[pair[0]] = pair[1]
		}
	}


	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    strVersions, err := json.Marshal(versions)
    common.CheckError(err, 2)
    io.WriteString(w, string(strVersions))

}
