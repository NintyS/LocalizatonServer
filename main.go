package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type User struct {
	Address   string
	Latitude  float64
	Longitude float64
	//Time      time.Time
}

var Addresses []User

func recPos(w http.ResponseWriter, req *http.Request) {

	body, _ := io.ReadAll(req.Body)

	fmt.Println("Body:", string(body))

	for _, v := range Addresses {
		if v.Address == req.RemoteAddr {
			fmt.Printf("Urządzenie: %s wysłało swoją pozycję: lat: %f, lon: %f o: %s\n", req.RemoteAddr, v.Latitude, v.Longitude, time.Now().Format(time.TimeOnly))
			return
		}
	}

	var user User

	err := json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Err:", err)
	}

	user.Address = req.RemoteAddr

	fmt.Printf("Urządzenie: %s wysłało swoją pozycję: lat: %f, lon: %f o: %s\n", user.Address, user.Latitude, user.Longitude, time.Now().Format(time.TimeOnly))

	Addresses = append(Addresses, user)
}

func RegDev(w http.ResponseWriter, req *http.Request) {

}

func main() {

	http.HandleFunc("/receivePosition", recPos)

	fmt.Println("Error:", http.ListenAndServe(":80", nil))
}
