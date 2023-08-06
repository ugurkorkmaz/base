package handler

import (
	"context"
	"log"

	pb "app/grpc/todo"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Todo struct implements gRPC methods for the TodoService server
type TodoServer struct {
	pb.TodoServiceServer
	todos []*pb.Task
}

// AddTask is a gRPC method to add a task
func (ts *TodoServer) AddTask(ctx context.Context, req *pb.AddTaskRequest) (*pb.Task, error) {
	// Create a new task and add it to the list
	todo := &pb.Task{
		Id:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
		CreateDate:  timestamppb.Now(),
		UpdateDate:  &timestamppb.Timestamp{},
	}

	ts.todos = append(ts.todos, todo)

	log.Printf("Adding task with title: %s\n", req.Title)

	// Return the added task as it is for demonstration purposes
	return todo, nil
}

// GetTask is a gRPC method to get a task
func (ts *TodoServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	// Get the task and return it to the client
	for _, todo := range ts.todos {
		if todo.Id == req.Id {
			return &pb.Task{
				Id:          todo.Id,
				Title:       todo.Title,
				Description: todo.Description,
				Completed:   todo.Completed,
				CreateDate:  todo.CreateDate,
				UpdateDate:  todo.UpdateDate,
			}, nil
		}
	}

	// Return an error if the task is not found
	return nil, status.Errorf(codes.NotFound, "Task not found")
}

// UpdateTask is a gRPC method to update a task
func (ts *TodoServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	// Update the task
	for _, todo := range ts.todos {
		if todo.Id == req.Id {
			todo.Title = req.Title
			todo.Description = req.Description
			todo.Completed = req.Completed
			todo.UpdateDate = timestamppb.Now()
			log.Printf("Updating task with id: %s\n", req.Id)

			// Return the updated task as it is for demonstration purposes
			return &pb.UpdateTaskResponse{
				Status: true,
			}, nil
		}
	}

	// Return an error if the task is not found
	return nil, status.Errorf(codes.NotFound, "Task not found")
}

// DeleteTask is a gRPC method to delete a task
func (ts *TodoServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	// Delete the task
	for i, todo := range ts.todos {
		if todo.Id == req.Id {
			ts.todos = append(ts.todos[:i], ts.todos[i+1:]...)
			log.Printf("Deleting task with id: %s\n", req.Id)

			// Return the deleted task for demonstration purposes
			return &pb.DeleteTaskResponse{
				Status: true,
			}, nil
		}
	}

	// Return an error if the task is not found
	return nil, status.Errorf(codes.NotFound, "Task not found")
}

// ListTasks is a gRPC method to list all tasks
func (ts *TodoServer) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.TaskList, error) {
	// Return all tasks
	log.Printf("Listing all tasks")

	var tasks []*pb.Task
	for _, todo := range ts.todos {
		tasks = append(tasks, &pb.Task{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreateDate:  todo.CreateDate,
			UpdateDate:  todo.UpdateDate,
		})
	}

	return &pb.TaskList{Tasks: tasks}, nil
}
