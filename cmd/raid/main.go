package main

import (
	"fmt"
	"os"
	"time"

	"github.com/falentio/raid-go"
)


func main() {
	start := time.Now()
	raids := make([]raid.Raid, 1000)
	for i := 0; i < 1000; i++ {
		raids[i] = raid.NewRaid().WithMessage(2131)
	}
	finishCreate := time.Now()
	for i := 0; i < 1000; i++ {
		fmt.Println(raids[i].String())
	}
	finishPrint := time.Now()
	fmt.Fprintf(os.Stderr, "1000 raid generated in %d ms\n", finishCreate.Sub(start).Milliseconds())
	fmt.Fprintf(os.Stderr, "1000 raid printed in %d ms\n", finishPrint.Sub(finishCreate).Milliseconds())
}