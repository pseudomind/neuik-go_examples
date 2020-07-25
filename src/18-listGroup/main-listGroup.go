package main

import (
	"github.com/pseudomind/neuik-go/neuik"
	"errors"
	"fmt"
)

func cb_RowActivated(args []interface{}) {
	var (
		err     error
		rowText string
		label   *neuik.Label
	)

	/* Typecast the first argument (elem *neuik.Element) */
	switch arg := args[0].(type) {
	case *neuik.Label:
		label = arg
	default:
		panic(errors.New("First argument `label *neuik.Label` is not of the right type"))
	}

	rowText, err = label.GetText()
	if err != nil {
		panic(err)
	}
	fmt.Println(rowText)
}

func main() {
	var (
		err  error
		mw   *neuik.Window
		lg   *neuik.ListGroup
		lbl1 *neuik.Label
		lbl2 *neuik.Label
		lbl3 *neuik.Label
		lbl4 *neuik.Label
		lbl5 *neuik.Label
		row1 *neuik.ListRow
		row2 *neuik.ListRow
		row3 *neuik.ListRow
		row4 *neuik.ListRow
		row5 *neuik.ListRow
	)

	if neuik.Init() {
		goto out
	}
	neuik.SetAppName("NEUIK Example: ListGroup")

	mw, err = neuik.NewWindow()
	if err != nil {
		goto out
	}
	mw.SetTitle("NEUIK Example: ListGroup")
	mw.Configure("Resizable")

	lbl1, _ = neuik.MakeLabel("Text Label #1")
	lbl2, _ = neuik.MakeLabel("Text Label #2")
	lbl3, _ = neuik.MakeLabel("Text Label #3")
	lbl4, _ = neuik.MakeLabel("Text Label #4")
	lbl5, _ = neuik.MakeLabel("Text Label #5")

	row1, _ = neuik.NewListRow()
	row2, _ = neuik.NewListRow()
	row3, _ = neuik.NewListRow()
	row4, _ = neuik.NewListRow()
	row5, _ = neuik.NewListRow()

	neuik.Container_AddElements(row1, lbl1)
	neuik.Element_SetCallback(row1, "OnActivated", cb_RowActivated, lbl1)

	neuik.Container_AddElements(row2, lbl2)
	neuik.Element_SetCallback(row2, "OnActivated", cb_RowActivated, lbl2)

	neuik.Container_AddElements(row3, lbl3)
	neuik.Element_SetCallback(row3, "OnActivated", cb_RowActivated, lbl3)

	neuik.Container_AddElements(row4, lbl4)
	neuik.Element_SetCallback(row4, "OnActivated", cb_RowActivated, lbl4)

	neuik.Container_AddElements(row5, lbl5)
	neuik.Element_SetCallback(row5, "OnActivated", cb_RowActivated, lbl5)

	lg, _ = neuik.NewListGroup()
	neuik.Element_Configure(lg, "FillAll", "PadAll=1")
	lg.AddRows(row1, row2, row3, row4, row5)

	mw.SetElement(lg)
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
