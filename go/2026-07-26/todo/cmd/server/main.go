package main

import (
	"context"
	"log"
	"time"

	"todo/internal/http"
	"todo/internal/store/entity"
)

func main() {
	server := http.NewServer()

	// Register sample data (services are stubs, so seed via stores directly)
	ctx := context.Background()
	now := time.Now()

	goal := &entity.Goal{
		ID:          "goal-1",
		Title:       "Learn Go",
		Description: "Build a layered web app with Gin",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	server.GoalStore.SaveGoal(ctx, goal)

	server.TaskStore.SaveTask(ctx, &entity.Task{
		ID:        "task-1",
		GoalID:    goal.ID,
		Title:     "Implement the goal service",
		Done:      false,
		CreatedAt: now,
		UpdatedAt: now,
	})

	log.Println("Starting server on :8080")
	log.Println("Sample goal: goal-1")
	log.Println("Sample task: task-1")

	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
