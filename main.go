package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	"github.com/Flaque/freak/cmds"
	"github.com/Flaque/freak/fs"
)

const help = `Usage:
$ freak <0-10>
$ freak 1 
$ freak 10
`

const freakfile = ".freak.csv"
const freakdevfile = ".freak.dev.csv"

func main() {

	// Parse arguments
	if len(os.Args) < 2 {
		fmt.Print(help)
		os.Exit(1)
	}
	arg := os.Args[1]

	// Support dev environments
	f := freakfile
	if os.Getenv("FREAK_DEV") == "TRUE" {
		fmt.Println("freak dev")
		f = freakdevfile
	}

	freakpath := path.Join(fs.HomeDir(), f)

	// Handle non number arguments
	intensity, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal("intensity must be an integer number")
	}

	err = cmds.Record(intensity, freakpath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Recorded '%d' \n", intensity)
}
