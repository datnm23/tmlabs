# Seach keyword in folder

This program searches for a keyword in all files within a specified directory and prints the filenames along with the lines containing the keyword. 

## Explaination

- Searches through all files in a given directory.
- Displays filenames and the lines where the keyword appears, along with line numbers.
- Case-insensitive keyword search.

## Approach 


- Using `fmt.Scan(folderName)` to reads user input and stores it in folderName. Any errors during input are handled by log.Fatal
- `os.ReadDir(folderName)`: Reads the contents of the directory specified by folderName and returns a slice of directory entries. If there’s an error reading the directory, the program logs the error and terminates.
- Check number of files in folder. if `len(files) == 0`, there is no file in folder, print message.
- `fmt.Scan(&keyWord)`: Reads the keyword input. Any errors during input are handled by log.Fatal.
- Iterates over each file in the directory.
  - Constructs the full path `filepath`.
  - `os.ReadFile(filepath)`: Reads the content of the file into data. If an error occurs, the program logs the error and terminates.
  - `strings.Contains`: Checks if the keyword exists in the file contents. `strings.ToLower` to converts the  content and keyword to lowercase for a case-insensitive search.If the keyword is found, prints the filename.
  - `strings.Split` :Splits the file content into lines and prints the line numbers and matching lines where the keyword appears.

## Usage

Run the code

```bash
go run . 
```

