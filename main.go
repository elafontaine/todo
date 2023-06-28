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
		<style> 
		input[type=checkbox]:checked + input[type=text].strikethrough {
			text-decoration: line-through;
			color: var(--bs-gray);
			text-decoration-color: var(--bs-gray);
		  }
		</style>
	</head>
	<body>
		{{range $i, $a := .}}
			<div><input id="formCheckBox-2" type="checkbox" /><input class="strikethrough" type="text" value="{{.Description}}" for="formCheckBox-1" />
		</div>
		{{else}}
		<div><strong>no rows, click the + button</strong></div>
		{{end}}
		<svg class="bi bi-plus-circle justify-content-md-end align-items-md-end" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" fill="currentColor" viewBox="0 0 16 16" style="font-size: 55px;position: absolute;bottom: 25px;right: 25px;">
    <path d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14zm0 1A8 8 0 1 0 8 0a8 8 0 0 0 0 16z"></path>
    <path d="M8 4a.5.5 0 0 1 .5.5v3h3a.5.5 0 0 1 0 1h-3v3a.5.5 0 0 1-1 0v-3h-3a.5.5 0 0 1 0-1h3v-3A.5.5 0 0 1 8 4z"></path>
</svg>
	</body>
</html>`

var todo_template *template.Template

const fileName = "tasks.csv"

var tasks []Task = []Task{}  //Global list with mutex (shared across requests)

func main() {
	var err error // required for the todo_template variable
	todo_template, err = template.New("root").Parse(tpl)
	if err != nil {
		panic(fmt.Sprintf("Couldn't prepare template, something is very wrong : %v", err))
	}
	tasks,err = getTasksFromFile(fileName)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/add", add)

	err = http.ListenAndServe(":5000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		panic(fmt.Sprintf("error starting server: %v\n", err))
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

func getTasksFromFile(fileName string) ([]Task, error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes_content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return nil, err
	}

	content := string(bytes_content)
	tasks, err := convertContentToTasks(content)
	if err != nil {
		log.Panicln("Unable to parse file as CSV for "+fileName, err)
		return nil, err
	}
	return tasks, nil
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.PostForm.Get("Description")
	}
	getRoot(w, r)
}
