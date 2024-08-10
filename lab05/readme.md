# Create command line app to making file having input size 
This program will write a command line application to make a file include the random character  "A-Za-z0-9" having a size X megabytes and input from the keyboard

## Explaination

Input filename fileName and an integer X, need to create a file **fileName.txt** having size is x megabytes . One character equals 1 byte so need x \* 1024 \* 1024 characters in **fileName.txt**

## Approach

- Using `fmt.Scan` to input fileName and an integer X, Create new text file named "filename.txt".
- __Maxsize__ = x \*1024 \*1024 byte, and one line have 256 character. Because down line use "\n" equals 1 by so we have number of lines have 256 character is _totalLinesFullCharacter_ = __Maxsize__ / 257 and _lastLine_ is _numCharInLastLine_ = __Maxsize__ % 257.
- create _contents_ is slice of string and each element of slice will be a content of line.
- _lineContent_ is content in a line. One line have 256 characters so iterate with 256 times, each time using `randCharacter()` choose a random chacrater, join to _lineContent_ by operator `+`. In last step, _lineContent_ will be contents of line with 256 character then append _lineContent_ to slice  _contents_. Repeat this step in _totalLinesFullCharacter_ times.
- _lastLineContent_-_ is content in the last line, we have _numCharInLastLine_ character so we will iterate _numCharInLastLine_time , each time using `randCharacter()` choose a random chacrater, join  to _lastLineContent_ by operator + .Done this step, append _lastLineContent_ to slice _contents_
- Join slice _contents_ by down line ("\n"). Then using `WriteStruct` to write contents to file.

## Usage

Run the code

```bash
go run .
```

## Problem

- If file is exists, the program will overide old contents.
- Because, x is unsigned integer so x cann't bigger than the max values that are defined in the math package. In this case x <= 18446744073709551615;
- if x in [0,18446744073709551615] and x very big this is lead to filename.txt very big make full hard drive capacity. so can't write into file.
