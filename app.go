package main

import (
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	flag.Set("logtostderr", "true")

	defer glog.Flush()

	rtr := mux.NewRouter()
	rtr.HandleFunc("/user/{name:[a-z]+}/profile", profile).Methods("GET")

	http.Handle("/", rtr)

	glog.Info("Listening on port ", 3000)

	http.ListenAndServe(":3000", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	userProfile := User{name, "email@domain.com"}

	js, err := json.Marshal(userProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
