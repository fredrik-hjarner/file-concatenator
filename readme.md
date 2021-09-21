# File Concatinator

A util program to help me print a code repo.

Concatenates all code files in a repo (or any folder really) into one text file. It adds some
`form feed` to make sure max one file per paper.

This program does not print it for you, it creates a text file. You have to print it yourself, so
you don't have to be afraid to be completely replaced by a machine, yet...

## Assumptions

- 100 characters fit per line.
- 75 lines fit on one page.

2 sheets per paper is nice, too.

### If printing with WordPad

Courier New 8px

margin-left: 30mm

margin-right: 5mm

margin-top: 15mm

margin-bottom: 5mm

## Install & use

Just do `go install`

then you'll be able to write `file-concatinator` in terminal while in a directory structure that
contains code!

So why don't you go right now to the computer root folder and print ALL code on your entire
computer?!

## Testing

### Run all tests

`go test ./... -cover -count 1 -test.v`
