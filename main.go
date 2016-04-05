package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("calc > ")

	for scanner.Scan() {
		line := scanner.Text()

		interpreter := NewInterpreter(line)
		value := interpreter.interpret()

		fmt.Println(value)

		fmt.Print("calc > ")
	}
}
