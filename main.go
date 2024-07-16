package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Users struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	When      string  `json:"time"`
}

var _users []Users

func getLocalization(w http.ResponseWriter, req *http.Request) {

	fmt.Println(_users)

	marshal, err := json.Marshal(_users)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Println("JSON:", string(marshal))
	fmt.Fprintf(w, string(marshal))
}

func setLocalization(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err.Error())
		return
	}
	for key, value := range req.Form {
		fmt.Println(key, value)
	}

	var u Users

	u.Name = req.Form.Get("name")
	var latitude = req.Form.Get("latitude")
	u.Latitude, _ = strconv.ParseFloat(latitude, 64)
	var longitude = req.Form.Get("longitude")
	u.Longitude, _ = strconv.ParseFloat(longitude, 64)
	u.when = req.Form.Get("time") //client gives server his local time

	fmt.Println("Stworzony użytkownik:", u)

	_users = append(_users, u)

}

func main() {

	http.HandleFunc("/getLocalization", getLocalization)
	http.HandleFunc("/setLocalization", setLocalization)

	http.ListenAndServe(":8090", nil)
}
