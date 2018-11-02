package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

const help = `Usage:
$ freak <0-10>
$ freak 1 
$ freak 10
`

func main() {

	// Parse arguments
	if len(os.Args) < 2 {
		fmt.Print(help)
		os.Exit(1)
	}
	arg := os.Args[1]

	freakpath := path.Join(homeDir(), freakfile)

	// Handle non number arguments
	quantity, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal("quantity must be an integer number")
	}

	data := []string{time.Now().Format(time.RFC1123), strconv.Itoa(quantity)}

	err = WriteToCSV(data, freakpath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Recorded '%d' \n", quantity)
}
