package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

const help = `Usage:
$ freak <0-10>
$ freak 1 
$ freak 10
`

const freakfile = ".freak.csv"

func homeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func createOrGetFile(name string) (*os.File, error) {
	var file *os.File
	var err error
	if fileExists(name) {
		file, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	} else {
		file, err = os.Create(name)
	}

	if err != nil {
		return nil, err
	}

	return file, nil
}

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

	// Create freak file if not exists, otherwise read
	file, err := createOrGetFile(freakpath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := []string{time.Now().Format(time.RFC1123), strconv.Itoa(quantity)}

	// Create a new writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write csv line at the end
	err = writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Recorded '%d' \n", quantity)
}
