package main

import (
	"io"
	"os"
	"strings"
)

func showFileDoc(path string) {
	file, err := os.Open(path)

	if err != nil {
		strerr := "doc: Can't open '" + path + "'.\n"
		os.Stderr.WriteString(strerr)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, file)

	if err != nil {
		strerr := "doc: Failed to read from '" + path + "'.\n"
		os.Stderr.WriteString(strerr)
		os.Exit(1)
	}
}

func showFolderDoc(path string) {
	// Send all documentation from a folder to the standard output.
}

func main() {
	if len(os.Args) == 1 {
		showFolderDoc("/doc/")
	} else if len(os.Args) == 2 {
		var topic string

		if strings.HasPrefix(os.Args[1], "./") {
			topic = os.Args[1]
		} else if strings.HasPrefix(os.Args[1], "../") {
			topic = os.Args[1]
		} else if strings.HasPrefix(os.Args[1], "/") {
			topic = os.Args[1]
		} else {
			topic = "/doc/" + os.Args[1]
		}

		node, err := os.Stat(topic)

		if err != nil {
			strerr := "doc: Can't access '" + os.Args[1] + "'.\n"
			os.Stderr.WriteString(strerr)
			os.Exit(1)
		}

		if node.IsDir() {
			showFolderDoc(topic)
		} else {
			showFileDoc(topic)
		}
	} else {
		os.Stderr.WriteString("doc: Invalid invocation.\n")
		os.Exit(1)
	}
}
