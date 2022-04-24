package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Projj struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Comment  string `json:"comment"`
	Date     string `json:"date"`
	Pid      int    `json:"project_id"`
}

func main() {
	post()
}

func get() {
	// link := "http://api.royal-server.ml/comments"
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("http://api.royal-server.ml/comments")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Projj
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
}

func post() {
	fmt.Println("2. Performing Http Post...")
	todo := Projj{6, "mabit", "lorem ipsum dolor sit amet", "mabit", 2}
	jsonReq, err := json.Marshal(todo)
	fmt.Println(bytes.NewBuffer(jsonReq))
	resp, err := http.Post("http://api.royal-server.ml/comments", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var todoStruct Projj
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)
}
