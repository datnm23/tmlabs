# Write string into file

This program will write a string to a file . If the file doesn't exists, it will create a new file then write content into the file.

**Note**

- If file is exists, the program will overide old contents.

## Explaination

To write a contents to the file, need check exists of the file. if it exists open the file and write new contents replace old contents in the file

## Approach

-`checkFileExisted`check if the file exists. using `os.Open()`  to opens the named file for reading. If successful,return `true` and we will have file can be used for reading. If there is an error, using `os.IsNotExist(err)`to check the cause of the error. If the cause is because the file does not exist, using `os.Create` create a new file. if not, print error message and stop
-`WriteStruct` write string in to the file. Using `fmt.Sprint` convert to  the default formats and convert it to []byte. Then using `os.WriteFile` to replace old content in file and write []byte to file;

## Usage

Create a text file with name: `contents.txt` or let program will create it.

Run the code

```bash
go run .
```