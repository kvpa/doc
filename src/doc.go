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
	nodes, err := os.ReadDir(path)

	if err != nil {
		strerr := "doc: Can't read contents of '" + path + "'.\n"
		os.Stderr.WriteString(strerr)
		os.Exit(1)
	}

	for i := 0; i < len(nodes); i++ {
		nodePath := path + "/" + nodes[i].Name()

		if nodes[i].IsDir() {
			showFolderDoc(nodePath)
		} else {
			showFileDoc(nodePath)
		}

		if i != len(nodes)-1 {
			os.Stdout.WriteString("\n\n")
		}
	}
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
