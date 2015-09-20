package server

import (
	"fmt"
	"github.com/bsteuber/go2cloud/tools/http/json"
	"github.com/bsteuber/go2cloud/tools/http/status"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func SgfUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	sgf := string(body)
	// TODO: Dispatch on unparseable sgf
	if err != nil {
		json.Response(w, status.InternalServerError, err)
	}

	// TODO: Do something useful with this
	fmt.Println(sgf)
	json.Response(w, status.OK, nil)
}

func Run(port string) {
	router := httprouter.New()
	router.POST("/upload", SgfUpload)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
