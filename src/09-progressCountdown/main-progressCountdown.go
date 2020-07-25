/* progress bar example */

package main

import (
	"github.com/pseudomind/neuik-go/neuik"
	"errors"
	"fmt"
	"time"
)

var TaskActive bool

func backgroundTask(taskTime int, progBar *neuik.ProgressBar, progMsg *neuik.Label) {
	var (
		frac     float64
		flTime   float64
		thisTime float64
		timeStep int = 50 // ms
	)

	//------------------------------------------------------------------------//
	// This is the only background task being run                             //
	//------------------------------------------------------------------------//
	TaskActive = true
	flTime = (float64)(taskTime)

	for {
		frac = thisTime / flTime
		if frac >= 1.0 {
			progBar.SetFraction(1.0)
			goto out
		}

		progBar.SetFraction(frac)
		time.Sleep(time.Millisecond * time.Duration(timeStep))
		thisTime += (float64)(timeStep) / 1000.0
	}

out:
	progMsg.SetText("")
	TaskActive = false
}

func cbStartCountdown(args []interface{}) {
	var (
		taskMsg  string
		taskTime int
		progBar  *neuik.ProgressBar
		progMsg  *neuik.Label
	)

	/* Typecast the first argument (taskTime int) */
	switch arg := args[0].(type) {
	case int:
		taskTime = arg
	default:
		panic(errors.New("First argument `taskTime int` is not of the right type"))
	}

	/* Typecast the second argument (progBar *neuik.ProgressBar) */
	switch arg := args[1].(type) {
	case *neuik.ProgressBar:
		progBar = arg
	default:
		panic(errors.New("Second argument `progBar *neuik.ProgressBar` is not of the right type"))
	}

	/* Typecast the third argument (progMsg *neuik.Label) */
	switch arg := args[2].(type) {
	case *neuik.Label:
		progMsg = arg
	default:
		panic(errors.New("Third argument `progBar *neuik.Label` is not of the right type"))
	}

	if !TaskActive {
		taskMsg = fmt.Sprintf("%ds countdown in progress...", taskTime)
		progMsg.SetText(taskMsg)

		go backgroundTask(taskTime, progBar, progMsg)
	}
}

func main() {
	var (
		mw         *neuik.Window
		vg0        *neuik.VGroup /* contains instrcution text */
		hg0        *neuik.HGroup /* contains the countdown buttons */
		label_inst *neuik.Label
		label_cnt  *neuik.Label
		label_msg  *neuik.Label
		btn_10s    *neuik.Button
		btn_30s    *neuik.Button
		btn_60s    *neuik.Button
		progBar    *neuik.ProgressBar
		hlineRes   *neuik.Line
		oFrame     *neuik.Frame
	)

	if neuik.Init() {
		goto out
	}

	mw, _ = neuik.NewWindow()
	mw.SetSize(400, 200)
	mw.SetTitle("ProgressBar -- Countdown example")

	/*------------------------------------------------------------------------*/
	/* Create and configure the text label                                    */
	/*------------------------------------------------------------------------*/
	label_inst, _ = neuik.MakeLabel(
		"Click on one of the following buttons to start a countdown:")
	neuik.Element_Configure(label_inst, "FillAll")
	label_cnt, _ = neuik.MakeLabel("Countdown for:")
	neuik.Element_Configure(label_cnt, "!FillAll")
	label_msg, _ = neuik.NewLabel()
	neuik.Element_Configure(label_msg, "HFill", "!VFill")

	/*------------------------------------------------------------------------*/
	/* Create and configure the progress bar                                  */
	/*------------------------------------------------------------------------*/
	progBar, _ = neuik.NewProgressBar()
	neuik.Element_Configure(progBar, "HFill", "!VFill")

	/*------------------------------------------------------------------------*/
	/* Create and configure the countdown buttons                             */
	/*------------------------------------------------------------------------*/
	btn_10s, _ = neuik.MakeButton("10s")
	neuik.Element_Configure(btn_10s, "!HFill", "VFill")
	neuik.Element_SetCallback(btn_10s, "OnClicked", cbStartCountdown,
		10, progBar, label_msg)

	btn_30s, _ = neuik.MakeButton("30s")
	neuik.Element_Configure(btn_30s, "!HFill", "VFill")
	neuik.Element_SetCallback(btn_30s, "OnClicked", cbStartCountdown,
		30, progBar, label_msg)

	btn_60s, _ = neuik.MakeButton("60s")
	neuik.Element_Configure(btn_60s, "!HFill", "VFill")
	neuik.Element_SetCallback(btn_60s, "OnClicked", cbStartCountdown,
		60, progBar, label_msg)

	/*------------------------------------------------------------------------*/
	/* Create the HLine                                                       */
	/*------------------------------------------------------------------------*/
	hlineRes, _ = neuik.NewHLine()
	hlineRes.SetThickness(1)

	/*------------------------------------------------------------------------*/
	/* Create and load the items into horizontal groups                       */
	/*------------------------------------------------------------------------*/
	hg0, _ = neuik.NewHGroup()
	neuik.Container_AddElements(hg0, label_cnt, btn_10s, btn_30s, btn_60s)
	hg0.SetHSpacing(3)

	/*------------------------------------------------------------------------*/
	/* Create and load the horizontal groups into the vertical group          */
	/*------------------------------------------------------------------------*/
	vg0, _ = neuik.NewVGroup()
	neuik.Element_Configure(vg0, "FillAll", "PadAll=5")
	neuik.Container_AddElements(vg0, label_inst, hg0, progBar, hlineRes, label_msg)
	vg0.SetVSpacing(10)

	oFrame, _ = neuik.NewFrame()
	neuik.Element_Configure(oFrame, "HFill", "PadAll=10")
	neuik.Container_SetElement(oFrame, vg0)

	mw.SetElement(oFrame)
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
