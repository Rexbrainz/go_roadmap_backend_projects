package main

import (
	"fmt"
	"os"

	"github.com/Rexbrainz/task-tracker/tracker"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Wrong number of inputs")
		os.Exit(0)
	}

	tracker.Track()
}