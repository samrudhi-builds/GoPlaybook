# Go

Go programs demonstrating fundamental programming concepts, standard library usage, and concurrency patterns.

## Basics

**basics_demo.go** - Basic syntax: print statements, variable declarations, multiple variable initialization

**hello_world.go** - Classic "Hello, world!" program

**string_pattern_checker.go** - Checks if input starts with 'i', contains 'a', ends with 'n' (case-insensitive)

**json_marshal_demo.go** - Converts Go map to JSON string

**read_names_from_file.go** - Reads name pairs from file, displays using structs

**interactive_sorted_integers.go** - Interactive sorted slice: add integers, auto-sorts after each entry, 'X' to exit

**float_to_int_truncate.go** - Type conversion demo: float to int truncation

**word_frequency_analyzer.go** - CLI tool to count word frequency in text or files using maps
- Usage: `go run word_frequency_analyzer.go -s "text here"` or `-ss filename.txt`

## Concurrency (Goroutines & Channels)

**goroutine_basics.go** - Introduction to goroutines: sequential vs concurrent execution

**channel_basics.go** - Channel fundamentals: communication between goroutines, buffered channels, select statements

**waitgroup_mutex.go** - Synchronization primitives: WaitGroups for coordination, Mutexes for protecting shared data

**concurrent_url_checker.go** - Real-world example: checking multiple URLs concurrently, worker pool pattern

## Usage

Run any program:
```bash
go run <filename.go>
```

Build executable:
```bash
go build <filename.go>
```

## Learning Path

1. Start with basics (hello_world.go → basics_demo.go)
2. Learn data structures (slices, maps, structs)
3. Master concurrency (goroutines → channels → sync primitives)
4. Practice with real examples (url_checker)

Requires Go 1.13+