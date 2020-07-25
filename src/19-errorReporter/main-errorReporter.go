package main

import (
	"github.com/pseudomind/neuik-go/neuik"
)

func main() {
	var (
		err error
		mw   *neuik.Window
		lg   *neuik.ListGroup
		// frm  *neuik.Frame
		lbl1 *neuik.Label
		lbl2 *neuik.Label
		lbl3 *neuik.Label
		lbl4 *neuik.Label
		lbl5 *neuik.Label
		lbl6 *neuik.Label
		// lbl7 *neuik.Label
		// lbl8 *neuik.Label
		// lbl9 *neuik.Label
		row1 *neuik.ListRow
		row2 *neuik.ListRow
		row3 *neuik.ListRow
		row4 *neuik.ListRow
		row5 *neuik.ListRow
	)

	if neuik.Init() {
		goto out
	}
	neuik.SetAppName("NEUIK Example: Error Reporter")

	mw, err = neuik.NewWindow()
	if err != nil {
		goto out
	}
	mw.SetTitle("This window has a button.")
	mw.Configure("Resizable")

	lbl1, _ = neuik.MakeLabel("Text Label #1")
	lbl2, _ = neuik.MakeLabel("Text Label #2")
	lbl3, _ = neuik.MakeLabel("Text Label #3")
	lbl4, _ = neuik.MakeLabel("Text Label #4")
	lbl5, _ = neuik.MakeLabel("Text Label #5")
	// lbl6, _ = neuik.MakeLabel("Text Label #6")
	// lbl7, _ = neuik.MakeLabel("Text Label #7")
	// lbl8, _ = neuik.MakeLabel("Text Label #8")
	// lbl9, _ = neuik.MakeLabel("Text Label #9")

	row1, _ = neuik.NewListRow()
	row2, _ = neuik.NewListRow()
	row3, _ = neuik.NewListRow()
	row4, _ = neuik.NewListRow()
	row5, _ = neuik.NewListRow()


	neuik.Container_AddElements(row1, lbl1)
	neuik.Container_AddElements(row2, lbl2)
	neuik.Container_AddElements(row3, lbl3)
	neuik.Container_AddElements(row4, lbl4)
	neuik.Container_AddElements(row5, lbl5)


	lg, _ = neuik.NewListGroup()
	neuik.Element_Configure(lg, "FillAll", "PadAll=1")
	// neuik.Container_AddElements(lg,
	// 	lbl1, lbl2, lbl3, lbl4, lbl5, lbl6, lbl7, lbl8, lbl9)
	neuik.Container_AddElements(lg,
		row1, row2, row3, row4, row5)

	// frm, _ = neuik.NewFrame()
	// neuik.Element_Configure(frm, "FillAll", "PadAll=1")
	// neuik.Container_SetElement(frm, lg)

	// mw.SetElement(frm)
	mw.SetElement(lbl6)
	if neuik.HasErrors() {
		goto out
	}
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
