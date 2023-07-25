package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCanLoadTaskListFromContent(t *testing.T) {
	tests := []struct {
		name           string
		content        string
		expectedResult []Task
	}{
		{name: "empty content", content: "", expectedResult: nil},
		{
			name:    "one task",
			content: "task 1,n,n",
			expectedResult: []Task{
				{Description: "task 1", completed: false, hidden: false},
			},
		},
		{
			name:    "2 tasks",
			content: "task 1,n,n\ntask 2,y,y",
			expectedResult: []Task{
				{Description: "task 1", completed: false, hidden: false},
				{Description: "task 2", completed: true, hidden: true},
			},
		},
		{
			name:    "3 tasks",
			content: "task 1,n,n\ntask 2,y,y\ntask 3,n,n\ntask 4,y,y",
			expectedResult: []Task{
				{Description: "task 1", completed: false, hidden: false},
				{Description: "task 2", completed: true, hidden: true},
				{Description: "task 3", completed: false, hidden: false},
				{Description: "task 4", completed: true, hidden: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := tt.content
			tasks, err := convertContentToTasks(content) // testee
			if err != nil {
				t.Errorf("Couldn't read from %s", content)
			}
			if !Equal(tasks, tt.expectedResult) {
				t.Errorf(" %v wasnt like %v ", tasks, tt.expectedResult)
			}
		})

	}
}

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
	json_content := map[string]interface{} {
		"description": description,
	}
	body, _ := json.Marshal(json_content)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST","http://example.com/add",bytes.NewReader(body))
	
	list, _ := add(w,r) //testee

	println(list)
	if !Equal(list, []Task{{Description: description, completed: false, hidden: false}}){
		t.Error("Task list weren't equal")
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
