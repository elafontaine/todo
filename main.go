package main

import (
	"errors"
	"fmt"
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

const fileName = "tasks.csv"


func main() {
	http.HandleFunc("/",getRoot)
	http.HandleFunc("/add", add)

	err:=  http.ListenAndServe(":5000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
/*	fmt.Println("Got / from ", r.RemoteAddr)
	// open file containing a list of tasks
	fh, err := os.Open(x)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	defer fh.Close()

	csvReader := csv.NewReader(fh)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+ fileName, err)
	}

	var tasks []Task

	
	_, err = io.WriteString(w, tpl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}*/
}

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.PostForm.Get("Description")
	}
	getRoot(w, r)
}