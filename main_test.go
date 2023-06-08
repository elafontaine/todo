package main

import (
	"encoding/csv"
	"strings"
	"testing"
)

func TestCanLoadTaskListFile(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := csv.NewReader(strings.NewReader(tt.content))
			records, err := r.ReadAll()
			if err != nil {
				t.Errorf("Couldn't read from %s", tt.content)
			}
			tasks:= []Task{}
			for _, record := range records {
				task := convertToTask(record)
				tasks = append(tasks, task)
			}
			if !Equal(tasks, tt.expectedResult) {
				t.Errorf(" %v wasnt like %v ", tasks, tt.expectedResult)
			}
		})

	}
}

func convertToTask(record []string) Task {
	task := Task{Description: record[0], completed: parseBool(record[1]), hidden: parseBool(record[2])}
	return task
}

func parseBool(s string) bool {
	if s == "y" || s == "Y" {
		return true
	}
	return false
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