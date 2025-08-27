package main

// flow of shit.
// first view should be a benchmark picker maybe, but that can come later
// second view should obviously be the benchmark table view.
// third view could be a scenario score graph view
// idk

// guess what this does
import (
	"log"
	"fmt"
	"os"
	// import bubble tea as btea because typing bubbletea is too damn long -.-
	btea "github.com/charmbracelet/bubbletea"
)

func main() {

	// debug log setup because blehhhh
	// god i do not like go syntax.
	f, err := os.Create("debug.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Hello from Bubble Tea debug!")
    }
}

