# Replace sensitive words and Print file content program

This program reads contents from a text file, and if contents have sensitive words like sex, fuck, drug, kill, replace them with s\*x, f\*ck, dr\*g, and k\*ll. After that, print it to the stdout.

**Note**

- A word should not started with any special characters.
- If a word end with a character like `.`, `,`, `!`, it may be the last word of a sentence, so it still valid.

## Explaination

The lab request reads an existing file and prints its contents to the console.
If the content of the file has some sentitive words, we need to change them by replacing the vowel of a sensitive word with *. Ex : sex - s\*x, fuck - f\*ck,...
To check sensitive words in their contents, we need to check word by word in the contents to avoid replacing part of the another word that contains a sensitive word. 

## Approach
1. First, read a file like lab01. We will have the contents of the file as a slice of byte. 
2. A word is a slice of character between a prefix and a suffix. The prefixs may be a blank space,comma,dot. And the suffixs are a blank space, comma, dot,exclamation mark (!).
So from a slice of byte, we will convert to a string. Then using `strings.Split` to split (by space) the content string into a slice of word. However, each word may has a prefix or suffix if it is adjacent to the word. Therefore, when we check the sensitive word, it will not match, leading to missing the sensitive word.
3. `checkPrefix`_`,_checkSuffix_ will discover a word with prefix and suffix. Return _true_ if contains prefix and suffix, and return _false_ if it does not. If _true_,split the prefix and suffix, we will have a word to the next step.
4. Using `strings.ToLower(word)` to convert the word to lower case to ensure it matches. And using _strings.Compare_ to check if a word is sensitive.
5. If it is a sensitive word, take all vowels from the  word and using `strings.ReplaceAll` replace its with *.
6. Make a new slices of word and append a word that has been checked and consored to the slice .
7. Print slice of words to the console


## Usage

Create a text file with name: `contents.txt`

Run the code

```bash
go run .
```
## Testing

```bash
go test
```

Result

```bash
=== RUN   TestReplaceAllSensitiveWords
    main_test.go:47: Testing case: 00
    main_test.go:47: Testing case: 01
    main_test.go:47: Testing case: 02
    main_test.go:47: Testing case: 03
    main_test.go:47: Testing case: 04
    main_test.go:47: Testing case: 05
    main_test.go:47: Testing case: 06
    main_test.go:47: Testing case: 07
--- PASS: TestReplaceAllSensitiveWords (0.00s)
PASS
ok      github.com/datnm23/tmlabs/lab02 0.002s
```