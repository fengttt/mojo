package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fengttt/mojo/pkg/mojo"
)

const (
	mjNone = 0
	mjEcho = 1
	mjDraw = 2
	mjBoth = 3

	mjstCmd          = 0
	mjstInData       = 1
	mjstDataReady    = 2
	mjstInExplain    = 3
	mjstExplainReady = 4
)

var databuf []string

func main() {
	var mojoMode int = mjEcho
	var mojoState int = mjstCmd
	// var mojoTmp string = "/tmp"

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		mojoState = handleLine(mojoState, text)
		if mojoMode&mjEcho != 0 {
			fmt.Println(text)
		}
		if mojoMode&mjDraw != 0 {
			if mojoState == mjstDataReady {
				mojo.PlotData(databuf)
			} else if mojoState == mjstExplainReady {
				mojo.ExplainPlan(databuf)
			}
		}
	}
}

func handleLine(st int, line string) int {
	return mjstCmd
}
