package main

import (
	"context"
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

	http.HandleFunc("/", getRoot)
	// http.HandleFunc("/add", add)

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

func addFormFunc(tasks *[]Task) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request)  {
		ctx := r.Context()
		ctx = context.WithValue(ctx,"tasks", tasks)
		addForm(w, r.WithContext(ctx))
	}
}
func addForm(w http.ResponseWriter, r *http.Request) {
	_, err := addNewTaskToAppList(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)	
		return
	}
	http.Redirect(w,r,"/", http.StatusFound)
}

func addNewTaskToAppList(r *http.Request) (tasks []Task, err error) {
	if r.Method == "POST" {
		myvar := r.Context().Value("tasks").(*[]Task)
		tasks = append(*myvar,
			Task{
				Description: r.PostFormValue("description"),
				completed:   false,
				hidden:      false,
			},
		)
		*myvar = tasks

	}
	return tasks, nil
}
