package main

import (
	"bufio"
	"fmt"
	"os"
	// "reflect"
)
func main() {
	fmt.Print("Please enter your name: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    // fmt.Println(reflect.TypeOf(input.Text()))

    fmt.Printf("Hello, %s!\n\n", scanner.Text())

    fmt.Println("Press the Enter key to exit.")
    os.Stdin.Read([]byte{0})
}
