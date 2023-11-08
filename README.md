
# mywc : A Simple Word Count Utility in Go

## Overview

`mywc` is a simple command-line utility written in Go that counts lines, words, and bytes in a text file. It provides similar functionality to the `wc` command in Unix/Linux.

## Usage



$ go run mywc.go [OPTIONS] <file_path>


- `[OPTIONS]`:
  - `-l`: Count lines.
  - `-w`: Count words.
  - `-c`: Count bytes.

- `<file_path>`: The path to the file you want to count.

