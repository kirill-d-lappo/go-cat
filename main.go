package main

func main() {
	config := NewCatConfig()

	if config.showVersion {
		printVersion()
		return
	}

	RunCat(config)
}

func printVersion() {
	versionText := `cat (GO-CAT project) 0.1.0
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Original cat is written by Torbjorn Granlund and Richard M. Stallman.

go version is written by Kirill Lappo as a part of go-cat project.

just for fun
`

	print(versionText)
}
