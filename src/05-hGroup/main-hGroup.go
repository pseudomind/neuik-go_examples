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
	var mw *neuik.Window // arg1

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
		hg   *neuik.HGroup
		btn1 *neuik.Button
		btn2 *neuik.Button
		btn3 *neuik.Button
	)

	if neuik.Init() {
		neuik.BacktraceErrors()
		return
	}

	mw, _ = neuik.NewWindow()
	mw.SetSize(640, 480)
	mw.SetTitle("Buttons in a HGroup.")

	btn1, _ = neuik.MakeButton("Update Title!")
	neuik.Element_Configure(btn1, "FillAll", "PadTop=20")
	neuik.Element_SetCallback(btn1, "OnClicked", cbButton_PushMe, mw)

	btn2, _ = neuik.MakeButton("Set BG Black")
	neuik.Element_Configure(btn2, "FillAll")
	neuik.Element_SetCallback(btn2, "OnClicked", cbSetBGColor, mw, "0,0,0,255")

	btn3, _ = neuik.MakeButton("Set BG Blue")
	neuik.Element_Configure(btn3, "FillAll")
	neuik.Element_SetCallback(btn3, "OnClicked", cbSetBGColor, mw, "0,0,200,255")

	hg, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg, "FillAll", "PadAll=10")
	neuik.Container_AddElements(hg, btn1, btn2, btn3)
	hg.SetHSpacing(3)

	mw.SetElement(hg)
	mw.Create()
	neuik.EventLoop(true)

	if neuik.HasErrors() {
		neuik.BacktraceErrors()
	}

	neuik.Quit()
}
