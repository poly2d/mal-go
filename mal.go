package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read(in string) string {
	return in
}

func eval(in string) string {
	return in
}

func print(in string) string {
	return in
}

func rep(in string) string {
	readRes := read(in)
	evalRes := eval(readRes)
	return print(evalRes)
}

func main() {
	const PROMPT string = "user> "
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(PROMPT)
		in, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("\ngoodbye")
			break
		}
		out := rep(in)
		fmt.Println(out)
	}
}
