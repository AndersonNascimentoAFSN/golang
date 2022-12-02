package webServer

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func Server() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/tasks", TaskHandler)
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Name: "Aprender Go",
		Done: true,
	}

	j, _ := json.Marshal(task)
	w.Write(j)

	// t := template.Must(template.ParseFiles("src/webServer/task.html"))
	// t.Execute(w, task)
}
