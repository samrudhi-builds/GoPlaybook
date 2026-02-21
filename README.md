# Go

Go programs demonstrating fundamental programming concepts and standard library usage.

## Files

**basics_demo.go** - Basic syntax: print statements, variable declarations, multiple variable initialization

**word_frequency_analyzer.go** - CLI tool to count word frequency in text or files using maps
- Usage: `go run word_frequency_analyzer.go -s "text here"` or `-ss filename.txt`

**hello_world.go** - Classic "Hello, world!" program

**string_pattern_checker.go** - Checks if input starts with 'i', contains 'a', ends with 'n' (case-insensitive)

**json_marshal_demo.go** - Converts Go map to JSON string

**read_names_from_file.go** - Reads name pairs from file, displays using structs

**interactive_sorted_integers.go** - Interactive sorted slice: add integers, auto-sorts after each entry, 'X' to exit

**float_to_int_truncate.go** - Type conversion demo: float to int truncation

## Usage

Run any program:
```bash
go run <filename.go>
```

Build executable:
```bash
go build <filename.go>
```

Requires Go 1.13+