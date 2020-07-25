package main

import (
	"github.com/pseudomind/neuik-go/neuik"
)

func main() {
	var (
		mw           *neuik.Window
		vg0          *neuik.VGroup
		hg0          *neuik.HGroup
		hg1          *neuik.HGroup
		hg2          *neuik.HGroup
		label_left   *neuik.Label
		label_center *neuik.Label
		label_right  *neuik.Label
		te_left      *neuik.TextEntry
		te_center    *neuik.TextEntry
		te_right     *neuik.TextEntry
	)

	if neuik.Init() {
		neuik.BacktraceErrors()
		return
	}

	mw, _ = neuik.NewWindow()
	mw.SetTitle("Justification/TextEntry Example")

	/*------------------------------------------------------------------------*/
	/* Create and configure the text labels                                   */
	/*------------------------------------------------------------------------*/
	label_left, _ = neuik.MakeLabel("JUSTIFY_LEFT:")
	neuik.Element_Configure(label_left, "FillAll", "HJustify=left")

	label_center, _ = neuik.MakeLabel("JUSTIFY_CENTER:")
	neuik.Element_Configure(label_center, "FillAll", "HJustify=center")

	label_right, _ = neuik.MakeLabel("JUSTIFY_RIGHT:")
	neuik.Element_Configure(label_right, "FillAll", "HJustify=right")

	/*------------------------------------------------------------------------*/
	/* Create and configure the text entry boxes                              */
	/*------------------------------------------------------------------------*/
	te_left, _ = neuik.NewTextEntry()
	neuik.Element_Configure(te_left, "FillAll", "HJustify=left")

	te_center, _ = neuik.NewTextEntry()
	neuik.Element_Configure(te_center, "FillAll")
	te_center.Configure("HJustify=center")

	te_right, _ = neuik.NewTextEntry()
	neuik.Element_Configure(te_right, "FillAll")
	te_right.Configure("HJustify=right")

	/*------------------------------------------------------------------------*/
	/* Create and load the items into horizontal groups                       */
	/*------------------------------------------------------------------------*/
	hg0, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg0, "FillAll")
	neuik.Container_AddElements(hg0, label_left, te_left)
	hg0.SetHSpacing(3)

	hg1, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg1, "FillAll")
	neuik.Container_AddElements(hg1, label_center, te_center)
	hg1.SetHSpacing(3)

	hg2, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg2, "FillAll")
	neuik.Container_AddElements(hg2, label_right, te_right)
	hg2.SetHSpacing(3)

	/*------------------------------------------------------------------------*/
	/* Create and load the horizontal groups into the vertical group          */
	/*------------------------------------------------------------------------*/
	vg0, _ = neuik.NewVGroup()
	neuik.Element_Configure(vg0, "HFill", "PadAll=5")
	neuik.Container_AddElements(vg0, hg0, hg1, hg2)
	vg0.SetVSpacing(3)

	mw.SetElement(vg0)

	/*------------------------------------------------------------------------*/
	/* Check that no errors were experienced before starting the event loop   */
	/*------------------------------------------------------------------------*/
	if neuik.HasErrors() {
		neuik.BacktraceErrors()
		neuik.Quit()
	}

	mw.Create()
	neuik.EventLoop(true)

	if neuik.HasErrors() {
		neuik.BacktraceErrors()
	}

	neuik.Quit()
}
