package main

import (
	"github.com/pseudomind/neuik-go/neuik"
	"errors"
)

func cbSetBGColor(args []interface{}) {
	var (
		mw     *neuik.Window // arg1
		clrStr string        // arg2
	)

	if len(args) < 2 {
		panic(errors.New("This function requires two arguments."))
	}

	/* Typecast the first argument (mw *neuik.Window) */
	switch arg := args[0].(type) {
	case *neuik.Window:
		mw = arg
	default:
		panic(errors.New("First argument `mw *neuik.Window` is not of the right type"))
	}

	/* Typecast the second argument (clrStr string) */
	switch arg := args[1].(type) {
	case string:
		clrStr = arg
	default:
		panic(errors.New("2nd argument `clrStr string` is not of the right type"))
	}

	mw.Configure("BGColor=" + clrStr)
}

func cbButton_PushMe(args []interface{}) {
	var mw *neuik.Window

	/* Typecast the first argument (mw *neuik.Window) */
	switch arg := args[0].(type) {
	case *neuik.Window:
		mw = arg
	}
	print("Button `Push Me!` pushed!!!\n")
	mw.SetTitle("The Button was pushed!")
}

func main() {
	var (
		mw   *neuik.Window
		vg   *neuik.VGroup
		btn1 *neuik.Button
		btn2 *neuik.Button
		btn3 *neuik.Button
		btn4 *neuik.Button
		btn5 *neuik.Button
	)

	if neuik.Init() {
		goto out
	}

	mw, _ = neuik.NewWindow()
	mw.SetTitle("Buttons in a VGroup.")

	btn1, _ = neuik.MakeButton("Update Title!")
	neuik.Element_Configure(btn1, "FillAll")
	neuik.Element_SetCallback(btn1, "OnClicked", cbButton_PushMe, mw)

	btn2, _ = neuik.MakeButton("Set BG Black")
	neuik.Element_Configure(btn2, "FillAll")
	neuik.Element_SetCallback(btn2, "OnClicked", cbSetBGColor, mw, "0,0,0,255")

	btn3, _ = neuik.MakeButton("Set BG Blue")
	neuik.Element_Configure(btn3, "FillAll")
	neuik.Element_SetCallback(btn3, "OnClicked", cbSetBGColor, mw, "0,0,200,255")

	btn4, _ = neuik.MakeButton("Set BG White")
	neuik.Element_Configure(btn4, "FillAll", "VScale=2.0")
	neuik.Element_SetCallback(btn4, "OnClicked", cbSetBGColor, mw, "255,255,255,255")

	btn5, _ = neuik.MakeButton("Set BG Pink")
	neuik.Element_Configure(btn5, "FillAll")
	neuik.Element_SetCallback(btn5, "OnClicked", cbSetBGColor, mw, "255,200,200,255")

	vg, _ = neuik.NewVGroup()
	neuik.Element_Configure(vg, "FillAll", "PadAll=10")
	neuik.Container_AddElements(vg, btn1, btn2, btn3, btn4, btn5)
	vg.SetVSpacing(3)

	mw.SetElement(vg)
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
