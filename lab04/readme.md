# Answer the question by input and write into file

This program will be answer the given set of questions and write the questions and answers in the file.

**Note**

- If file is exists, the program will overide old contents.

## Explaination

Enter answers from the keyboard. Then print questions and answers to the "person.txt'

## Appoach
- Creat list of questions.
- Print question to the console.
- Create new slice of string to store question and answer named _result_.
- Iterate through each question in list of question. Declares variable named _answer_ and  type string. Using `fmt.Scan` to input answer and scan to _answer_. Append question and answer to the _result_.
- using `strings.Joint` to join (by down line ) _data_ and using method `WriteStruct` packaged in lab03 to write _data_ to person.txt
## Usage

Create a text file with name: `person.txt`

Run the code

```bash
go run .
```