package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url string = "https://randomuser.me/api/?inc=gender,name,nat,email,dob,login"

type randName struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}
type randDob struct {
  Date string `json:"date"`
  Age int `json"age"`
}
type randLogin struct {
  UUID string `json:"uuid"`
  Username string `json:"username"`
  Password string `json:"password"`
  Salt string `json:"salt"`
  MD5 string `json:"md5"`
  SHA1 string `json:"sha1"`
  SHA256 string `json:"sha256"`
}
type randResults struct {
	Gender string   `json:"gender"`
	Name   randName `json:"name"`
	Email  string   `json:"email"`
	Nat    string   `json"nat"`
  DOB randDob `json:"dob"`
  Login randLogin `json:"login"`
}
type user struct {
	Results []randResults `json:"results"`
}

func main() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var u user
	err = json.Unmarshal(data, &u)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("user: %+v", u)
}

