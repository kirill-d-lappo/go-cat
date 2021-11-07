# Go-GNU: cat

## Description

cat - concatenate files and print on the standard output

## Usage:

```bash
cat [option] [file]…
```

## Description
Concatenate FILE(s) to standard output.

With no FILE, or when FILE is -, read standard input.

* `--help`

    display help and exit

* `-u`

    (ignored)

* `--version`

    output version information and exit

### Not Implemented

* `-A`, `--show-all`

    equivalent to -vET

* `-b`, `--number-nonblank`

    number nonempty output lines, overrides -n

* `-e`

    equivalent to -vE

* `-E`, `--show-ends`
*
    display $ at end of each line

* `-n`, `--number`

    number all output lines

* `-s`, `--squeeze-blank`

    suppress repeated empty output lines

* `-t`

    equivalent to -vT

* `-T`, `--show-tabs`

    display TAB characters as ^I

* `-v`, `--show-nonprinting`

    use ^ and M- notation, except for LFD and TAB

## Example

```bash
# Output f's contents, then standard input, then g's contents.
cat f - g

# Copy standard input to standard output.
cat
```