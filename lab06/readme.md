# Create command line app to seach file line
This program will write a command line application input file name and Y. Return print y th line in filename

## Explaination

Input filename fileName and an integer Y, need to read a file **fileName.txt**. And get line Yth in file then print to console 

## Approach

- Using `fmt.Scan` to input fileName and an integer Y, Read text file named "filename.txt".The data will be a slice of byte named _contents_
so we will convert them to string . If there was an error while reading the file, the app will print error message and stop.

- Split
## Usage

Run the code

```bash
go run .
```

## Problem

- If file is exists, the program will overide old contents.
- Because, x is unsigned integer so x cann't bigger than the max values that are defined in the math package. In this case x <= 18446744073709551615;
- if x in [0,18446744073709551615] and x very big this is lead to filename.txt very big make full hard drive capacity. so can't write into file.
