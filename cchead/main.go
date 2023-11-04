package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func start() (int, int) {
	numLines := flag.Int("n", 0, "number of lines")
	numBytes := flag.Int("c", 0, "number of bytes")
	flag.Parse()

	// if no flags are provided, perform default behavior of 10 lines
	if flag.NFlag() == 0 {
		*numLines = 10
	}

	// display error and exit if -n and -c flags are used simultaneously
	if flag.NFlag() > 1 {
		fmt.Println("cchead: can't combine line and byte counts")
		os.Exit(1)
	}

	// display error and exit if a negative number is given for bytes
	if *numBytes < 0 {
		fmt.Println("cchead: illegal byte count --", *numBytes)
		os.Exit(1)
	}

	// display error and exit if a negative number is given for lines
	if *numLines < 0 {
		fmt.Println("cchead: illegal line count --", *numLines)
		os.Exit(1)
	}

	// return lines and bytes as ints for ease of use
	return *numLines, *numBytes
}

func printBytes(reader *bufio.Reader, numBytes int) {
	if numBytes > 0 {
		content, _ := reader.Peek(numBytes)
		fmt.Printf("%s\n", content)
	}
}

func printLines(reader *bufio.Reader, numLines int) {
	if numLines > 0 {
		for counter := 0; counter < numLines; counter++ {
			content, _ := reader.ReadString('\n')
			if content != "" {
				fmt.Printf("%s", content)
			}
		}
	}
}

func main() {
	numLines, numBytes := start()

	// file path(s) must be given as the last argument(s)
	// every execution will have one arg
	numFiles := len(flag.Args())

	// if no file is provided, echo back user input
	if numFiles == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for i := numLines; i > 0; i-- {
			if scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
			}
		}
		os.Exit(0)
	}

	for i := 0; i < numFiles; i++ {
		filePath := flag.Args()[i]
		if numFiles > 1 {
			fmt.Printf("==> %s <==\n", filePath)
		}
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Could not open file --", filePath)
			os.Exit(1)
		}

		reader := bufio.NewReader(file)
		printBytes(reader, numBytes)
		printLines(reader, numLines)

		/* this is dumb, but it makes output match the legit `head` utility
		if i < numFiles {
			fmt.Println()
		}*/
	}
}
