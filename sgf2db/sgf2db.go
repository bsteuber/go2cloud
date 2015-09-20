package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func SgfUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// TODO: Dispatch on unparseable sgf
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// TODO: Do something useful with this
	fmt.Println(string(body))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(`{"status" : "ok"}`); err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()
	router.POST("/upload", SgfUpload)
	log.Fatal(http.ListenAndServe(":8080", router))
}
