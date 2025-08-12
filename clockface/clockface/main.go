package main

import (
	"os"
	"time"

	"github.com/madhusudhan-reddy-oneture/gotbd/clockface/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
