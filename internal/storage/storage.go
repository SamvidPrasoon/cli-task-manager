package storage

import (
	"cli-task-manager/internal/task"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const defaultFile = "tasks.json"

type Store struct {
	filePath string
}

func New(filePath string) *Store {
	if filePath == "" {
		filePath = defaultFile
	}
	return &Store{filePath: filePath}
}

func (s *Store) Load() ([]task.Task, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []task.Task{}, nil
		}
		return nil, fmt.Errorf("reading task file: %w", err)
	}
	var tasks []task.Task
	// unmarshal means parse
	err = json.Unmarshal(data, &tasks)
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("parsing task file: %w", err)
	}
	return tasks, nil
}

func (s *Store) Save(tasks []task.Task) error {
	// marshal means stringify
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("serializing tasks :%w", err)
	}
	if err := os.WriteFile(s.filePath, data, 0644); err != nil {
		return fmt.Errorf("writing task %w", err)
	}
	return nil
}
