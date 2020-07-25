package main

import (
	"github.com/pseudomind/neuik-go/neuik"
)

func cbButtonPushed(args []interface{}) {
	var mw *neuik.Window

	/* Typecast the first argument (mw *neuik.Window) */
	switch arg := args[0].(type) {
	case *neuik.Window:
		mw = arg
	}

	mw.SetTitle("The Button was pushed!")
}

func main() {
	var (
		mw  *neuik.Window
		btn *neuik.Button
	)

	if neuik.Init() {
		goto out
	}

	mw, _ = neuik.NewWindow()
	mw.SetTitle("This window has a button.")

	btn, _ = neuik.MakeButton("Push Me!")
	neuik.Element_SetCallback(btn, "OnClicked", cbButtonPushed, mw)

	mw.SetElement(btn)
	mw.Create()

	if !neuik.HasErrors() {
		neuik.EventLoop(true)
	}
out:
	if neuik.HasErrors() {
		neuik.BacktraceErrors()
	}

	neuik.Quit()
}
