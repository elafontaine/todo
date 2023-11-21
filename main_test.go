package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestCreateTaskFileIfDoesNotExist(t *testing.T) {
	fileName := "random-filename.test"
	for i := 0; i < 2; i++ { // 0- no file, 1- file should exist and still succeed
		testGetTasksFromFile(fileName, t)
	}
	err := os.Remove(fileName)
	if err != nil {
		t.Errorf("Couldn't delete test file : %v", err)
	}
}

func testGetTasksFromFile(fileName string, t *testing.T) {
	tasks, err := getTasksFromFile(fileName) 		//testee
	if err != nil {
		t.Error("Something very wrong happened")
	}
	if len(tasks) != 0 {
		t.Error("Content of the file isn't empty!")
	}
}

func TestCanAddTaskToTasksList(t *testing.T){
	description := "New task to add"
	formContent := "description="+ description
	r := httptest.NewRequest("POST","http://example.com/tasks/",strings.NewReader(formContent))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tasks := []Task{}
	ctx := context.WithValue(r.Context(),"tasks",&tasks)
	r = r.WithContext(ctx)
	output_tasks,_ := addNewTaskToAppList(r) //testee

	fmt.Printf("Tasks : %v", output_tasks)
	if !Equal(tasks, []Task{{Description: description, completed: false, hidden: false}}){
		t.Error("Task list weren't equal")
	}
}


func TestCanAddFormReturnALocationToGetRoot(t *testing.T) {
	description := "New task to add"
	params := url.Values{}
	params.Add("description", description)
	tasks := []Task{}

	r := httptest.NewRequest("POST","http://example.com/tasks/", strings.NewReader(params.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	
	addFormFunc(&tasks)(w,r) //testee

	if w.Result().StatusCode != http.StatusFound {
		t.Error("Did not receive expected redirect to root")
	}
	if w.Result().Header["Location"][0] != "/" {
		t.Error("Did not receive the location header to point to /")
	}
}

func TestCanGetTasksFromUri(t *testing.T) {
	description := "New task to add"
	params := url.Values{}
	params.Add("description", description)
	tasks := []Task{}

	r := httptest.NewRequest("GET","http://example.com/tasks/", strings.NewReader(params.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	
	addFormFunc(&tasks)(w,r) //testee

	if w.Result().StatusCode != http.StatusFound {
		t.Error("Did not receive expected redirect to root")
	}
	if w.Result().Header["Location"][0] != "/" {
		t.Error("Did not receive the location header to point to /")
	}
}

func Equal(a, b []Task) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
