# Go Reloaded

A command-line text processing program written in Go.

Go Reloaded reads text from an input file, applies a set of formatting and transformation rules, then writes the processed result to an output file.

---

## Table of Contents

* [Overview](#overview)
* [Features](#features)
* [Usage](#usage)
* [Examples](#examples)
* [Supported Modifiers](#supported-modifiers)
* [Text Formatting Rules](#text-formatting-rules)
* [Project Structure](#project-structure)
* [Implementation Details](#implementation-details)
* [Testing](#testing)
* [Error Handling](#error-handling)
* [Built With](#built-with)

---

## Overview

This project processes plain text using inline modifiers.

A modifier is a special instruction written inside the text. It changes the word or words that come before it.

Example:

```text
hello (up)
```

Output:

```text
HELLO
```

The program supports case conversion, number base conversion, punctuation cleanup, quote formatting, and article correction.

---

## Features

* Read input from a text file
* Write processed text to an output file
* Convert hexadecimal values to decimal
* Convert binary values to decimal
* Apply uppercase, lowercase, and capitalization modifiers
* Support numbered modifiers such as `(up, 2)`, `(low, 3)`, and `(cap, 4)`
* Fix spacing around punctuation marks
* Fix spacing around single quotes
* Replace `a` with `an` when needed
* Handle invalid input safely without crashing
* Include unit tests for core text-processing behavior

---

## Usage

Run the program with:

```bash
go run . input.txt output.txt
```

Example:

```bash
go run . sample.txt result.txt
```

The program expects exactly two arguments:

```text
input.txt   -> the file to read from
output.txt  -> the file to write the processed result into
```

If the wrong number of arguments is provided, the program prints:

```text
Usage: go run . input.txt output.txt
```

---

## Examples

### Example Input

```text
it was a amazing day . i saw 1E (hex) birds and 10 (bin) cats . this is great (up, 2)
```

### Example Output

```text
it was an amazing day. i saw 30 birds and 2 cats. this IS GREAT
```

---

## Supported Modifiers

Modifiers are removed from the final output after they are applied.

### Uppercase Modifier

Converts the previous word to uppercase.

```text
hello (up)
```

Output:

```text
HELLO
```

---

### Lowercase Modifier

Converts the previous word to lowercase.

```text
HELLO (low)
```

Output:

```text
hello
```

---

### Capitalize Modifier

Capitalizes the previous word.

```text
brooklyn (cap)
```

Output:

```text
Brooklyn
```

---

### Hexadecimal Modifier

Converts the previous hexadecimal value to decimal.

```text
1E (hex)
```

Output:

```text
30
```

---

### Binary Modifier

Converts the previous binary value to decimal.

```text
10 (bin)
```

Output:

```text
2
```

---

## Numbered Modifiers

Numbered modifiers apply a transformation to multiple previous words.

The supported numbered modifiers are:

```text
(up, n)
(low, n)
(cap, n)
```

Where `n` is the number of previous words to modify.

---

### Uppercase Multiple Words

```text
this is great (up, 2)
```

Output:

```text
this IS GREAT
```

---

### Lowercase Multiple Words

```text
THIS IS GREAT (low, 2)
```

Output:

```text
THIS is great
```

---

### Capitalize Multiple Words

```text
hello from brooklyn (cap, 2)
```

Output:

```text
hello From Brooklyn
```

---

## Text Formatting Rules

### Punctuation

The program removes unnecessary spaces before punctuation marks.

Supported punctuation marks:

```text
. , ! ? : ;
```

Example:

```text
Hello , world !
```

Output:

```text
Hello, world!
```

The program also handles punctuation attached to the beginning of a word.

Example:

```text
Hello ,world
```

Output:

```text
Hello, world
```

---

### Single Quotes

The program fixes spacing inside single quotes.

Example:

```text
say ' hello ' to me
```

Output:

```text
say 'hello' to me
```

Opening quotes are attached to the next word, and closing quotes are attached to the previous word.

---

### Article Correction

The program changes `a` to `an` before words that start with a vowel or `h`.

Example:

```text
a apple
```

Output:

```text
an apple
```

Example:

```text
A orange
```

Output:

```text
An orange
```

Checked starting letters:

```text
a e i o u h
```

---

## Project Structure

```text
.
├── go.mod
├── main.go
├── main_test.go
├── input.txt
└── output.txt
```

---

## Implementation Details

The program follows a clear text-processing pipeline:

```text
Read input file
      ↓
Split text into words
      ↓
Apply modifiers
      ↓
Fix punctuation
      ↓
Fix quotes
      ↓
Fix articles
      ↓
Join words back into text
      ↓
Write output file
```

The main pipeline is handled by:

```go
func processText(text string) string
```

---

## Main Components

### File Handling

```go
func readFile(filename string) (string, error)
func writeFile(filename string, content string) error
```

These functions handle reading from the input file and writing to the output file.

---

### Modifier Processing

```go
func processModifiers(words []string) []string
```

This function detects and applies all supported modifiers:

* `(up)`
* `(low)`
* `(cap)`
* `(hex)`
* `(bin)`
* `(up, n)`
* `(low, n)`
* `(cap, n)`

---

### Number Conversion

```go
func hexToDecimal(s string) string
func binToDecimal(s string) string
```

These functions convert hexadecimal and binary values into decimal strings.

If conversion fails, the original value is returned unchanged.

---

### Text Formatting

```go
func fixPunctuation(words []string) []string
func fixQuotes(words []string) []string
func fixArticles(words []string) []string
```

These functions handle punctuation spacing, quote formatting, and article correction.

---

## Testing

Unit tests are included in:

```text
main_test.go
```

Run all tests:

```bash
go test
```

Run tests with detailed output:

```bash
go test -v
```

The tests cover:

* Hexadecimal conversion
* Binary conversion
* Single-word modifiers
* Numbered modifiers
* Punctuation spacing
* Single quote formatting
* Article correction

---

## Error Handling

The program handles the following cases:

* Incorrect number of command-line arguments
* Failure to read the input file
* Failure to write to the output file
* Invalid hexadecimal values
* Invalid binary values
* Invalid numbered modifier values
* Modifiers appearing without previous words

When a value cannot be converted, the original word is kept unchanged.

---

## Built With

* Go
* `fmt`
* `os`
* `strconv`
* `strings`
* Go testing package

---

## Purpose

This project was built as part of the Go Reloaded exercise.

It focuses on file handling, string manipulation, parsing, number conversion, edge case handling, and building a clean text-processing pipeline in Go.
