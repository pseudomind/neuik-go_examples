package main

import (
	"github.com/pseudomind/neuik-go/neuik"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ClearOnNextOp bool
	LastValue     float64
	LastOp        byte = 0
	Buffer        []byte
)

func cbBtnPushed(args []interface{}) {
	var (
		disp    *neuik.Button // arg1
		btnOp   string        // arg2
		dispStr string
		thisVal float64
		result  float64
	)

	if len(args) < 2 {
		panic(errors.New("This function requires two arguments."))
	}

	/* Typecast the first argument (disp *neuik.Button) */
	switch arg := args[0].(type) {
	case *neuik.Button:
		disp = arg
	default:
		panic(errors.New("1st argument `disp *neuik.Button` is not of the right type"))
	}

	/* Typecast the second argument (btnOp string) */
	switch arg := args[1].(type) {
	case string:
		btnOp = arg
	default:
		panic(errors.New("2nd argument `btnOp string` is not of the right type"))
	}

	if ClearOnNextOp {
		//--------------------------------------------------------------------//
		// After a `+,-,*,/` is used, the next operation will clear the       //
		// display before continuing.                                         //
		//--------------------------------------------------------------------//
		ClearOnNextOp = false
		disp.SetText("")
	}
	dispStr, _ = disp.GetText()

	switch btnOp {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		//--------------------------------------------------------------------//
		// A number button, add it to the end of the display.                 //
		//--------------------------------------------------------------------//
		dispStr = dispStr + btnOp
		disp.SetText(dispStr)
	case ".":
		if !strings.Contains(dispStr, ".") {
			dispStr = dispStr + btnOp
			disp.SetText(dispStr)
		}
	case "CLEAR":
		disp.SetText("")
		LastOp = 0
		LastValue = 0.0
	case "+":
		LastOp = '+'
		LastValue, _ = strconv.ParseFloat(dispStr, 64)
		ClearOnNextOp = true
	case "-":
		LastOp = '-'
		LastValue, _ = strconv.ParseFloat(dispStr, 64)
		ClearOnNextOp = true
	case "*":
		LastOp = '*'
		LastValue, _ = strconv.ParseFloat(dispStr, 64)
		ClearOnNextOp = true
	case "/":
		LastOp = '/'
		LastValue, _ = strconv.ParseFloat(dispStr, 64)
		ClearOnNextOp = true
	case "=":
		if LastOp == 0 {
			return
		}
		thisVal, _ = strconv.ParseFloat(dispStr, 64)

		switch LastOp {
		case 0:
		case '+':
			result = LastValue + thisVal
		case '-':
			result = LastValue - thisVal
		case '*':
			result = LastValue * thisVal
		case '/':
			if thisVal == 0.0 {
				disp.SetText("ERROR: Attempted Div By Zero")
				return
			}
			result = LastValue / thisVal
		}
		dispStr = fmt.Sprintf("%g", result)
		disp.SetText(dispStr)
	}
}

