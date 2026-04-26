package cmd

import (
	"cli-task-manager/internal/storage"
	"cli-task-manager/internal/task"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var store = storage.New("")

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A production-grade CLI task manager",
	Long:  `Manage your tasks from the terminal with persistent JSON storage.`,
}

// flags
// filter variable
var filter string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	listCmd.Flags().StringVarP(
		&filter,
		"filter",
		"f",
		"",
		"filter by status",
	)

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := store.Load()
		if err != nil {
			return fmt.Errorf("loading tasks: %w", err)
		}
		newTask := task.New(len(tasks)+1, args[0])
		tasks = append(tasks, newTask)
		if err := store.Save(tasks); err != nil {
			return fmt.Errorf("saving tasks:%w", err)
		}
		fmt.Printf("✅ Added task #%d: %s\n", newTask.Id, newTask.Title)
		return nil
	},
}
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := store.Load()
		if err != nil {
			return err
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks yet. Add one with: tasks add \"your task\"")
			return nil
		}
		fmt.Println("filter", filter)
		for _, t := range tasks {
			status := "[ ]"
			if t.IsDone() {
				status = "[✓]"
			}
			if t.IsPending() {
				status = "[...]"
			}
			if t.IsDone() && filter == string(task.StatusDone) {
				fmt.Printf("%s #%d %s\n", status, t.Id, t.Title)
			}
			if t.IsPending() && filter == string(task.StatusPending) {
				fmt.Printf("%s #%d %s\n", status, t.Id, t.Title)
			}
			if filter == "" {
				fmt.Printf("%s #%d %s\n", status, t.Id, t.Title)
			}

		}
		return nil
	},
}
var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := store.Load()
		if err != nil {
			return err
		}
		id := 0
		fmt.Sscanf(args[0], "%d", &id)
		found := false
		for i, t := range tasks {
			if t.Id == id {
				tasks[i] = t.MarkDone()
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("task #%d not found", id)
		}
		if err := store.Save(tasks); err != nil {
			return err
		}
		fmt.Printf("✅ Task #%d marked as done\n", id)
		return nil
	},
}
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := store.Load()
		if err != nil {
			return err
		}
		id := 0
		fmt.Sscanf(args[0], "%d", &id)
		newTasks := tasks[:0] // reuse the backing array — memory efficient
		for _, t := range tasks {
			if t.Id != id {
				newTasks = append(newTasks, t)
			}
		}
		if len(newTasks) == len(tasks) {
			return fmt.Errorf("task #%d not found", id)
		}
		if err := store.Save(newTasks); err != nil {
			return err
		}
		fmt.Printf("🗑️  Task #%d deleted\n", id)
		return nil
	},
}
