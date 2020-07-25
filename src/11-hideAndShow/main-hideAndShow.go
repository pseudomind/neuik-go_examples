/* example -- hide and show NEUIK elements */
package main

import (
	"github.com/pseudomind/neuik-go/neuik"
	"errors"
)

func cb_SetShown(args []interface{}) {
	var (
		elem    neuik.Element
		isShown bool
	)

	/* Typecast the first argument (elem *neuik.Element) */
	switch arg := args[0].(type) {
	case neuik.Element:
		elem = arg
	default:
		panic(errors.New("First argument `elem *neuik.Element` is not of the right type"))
	}

	/* Typecast the second argument (isShown bool) */
	switch arg := args[1].(type) {
	case bool:
		isShown = arg
	default:
		panic(errors.New("Second argument `isShown bool` is not of the right type"))
	}

	if isShown {
		neuik.Element_Configure(elem, "Show")
	} else {
		neuik.Element_Configure(elem, "!Show")
	}
}

func main() {
	var (
		mw            *neuik.Window
		vg0           *neuik.VGroup
		vgSHLabels    *neuik.VGroup
		vgSHButtons   *neuik.VGroup
		hgBtns        *neuik.HGroup
		shItems       *neuik.HGroup
		shBtn_Labels  *neuik.ToggleButton
		shBtn_Buttons *neuik.ToggleButton
		label_1       *neuik.Label
		label_2       *neuik.Label
		label_3       *neuik.Label
		btn_1         *neuik.Button
		btn_2         *neuik.Button
		btn_3         *neuik.Button
		shFrame       *neuik.Frame
	)

	if neuik.Init() {
		goto out
	}

	mw, _ = neuik.NewWindow()
	mw.SetSize(300, 150)
	mw.SetTitle("Show/Hide Elements -- example")

	/*------------------------------------------------------------------------*/
	/* Create and configure the text labels and buttons                       */
	/*------------------------------------------------------------------------*/
	label_1, _ = neuik.MakeLabel("These are labels")
	label_2, _ = neuik.MakeLabel("within a single")
	label_3, _ = neuik.MakeLabel("vertical group.")

	vgSHLabels, _ = neuik.NewVGroup()
	neuik.Element_Configure(vgSHLabels, "!Show")
	neuik.Container_AddElements(vgSHLabels, label_1, label_2, label_3)

	shBtn_Labels, _ = neuik.MakeToggleButton("Show Labels")

	/* set the on activated callback */
	neuik.Element_SetCallback(shBtn_Labels, "OnActivated",
		cb_SetShown, vgSHLabels, true)

	/* set the on deActivated callback */
	neuik.Element_SetCallback(shBtn_Labels, "OnDeactivated",
		cb_SetShown, vgSHLabels, false)

	btn_1, _ = neuik.MakeButton("Button 1")
	btn_2, _ = neuik.MakeButton("Button 2")
	btn_3, _ = neuik.MakeButton("Button 3")

	vgSHButtons, _ = neuik.NewVGroup()
	neuik.Element_Configure(vgSHButtons, "!Show")
	neuik.Container_AddElements(vgSHButtons, btn_1, btn_2, btn_3)

	shBtn_Buttons, _ = neuik.MakeToggleButton("Show Buttons")

	/* set the on activated callback */
	neuik.Element_SetCallback(shBtn_Buttons, "OnActivated",
		cb_SetShown, vgSHButtons, true)

	/* set the on deActivated callback */
	neuik.Element_SetCallback(shBtn_Buttons, "OnDeactivated",
		cb_SetShown, vgSHButtons, false)

	shItems, _ = neuik.NewHGroup()
	neuik.Container_AddElements(shItems, vgSHLabels, vgSHButtons)

	shFrame, _ = neuik.NewFrame()
	neuik.Container_SetElement(shFrame, shItems)

	/*------------------------------------------------------------------------*/
	/* Create and load the horizontal groups into the vertical group          */
	/*------------------------------------------------------------------------*/
	hgBtns, _ = neuik.NewHGroup()
	neuik.Container_AddElements(hgBtns, shBtn_Labels, shBtn_Buttons)

	vg0, _ = neuik.NewVGroup()
	neuik.Element_Configure(vg0, "VJustify=top")
	neuik.Container_AddElements(vg0, hgBtns, shFrame)
	vg0.SetVSpacing(10)

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
