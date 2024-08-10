# Print file content program

This program read content from a text file and print it to the stdout.

## Explaination

The lab request reads an existing file and prints its contents to console

## Approach

- Create a file .txt as required. This file contains only plain text contents.
- Use os.ReadFile() method to read a file by using the path as its parameter. This method returns either the data of the file or an error.The data will be a slice of byte []byte
so we will convert to string and print to console.
- If there was an error while reading the file, the app will print error message and stop.

## Usage

Create a text file with name: `contents.txt`

Run the code

```bash
go run .
```
