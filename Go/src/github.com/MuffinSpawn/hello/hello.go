package main

// import "bufio"
import (
	"flag"
	"fmt"
	"io"
	"os"
)
// import "os"
func main() {
	arguments := parse_command_line_arguments()
	if len(arguments.files) < 2 {
		fmt.Printf("Source and destination files not given.\n")
	} else {
	    fmt.Printf("Hello, %s, times %d!\n", *arguments.name, *arguments.multiplier)
	    fmt.Printf("Source/Destination file: %v/%v\n", arguments.files[0], arguments.files[1])
	    /*
	    input := bufio.NewScanner(os.Stdin)
	    fmt.Println(input.Text)
	    */

	    _, err := Copy(arguments.files[0], arguments.files[1])
	    if err != nil {
	    	fmt.Printf("Copy failed. %v\n", err)
	    }
	}
}


func parse_command_line_arguments() Arguments {
	arguments := Arguments {
		flag.Int("multiplier", 1, "intensity multiplier"),
		flag.String("name", "unknown person", "name of the person to greet"),
		[]string{},
	}

	flag.Parse()

	arguments.files = flag.Args()

	return arguments
}


func Copy(src, dst string) (uint64, error) {
    var bytes_copied int64 = 0
    var is_regular_file bool
    var err error
    var source *os.File
    var destination *os.File

    is_regular_file, err = IsRegularFile(src)

    if is_regular_file {
	    source, err = os.Open(src)
    }

    if err == nil {
	    defer source.Close()
	    is_regular_file, err = IsRegularFile(dst)
    }

    if is_regular_file {
    	err = fmt.Errorf("File %s already exists", dst)
    } else {
	    destination, err = os.Create(dst)
    }

    if err == nil {
	    defer destination.Close()
	    bytes_copied, err = io.Copy(destination, source)
    }

    return uint64(bytes_copied), err
}


func IsRegularFile(file string) (bool, error) {
	is_regular_file :=  false
	var err error = nil
    source_file_stats, err := os.Stat(file)

    if err == nil {
	    is_regular_file = source_file_stats.Mode().IsRegular()
    } else {
    	err = fmt.Errorf("%s is not a regular file", file)
    }

    return is_regular_file, err
}


type Arguments struct {
	multiplier *int
	name *string
	files []string
}
