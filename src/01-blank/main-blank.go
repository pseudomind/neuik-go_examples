package main

import (
	"github.com/pseudomind/neuik-go/neuik"
)

func main() {
	var mw *neuik.Window

	neuik.Init() /* Initialize the NEUIK library */

	mw, _ = neuik.NewWindow() /* Create a new Window */
	mw.SetTitle("Quick Demo")
	mw.SetSize(200, 100)
	mw.Create() /* Create and show the window */

	//------------------------------------------------------------------------//
	// Check for any errors which may have occurred before starting the loop  //
	//------------------------------------------------------------------------//
	if !neuik.HasErrors() {
		neuik.EventLoop(true) /* Start capturing events 1=killOnError */
	}

	//------------------------------------------------------------------------//
	// Check for any errors which may have occurred during the event loop     //
	//------------------------------------------------------------------------//
	if neuik.HasErrors() {
		neuik.BacktraceErrors()
	}

	neuik.Quit()
}
