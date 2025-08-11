package main

import (
	"flag"
	"fmt"
)

func printOutput(key, message string) {
	fmt.Printf("::set-output name=%s::%s\n", key, message)
}

func main() {
	var msg = flag.String("message", "Hello From Go GithHub Action", "action -message <your message>")

	flag.Parse()

	fmt.Println(*msg)

	printOutput("message", *msg)
}
