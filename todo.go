package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func getTasksFromFile(fileName string) (tasks []Task, err error) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	bytes_content, err := os.ReadFile(fileName)
	if err != nil {
		return
	}

	content := string(bytes_content)
	tasks, err = convertContentToTasks(content)
	if err != nil {
		log.Panicln("Unable to parse file as CSV for "+fileName, err)
		return
	}
	return tasks, nil
}

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
