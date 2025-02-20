package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var pl = fmt.Println

func main() {
	pl("What is your name??")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	name = strings.TrimSpace(name) // Removes the trailing newline
	pl("Hello", name)
}
