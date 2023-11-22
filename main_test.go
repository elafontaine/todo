package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestCanAddFormReturnALocationToGetRoot(t *testing.T) {
	description := "New task to add"
	params := url.Values{}
	params.Add("description", description)
	tasks := []Task{}

	r := httptest.NewRequest("POST", "http://example.com/tasks/", strings.NewReader(params.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	addFormFunc(&tasks)(w, r) //testee

	if !Equal(tasks, []Task{{Description: description, completed: false, hidden: false}}) {
		t.Error("Task list weren't equal")
	}
	if w.Result().StatusCode != http.StatusFound {
		t.Error("Did not receive expected redirect to root")
	}
	if w.Result().Header["Location"][0] != "/" {
		t.Error("Did not receive the location header to point to /")
	}
}

func TestAddFormHasWrongInput(t *testing.T){
	
}

func TestCanGetTasksFromUri(t *testing.T) {
	tasks := []Task{}

	r := httptest.NewRequest("GET", "http://example.com/tasks/", nil)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	addFormFunc(&tasks)(w, r) //testee

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
