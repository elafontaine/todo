package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Task struct {
	completed   bool
	hidden      bool
	Description string
}

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Eric's todo</title>
	</head>
	<body>
		{{range .}}<div>{{.Description}}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
	</body>
</html>`

var todo_template *template.Template

const fileName = "tasks.csv"

var tasks []Task = []Task{}


func main() {
	var err error // required for the todo_template variable
	todo_template, err = template.New("root").Parse(tpl)
	if err != nil {
		log.Fatalf("Couldn't prepare template, something is very wrong : %v", err)
	}
	tasks = getTasks(fileName)
	if err != nil {
		log.Fatalf("Couldn't prepare task list from file that exist: %v", err)
	}

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/add", add)

	err = http.ListenAndServe(":5000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got / from ", r.RemoteAddr)
	// open file containing a list of tasks
	err := todo_template.Execute(w, tasks)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func GetTasks() ([]Task) {
	return tasks
}

func getTasks(fileName string) ([]Task) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		panic(err)
	}

	bytes_content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		panic(err)
	}

	content := string(bytes_content)
	tasks, err := convertContentToTasks(content)
	if err != nil {
		log.Panicln("Unable to parse file as CSV for "+fileName, err)
		panic(err)
	}
	return tasks
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.PostForm.Get("Description")
	}
	getRoot(w, r)
}
