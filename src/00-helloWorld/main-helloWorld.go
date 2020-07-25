package main

import (
	"github.com/pseudomind/neuik-go/neuik"
)

func main() {
	var (
		mw  *neuik.Window
		lbl *neuik.Label
	)

	neuik.Init() /* Initialize the NEUIK library */

	mw, _ = neuik.NewWindow() /* Create a new Window */
	mw.SetTitle("Quick Demo")
	mw.SetSize(200, 100)

	lbl, _ = neuik.MakeLabel("Hello World") /* Create a new Label */
	/* set the label as the element contained by the window */
	mw.SetElement(lbl)
	mw.Create() /* Create and show the window */

	//------------------------------------------------------------------------//
	// Check for any errors which may have occurred before starting the loop  //
	//------------------------------------------------------------------------//
	if neuik.HasErrors() {
		neuik.BacktraceErrors()
		return
	}
	neuik.EventLoop(true) /* Start capturing events 1=killOnError */

	//------------------------------------------------------------------------//
	// Check for any errors which may have occurred during the event loop     //
	//------------------------------------------------------------------------//
	if neuik.HasErrors() {
		neuik.BacktraceErrors()
		return
	}

	neuik.Quit()
}
