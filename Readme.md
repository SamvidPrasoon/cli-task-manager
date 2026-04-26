### Folder Structure
```
01-cli-task-manager/
├── main.go                      ← entry point, calls cmd.Execute()
├── go.mod                       ← module definition
├── cmd/
│   └── root.go                  ← all cobra commands (add, list, done, delete)
└── internal/
    ├── task/
    │   ├── task.go              ← Task type, New(), IsDone(), MarkDone()
    │   └── task_test.go         ← table-driven tests
    └── storage/
        └── storage.go           ← JSON file persistence
```
### Run It

```bash
cd 01-cli-task-manager
go run main.go add "Learn Go interfaces"
go run main.go add "Build a REST API"
go run main.go list
go run main.go done 1
go run main.go list
go run main.go delete 2
go test ./...
go run main.go list --filter=done
go run main.go list --filter=pending
```