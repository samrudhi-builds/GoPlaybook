package main

import (
	"fmt"
	"time"
)

// Without goroutines - runs sequentially
func sequentialTasks() {
	fmt.Println("\n=== Sequential Execution ===")
	start := time.Now()
	
	task("Task 1", 1)
	task("Task 2", 1)
	task("Task 3", 1)
	
	fmt.Printf("Sequential took: %v\n", time.Since(start))
}

// With goroutines - runs concurrently
func concurrentTasks() {
	fmt.Println("\n=== Concurrent Execution ===")
	start := time.Now()
	
	// The 'go' keyword makes it run concurrently
	go task("Task 1", 1)
	go task("Task 2", 1)
	go task("Task 3", 1)
	
	// Wait for goroutines to complete
	// (We'll learn better ways with channels next)
	time.Sleep(2 * time.Second)
	
	fmt.Printf("Concurrent took: %v\n", time.Since(start))
}

func task(name string, seconds int) {
	fmt.Printf("%s starting...\n", name)
	time.Sleep(time.Duration(seconds) * time.Second)
	fmt.Printf("%s completed!\n", name)
}

func main() {
	fmt.Println("GOROUTINES DEMO: Sequential vs Concurrent")
	fmt.Println("==========================================")
	
	// Run tasks one after another (slow)
	sequentialTasks()
	
	// Run tasks at the same time (fast)
	concurrentTasks()
	
	fmt.Println("\nKey takeaway: Goroutines run concurrently, making programs faster!")
}
