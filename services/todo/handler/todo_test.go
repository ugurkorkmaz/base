package handler

import (
	"context"
	"testing"

	pb "app/grpc/todo"
)

func TestTodoServer_AddTask(t *testing.T) {
	// Initialize the TodoServer
	ts := &TodoServer{}

	// Prepare a request with a new task
	req := &pb.Task{
		Id:          "task-1",
		Title:       "Buy groceries",
		Description: "Buy fruits and vegetables",
		Completed:   false,
	}

	// Call the AddTask method
	resp, err := ts.AddTask(context.Background(), req)
	if err != nil {
		t.Fatalf("AddTask failed with error: %v", err)
	}

	// Check if the response matches the request
	if resp != req {
		t.Errorf("Expected response: %v, but got: %v", req, resp)
	}

	// Check if the task was added to the todos list
	if len(ts.todos) != 1 {
		t.Errorf("Expected 1 task in the list, but got: %d", len(ts.todos))
	}
}

func TestTodoServer_GetTask(t *testing.T) {
	// Initialize the TodoServer
	ts := &TodoServer{}

	// Add a task to the list
	task := &pb.Task{
		Id:          "task-1",
		Title:       "Buy groceries",
		Description: "Buy fruits and vegetables",
		Completed:   false,
	}
	ts.todos = append(ts.todos, task)

	// Prepare a request to get the added task
	req := &pb.GetTaskRequest{Id: "task-1"}

	// Call the GetTask method
	resp, err := ts.GetTask(context.Background(), req)
	if err != nil {
		t.Fatalf("GetTask failed with error: %v", err)
	}

	// Check if the response matches the added task
	if resp != task {
		t.Errorf("Expected response: %v, but got: %v", task, resp)
	}
}

func TestTodoServer_UpdateTask(t *testing.T) {
	// Initialize the TodoServer
	ts := &TodoServer{}

	// Add a task to the list
	task := &pb.Task{
		Id:          "task-1",
		Title:       "Buy groceries",
		Description: "Buy fruits and vegetables",
		Completed:   false,
	}
	ts.todos = append(ts.todos, task)

	// Prepare a request to update the task
	req := &pb.Task{
		Id:          "task-1",
		Title:       "Buy groceries and household items",
		Description: "Buy fruits, vegetables, and cleaning supplies",
		Completed:   true,
	}

	// Call the UpdateTask method
	resp, err := ts.UpdateTask(context.Background(), req)
	if err != nil {
		t.Fatalf("UpdateTask failed with error: %v", err)
	}

	// Check if the response matches the updated task
	if resp != req {
		t.Errorf("Expected response: %v, but got: %v", req, resp)
	}

	// Check if the task was updated in the todos list
	updatedTask := ts.todos[0]
	if updatedTask.Title != req.Title || updatedTask.Description != req.Description || !updatedTask.Completed {
		t.Errorf("Task was not updated correctly")
	}
}

func TestTodoServer_DeleteTask(t *testing.T) {
	// Initialize the TodoServer
	ts := &TodoServer{}

	// Add a task to the list
	task := &pb.Task{
		Id:          "task-1",
		Title:       "Buy groceries",
		Description: "Buy fruits and vegetables",
		Completed:   false,
	}
	ts.todos = append(ts.todos, task)

	// Prepare a request to delete the task
	req := &pb.DeleteTaskRequest{Id: "task-1"}

	// Call the DeleteTask method
	resp, err := ts.DeleteTask(context.Background(), req)
	if err != nil {
		t.Fatalf("DeleteTask failed with error: %v", err)
	}

	// Check if the response matches the deleted task
	if resp != task {
		t.Errorf("Expected response: %v, but got: %v", task, resp)
	}

	// Check if the task was deleted from the todos list
	if len(ts.todos) != 0 {
		t.Errorf("Expected 0 tasks in the list after deletion, but got: %d", len(ts.todos))
	}
}

func TestTodoServer_ListTasks(t *testing.T) {
	// Initialize the TodoServer
	ts := &TodoServer{}

	// Add some tasks to the list
	task1 := &pb.Task{
		Id:          "task-1",
		Title:       "Buy groceries",
		Description: "Buy fruits and vegetables",
		Completed:   false,
	}
	task2 := &pb.Task{
		Id:          "task-2",
		Title:       "Read a book",
		Description: "Read 'The Golang Book'",
		Completed:   true,
	}
	ts.todos = append(ts.todos, task1, task2)

	// Prepare a request to list all tasks
	req := &pb.ListTasksRequest{}

	// Call the ListTasks method
	resp, err := ts.ListTasks(context.Background(), req)
	if err != nil {
		t.Fatalf("ListTasks failed with error: %v", err)
	}

	// Check if the response contains all tasks in the todos list
	if len(resp.Tasks) != 2 {
		t.Errorf("Expected 2 tasks in the response, but got: %d", len(resp.Tasks))
	}

	// Check if the tasks in the response match the added tasks
	if resp.Tasks[0] != task1 || resp.Tasks[1] != task2 {
		t.Errorf("Tasks in the response do not match the added tasks")
	}
}
