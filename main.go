package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"net/http"
)

type Todos struct {
	UserId    int    `json:"userdId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	//receive Get response
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error getting the response: ", res.Status)
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	// fmt.Println("Data: ", string(data))

	var todo Todos
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Todos: ", todo)

	fmt.Println("Title response: ", todo.Title)
	fmt.Println("Completed response: ", todo.Completed)
}

func performPostRequest() {
	todo := Todos{
			UserId:    1,
			Title:     "Learn golang",
			Completed: true,
	}

	//convert struct to json data
	jsonData, err := json.Marshal(&todo)
	if err != nil {
		fmt.Println("Error marshalling: ", err)
		return
	}

	//convert json data into string
	jsonString := string(jsonData)

	//convert string to io.reader
	jsonReader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos"

	//send POST request
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error posting: ", err)
		return
	}

	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response data: ", string(data))

	fmt.Println(res.Status)

}

func performUpdateRequest() {
	todo := Todos{
		UserId:    10213,
		Title:     "Learn golang programme",
		Completed: false,
	}

	//convert struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling data: ", err)
		return
	}

	//convert json to string
	jsonString := string(jsonData)

	// convert string to io.reader
	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request: ", err)
		return
	}
	req.Header.Add("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request: ", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading the data: ", err)
		return
	}
	fmt.Println("Response: ", string(data))
	fmt.Println("status response: ",res.Status)

}

func performDeleteRequest() {
	//delete with id
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create a DELETE request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error requesting data: ", err)
		return
	}

	//send client request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error sending request", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading data: ", err)
		return
	}
	fmt.Println(string(data))
	fmt.Println(res.Status)
}

func main() {
	//performGetRequest()
	//performPostRequest()
	//performUpdateRequest()
	performDeleteRequest()
}
