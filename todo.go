package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func convertContentToTasks(content string) (tasks []Task, err error) {
	r := csv.NewReader(strings.NewReader(content))
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Couldn't read (%w) from content %s", err, content)
	}
	tasks = []Task{}
	for _, record := range records {
		task := convertToTask(record)
		tasks = append(tasks, task)
	}
	return tasks, nil
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