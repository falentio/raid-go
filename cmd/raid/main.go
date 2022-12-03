package main

import (
	"fmt"
	"os"
	"time"

	"github.com/falentio/raid-go"
)


func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		fmt.Println(raid.NewRaid().WithMessage(2131).String())
	}
	delta := time.Now().Sub(start)
	fmt.Fprintf(os.Stderr, "1000 raid generated in %d ms\n", delta.Milliseconds())
}