func main() {
	var (
		mw         *neuik.Window
		hg         *neuik.HGroup
		vg0        *neuik.VGroup
		vg1        *neuik.VGroup
		vg2        *neuik.VGroup
		vg3        *neuik.VGroup
		vg4        *neuik.VGroup
		btn0       *neuik.Button
		btn1       *neuik.Button
		btn2       *neuik.Button
		btn3       *neuik.Button
		btn4       *neuik.Button
		btn5       *neuik.Button
		btn6       *neuik.Button
		btn7       *neuik.Button
		btn8       *neuik.Button
		btn9       *neuik.Button
		btnDot     *neuik.Button
		btnClr     *neuik.Button
		btnPlus    *neuik.Button
		btnMinus   *neuik.Button
		btnMult    *neuik.Button
		btnDiv     *neuik.Button
		btnEqual   *neuik.Button
		btnDisplay *neuik.Button
		dFontSize  string = "FontSize=18"
		dFontBlue  string = "FontColor=0,0,200,0"
	)

	if neuik.Init() {
		neuik.BacktraceErrors()
		return
	}

	/*------------------------------------------------------------------------*/
	/* Numeric Buttons                                                        */
	/*------------------------------------------------------------------------*/
	mw, _ = neuik.NewWindow()
	mw.SetTitle("Basic Calculator")
	mw.Configure("BGColor=180,180,180,0")

	btnDisplay, _ = neuik.NewButton()
	btnDisplay.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btnDisplay, "HFill")

	btn0, _ = neuik.MakeButton("0")
	btn0.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn0, "FillAll")
	neuik.Element_SetCallback(btn0, "OnClicked", cbBtnPushed, btnDisplay, "0")

	btn1, _ = neuik.MakeButton("1")
	btn1.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn1, "FillAll")
	neuik.Element_SetCallback(btn1, "OnClicked", cbBtnPushed, btnDisplay, "1")

	btn2, _ = neuik.MakeButton("2")
	btn2.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn2, "FillAll")
	neuik.Element_SetCallback(btn2, "OnClicked", cbBtnPushed, btnDisplay, "2")

	btn3, _ = neuik.MakeButton("3")
	btn3.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn3, "FillAll")
	neuik.Element_SetCallback(btn3, "OnClicked", cbBtnPushed, btnDisplay, "3")

	btn4, _ = neuik.MakeButton("4")
	btn4.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn4, "FillAll")
	neuik.Element_SetCallback(btn4, "OnClicked", cbBtnPushed, btnDisplay, "4")

	btn5, _ = neuik.MakeButton("5")
	btn5.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn5, "FillAll")
	neuik.Element_SetCallback(btn5, "OnClicked", cbBtnPushed, btnDisplay, "5")

	btn6, _ = neuik.MakeButton("6")
	btn6.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn6, "FillAll")
	neuik.Element_SetCallback(btn6, "OnClicked", cbBtnPushed, btnDisplay, "6")

	btn7, _ = neuik.MakeButton("7")
	btn7.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn7, "FillAll")
	neuik.Element_SetCallback(btn7, "OnClicked", cbBtnPushed, btnDisplay, "7")

	btn8, _ = neuik.MakeButton("8")
	btn8.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn8, "FillAll")
	neuik.Element_SetCallback(btn8, "OnClicked", cbBtnPushed, btnDisplay, "8")

	btn9, _ = neuik.MakeButton("9")
	btn9.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btn9, "FillAll")
	neuik.Element_SetCallback(btn9, "OnClicked", cbBtnPushed, btnDisplay, "9")

	/*------------------------------------------------------------------------*/
	/* Special Buttons                                                        */
	/*------------------------------------------------------------------------*/
	btnDot, _ = neuik.MakeButton(".")
	btnDot.Configure(dFontSize, "FontBold")
	neuik.Element_Configure(btnDot, "FillAll")
	neuik.Element_SetCallback(btnDot, "OnClicked", cbBtnPushed, btnDisplay, ".")

	btnClr, _ = neuik.MakeButton("C")
	btnClr.Configure(dFontSize, dFontBlue)
	neuik.Element_Configure(btnClr, "FillAll")
	neuik.Element_SetCallback(btnClr, "OnClicked", cbBtnPushed, btnDisplay, "CLEAR")

	btnPlus, _ = neuik.MakeButton("+")
	btnPlus.Configure(dFontSize, dFontBlue)
	neuik.Element_Configure(btnPlus, "FillAll")
	neuik.Element_SetCallback(btnPlus, "OnClicked", cbBtnPushed, btnDisplay, "+")

	btnMinus, _ = neuik.MakeButton("-")
	btnMinus.Configure(dFontSize, dFontBlue)
	neuik.Element_Configure(btnMinus, "FillAll")
	neuik.Element_SetCallback(btnMinus, "OnClicked", cbBtnPushed, btnDisplay, "-")

	btnMult, _ = neuik.MakeButton("*")
	btnMult.Configure(dFontSize, dFontBlue)
	neuik.Element_Configure(btnMult, "FillAll")
	neuik.Element_SetCallback(btnMult, "OnClicked", cbBtnPushed, btnDisplay, "*")

	btnDiv, _ = neuik.MakeButton("/")
	btnDiv.Configure(dFontSize, dFontBlue)
	neuik.Element_Configure(btnDiv, "FillAll")
	neuik.Element_SetCallback(btnDiv, "OnClicked", cbBtnPushed, btnDisplay, "/")

	btnEqual, _ = neuik.MakeButton("=")
	btnEqual.Configure(dFontSize, dFontBlue)
	neuik.Element_Configure(btnEqual, "FillAll")
	neuik.Element_SetCallback(btnEqual, "OnClicked", cbBtnPushed, btnDisplay, "=")

	vg1, _ = neuik.NewVGroup()
	vg1.SetVSpacing(3)
	neuik.Element_Configure(vg1, "FillAll", "HScale=3.0")
	neuik.Container_AddElements(vg1, btn7, btn4, btn1, btn0)

	vg2, _ = neuik.NewVGroup()
	vg2.SetVSpacing(3)
	neuik.Element_Configure(vg2, "FillAll", "HScale=3.0")
	neuik.Container_AddElements(vg2, btn8, btn5, btn2, btnDot)

	vg3, _ = neuik.NewVGroup()
	vg3.SetVSpacing(3)
	neuik.Element_Configure(vg3, "FillAll", "HScale=3.0")
	neuik.Container_AddElements(vg3, btn9, btn6, btn3, btnClr)

	vg4, _ = neuik.NewVGroup()
	vg4.SetVSpacing(3)
	neuik.Element_Configure(vg4, "FillAll", "HScale=2.0")
	neuik.Container_AddElements(vg4, btnDiv, btnMult, btnMinus, btnPlus, btnEqual)

	hg, _ = neuik.NewHGroup()
	hg.SetHSpacing(3)
	neuik.Element_Configure(hg, "FillAll")
	neuik.Container_AddElements(hg, vg1, vg2, vg3, vg4)

	vg0, _ = neuik.NewVGroup()
	vg0.SetVSpacing(10)
	neuik.Element_Configure(vg0, "FillAll", "PadAll=10")
	neuik.Container_AddElements(vg0, btnDisplay, hg)

	mw.SetElement(vg0)
	mw.Create()

	if !neuik.HasErrors() {
		neuik.EventLoop(true)
	}

	if neuik.HasErrors() {
		neuik.BacktraceErrors()
	}

	neuik.Quit()
}
