/* example that counts time since the program started */
package main

import (
	"github.com/pseudomind/neuik-go/neuik"
	"errors"
	"fmt"
	"time"
)

func backgroundTask(timeLbl *neuik.Label) {
	var (
		thisTime  uint64
		oneSecond time.Duration = time.Duration(1000) * time.Millisecond
	)

	for {
		time.Sleep(oneSecond)
		thisTime += 1
		timeLbl.SetText(fmt.Sprintf("%d", thisTime))
	}
}

func cbOnWindowCreate(args []interface{}) {
	var timeLbl *neuik.Label

	/* Typecast the first/only argument (timeLbl *neuik.Label) */
	switch arg := args[0].(type) {
	case *neuik.Label:
		timeLbl = arg
	default:
		panic(errors.New("First/only argument `timeLbl *neuik.Label` is not of the right type"))
	}

	go backgroundTask(timeLbl)
}

func main() {
	var (
		mw         *neuik.Window
		vg0        *neuik.VGroup
		label_top  *neuik.Label
		label_elap *neuik.Label
		label_bot  *neuik.Label
	)

	if neuik.Init() {
		goto out
	}

	mw, _ = neuik.NewWindow()
	mw.SetSize(300, 150)
	mw.SetTitle("Time since started -- example")

	/*------------------------------------------------------------------------*/
	/* Create and configure the text label                                    */
	/*------------------------------------------------------------------------*/
	label_top, _ = neuik.MakeLabel("This program started")
	label_elap, _ = neuik.MakeLabel("0")
	label_bot, _ = neuik.MakeLabel("seconds ago.")

	/*------------------------------------------------------------------------*/
	/* Create and load the horizontal groups into the vertical group          */
	/*------------------------------------------------------------------------*/
	vg0, _ = neuik.NewVGroup()
	vg0.SetVSpacing(10)
	neuik.Container_AddElements(vg0, label_top, label_elap, label_bot)

	mw.SetCallback("OnCreated", cbOnWindowCreate, label_elap)
	mw.SetElement(vg0)

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
