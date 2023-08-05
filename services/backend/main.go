package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "app/grpc/todo"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Todo struct is used to hold task information
type Todo struct {
	Id          string
	Title       string
	Description string
	Completed   bool
	DueDate     *timestamp.Timestamp
}

// Server struct implements gRPC methods for the TodoService server
type Server struct {
	pb.TodoServiceServer
	todos []*Todo // List of tasks
}

// AddTask is a gRPC method to add a task
func (s *Server) AddTask(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	// Create a new task and add it to the list
	todo := &Todo{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Completed:   req.Completed,
		DueDate: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   0,
		},
	}

	s.todos = append(s.todos, todo)

	log.Printf("Adding task with title: %s\n", req.Title)

	// Return the added task as it is for demonstration purposes
	return req, nil
}

// GetTask is a gRPC method to get a task
func (s *Server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	// Get the task and return it to the client
	for _, todo := range s.todos {
		if todo.Id == req.Id {
			return &pb.Task{
				Id:          todo.Id,
				Title:       todo.Title,
				Description: todo.Description,
				Completed:   todo.Completed,
				DueDate:     todo.DueDate,
			}, nil
		}
	}

	// Return an error if the task is not found

	return nil, status.Errorf(codes.NotFound, "Task not found")
}

// UpdateTask is a gRPC method to update a task
func (s *Server) UpdateTask(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	// Update the task
	for _, todo := range s.todos {
		if todo.Id == req.Id {
			todo.Title = req.Title
			todo.Description = req.Description
			todo.Completed = req.Completed
			todo.DueDate = &timestamp.Timestamp{
				Seconds: time.Now().Unix(),
				Nanos:   0,
			}
			log.Printf("Updating task with id: %s\n", req.Id)

			// Return the updated task as it is for demonstration purposes
			return req, nil
		}
	}

	// Return an error if the task is not found
	return nil, status.Errorf(codes.NotFound, "Task not found")
}

// DeleteTask is a gRPC method to delete a task
func (s *Server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.Task, error) {
	// Delete the task
	for i, todo := range s.todos {
		if todo.Id == req.Id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			log.Printf("Deleting task with id: %s\n", req.Id)

			// Return the deleted task for demonstration purposes
			return &pb.Task{
				Id:          todo.Id,
				Title:       todo.Title,
				Description: todo.Description,
				Completed:   todo.Completed,
				DueDate:     todo.DueDate,
			}, nil
		}
	}

	// Return an error if the task is not found
	return nil, status.Errorf(codes.NotFound, "Task not found")
}

// ListTasks is a gRPC method to list all tasks
func (s *Server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.TaskList, error) {
	// Return all tasks
	log.Printf("Listing all tasks")

	var tasks []*pb.Task
	for _, todo := range s.todos {
		tasks = append(tasks, &pb.Task{
			Id:          todo.Id,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			DueDate:     todo.DueDate,
		})
	}

	return &pb.TaskList{Tasks: tasks}, nil
}

func main() {
	// Start and listen to the Golang gRPC server

	// Specify the server address, for example, "localhost:50051"
	address := "localhost:50051"

	// Create a network connection to listen on TCP
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the TodoService server
	pb.RegisterTodoServiceServer(server, &Server{})

	// Start the server and listen for client requests
	log.Printf("Server listening on %s\n", address)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
