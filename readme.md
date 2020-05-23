# Go-GNU: cat

## Description

`cat` copies each file (‘-’ means standard input), or standard input if none are given, to standard output.

## Usage:

```bash
cat [option] [file]…
```

## Example

```bash
# Output f's contents, then standard input, then g's contents.
cat f - g

# Copy standard input to standard output.
cat
```