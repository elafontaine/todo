package main

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Task struct {
	completed   bool
	hidden      bool
	Description string
}

var todo_template *template.Template

const fileName = "tasks.csv"

var tasks []Task = []Task{}

func main() {
	var err error // required for the todo_template variable
	todo_template, err = template.New("root").Parse(tpl)
	if err != nil {
		panic(fmt.Sprintf("Couldn't prepare template, something is very wrong : %v", err))
	}
	tasks, err = getTasksFromFile(fileName)
	if err != nil {
		panic(err)
	}

	asset_directory := http.Dir("assets")
	asset_fileserver := http.FileServer(asset_directory)
	http.Handle("/assets/", http.StripPrefix("/assets/",asset_fileserver))
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/add", addFormFunc(&tasks))

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

func addFormFunc(tasks *[]Task) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "tasks", tasks)
		addForm(w, r.WithContext(ctx))
	}
}
func addForm(w http.ResponseWriter, r *http.Request) {
	addNewTaskToAppList(r)
	http.Redirect(w, r, "/", http.StatusFound)
}

func addNewTaskToAppList(r *http.Request) {
	if r.PostFormValue("description") != "" {
		myvar := r.Context().Value("tasks").(*[]Task)
		*myvar = append(*myvar,
			Task{
				Description: r.PostFormValue("description"),
				completed:   false,
				hidden:      false,
			},
		)
		// https://stackoverflow.com/a/38105687/2184575
	}
}
