package main

import (
	"github.com/pseudomind/neuik-go/neuik"
)

func main() {
	var (
		mw   *neuik.Window
		fg   *neuik.FlowGroup
		btn1 *neuik.Button
		btn2 *neuik.Button
		btn3 *neuik.Button
		btn4 *neuik.Button
		btn5 *neuik.Button
		btn6 *neuik.Button
		btn7 *neuik.Button
		btn8 *neuik.Button
		btn9 *neuik.Button
	)

	if neuik.Init() {
		goto out
	}

	mw, _ = neuik.NewWindow()
	mw.SetTitle("This window has a button.")
	mw.Configure("Resizable")

	btn1, _ = neuik.MakeButton("Button #1")
	btn2, _ = neuik.MakeButton("Button #2")
	btn3, _ = neuik.MakeButton("Button #3")
	btn4, _ = neuik.MakeButton("Button #4")
	btn5, _ = neuik.MakeButton("Button #5")
	btn6, _ = neuik.MakeButton("Button #6")
	btn7, _ = neuik.MakeButton("Button #7")
	btn8, _ = neuik.MakeButton("Button #8")
	btn9, _ = neuik.MakeButton("Button #9")

	fg, _ = neuik.NewFlowGroup()
	neuik.Element_Configure(fg, "FillAll", "PadAll=1")
	neuik.Container_AddElements(fg,
		btn1, btn2, btn3, btn4, btn5, btn6, btn7, btn8, btn9)

	mw.SetElement(fg)
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
