# Print tree directory

This program will print lists the structure of a directory

## Approach

The program employs a recursive approach to traverse and print the directory structure.

- It verifies whether the specified directory path exists using `os.Open` If the directory does not exist, the program prints an error message and terminates.
- For each directory, the program lists all entries (files, subdirectories).If an entry is a subdirectory, the function calls itself recursively to process the contents of that subdirectory.If an entry is a regular file, it is printed directly.
- `printFormatting` will format the output . The depth of each directory or file in the hierarchy is represented using indentation. The `depth` parameter determines how many levels deep the current entry is and controls the amount of indentation to visually represent the directory tree. Each entry is printed with |-- to indicate the start of a new item.
- `printDirectory` initializes the recursive traversal and handles the starting point of the directory listing. And ituse a recursive function `printDirRec` to traverse directories. `printDirRec` responsible for traversing directories and printing their contents. Sart the Process: `printDirectory` is called with the initial path and a depth of `0`. In recursive, use `os.ReadDir` retrieves all entries in the current directory and printed with appropriate indentation. Recursively call `printDirRec` for subdirectories and. Print file names with indentation based on depth. The recursion ends when all entries in the current directory and its subdirectories have been processed. Errors encountered while reading directories are handled using log.Fatal, which prints the error message and terminates the program.

## Usage

Run the code

```bash
go run . <directory_path>
```
<directory_path>: The path to the directory you want to list.

## Result
```bash

|--lab08
| |--gunfile
| | |--abc.txt
| | |--def.txt
| |--hyperloop
| |--neuralchip
| |--tesla
| | |--readme.md
| |--tree.go
| |--twitter

```