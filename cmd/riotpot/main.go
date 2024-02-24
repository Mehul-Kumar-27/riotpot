// Main Application
package main

import (
	"fmt"

	"github.com/riotpot/cmd/riotpot/parse"
)

// `main` starts all the submodules containing the emulator services.
// It is the first function called when the application is run.
// It also acts as an orchestrator, which dictates the functioning of the application.
func main() {
	// Say Hi, don't be rude!
	fmt.Println("░▒▓███ RIoIPot ███▓▒░")

	// Parse the flags
	cmd := parse.NewRiotpotCommand()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
