package main

// import "bufio"
import (
	"flag"
	"fmt"
)
// import "os"
func main() {
	arguments := parse_command_line_arguments()

    fmt.Printf("Hello, %s, times %d!\n", *arguments.name, *arguments.multiplier)
    fmt.Printf("You listed the following items: %v\n", arguments.items)
    /*
    input := bufio.NewScanner(os.Stdin)
    fmt.Println(input.Text)
    */


}


func parse_command_line_arguments() Arguments {
	arguments := Arguments {
		flag.Int("multiplier", 1, "intensity multiplier"),
		flag.String("name", "unknown person", "name of the person to greet"),
		[]string{},
	}

	flag.Parse()

	arguments.items = flag.Args()

	return arguments
}


type Arguments struct {
	multiplier *int
	name *string
	items []string
}
