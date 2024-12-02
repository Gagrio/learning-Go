package main

import (
	"fmt"
	"net/http"
)

var crashCourse = "Watch GO crash course"
var fullCourse = "Watch Go Full course"
var goPlay = "GO play"
var taskItems = []string{crashCourse, fullCourse, goPlay}

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/showTasks", showTasks)

	http.ListenAndServe("localhost:8080", nil)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to my Todolist App!"
	fmt.Fprintf(writer, greeting)
}
