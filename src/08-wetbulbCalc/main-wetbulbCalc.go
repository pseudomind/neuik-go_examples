package main

import (
	"github.com/pseudomind/neuik-go/neuik"
	"errors"
)

// void cbFunc_TE_Act(
// 	void * window,
// 	void * teInp,
// 	void * teOutp)
// {
// 	NEUIK_TextEntry *  outp    = NULL;
// 	NEUIK_TextEntry ** teArry  = NULL;

// 	outp   = (NEUIK_TextEntry *)teOutp;
// 	teArry = (NEUIK_TextEntry **)teInp;

// 	if (NEUIK_TextEntry_GetText(teArry[0]) != NULL &&
// 		NEUIK_TextEntry_GetText(teArry[1]) != NULL &&
// 		NEUIK_TextEntry_GetText(teArry[2]) != NULL)
// 	{
// 		if ((*NEUIK_TextEntry_GetText(teArry[0])) != '\0' &&
// 			(*NEUIK_TextEntry_GetText(teArry[1])) != '\0' &&
// 			(*NEUIK_TextEntry_GetText(teArry[2])) != '\0'
// 		)
// 			NEUIK_TextEntry_SetText(outp, "Do a Calculation...");
// 	}
// }

func cbFunc_TE_Act(args []interface{}) {
	var (
		te *neuik.TextEntry
	)

	if len(args) < 1 {
		panic(errors.New("This function requires one argument."))
	}

	/* Typecast the first argument (te *neuik.TextEntry) */
	switch arg := args[0].(type) {
	case *neuik.TextEntry:
		te = arg
	default:
		panic(errors.New("First argument `te *neuik.TextEntry` is not of the right type"))
	}

	te.SetText("Do a Calculation...")
}

func main() {
	var (
		mw         *neuik.Window
		vg0        *neuik.VGroup /* contains the display and horiz. group */
		hg0        *neuik.HGroup /* T_DB */
		hg1        *neuik.HGroup /* P    */
		hg2        *neuik.HGroup /* rH   */
		hg3        *neuik.HGroup /* T_WB */
		label_T_DB *neuik.Label
		label_P    *neuik.Label
		label_rH   *neuik.Label
		label_T_WB *neuik.Label
		te_T_DB    *neuik.TextEntry
		te_P       *neuik.TextEntry
		te_rH      *neuik.TextEntry
		te_T_WB    *neuik.TextEntry
		hlineRes   *neuik.Line
		oFrame     *neuik.Frame
		teInp      [3]neuik.Element
	)

	if neuik.Init() {
		goto out
	}

	mw, _ = neuik.NewWindow()
	mw.SetTitle("Wetbulb Temperature Calculator")

	/*------------------------------------------------------------------------*/
	/* Create and configure the text labels                                   */
	/*------------------------------------------------------------------------*/
	label_T_DB, _ = neuik.MakeLabel("T_DB (F) :")
	neuik.Element_Configure(label_T_DB, "FillAll", "HJustify=right")

	label_P, _ = neuik.MakeLabel("P (psi) :")
	neuik.Element_Configure(label_P, "FillAll", "HJustify=right")

	label_rH, _ = neuik.MakeLabel("rH (%) :")
	neuik.Element_Configure(label_rH, "FillAll", "HJustify=right")

	label_T_WB, _ = neuik.MakeLabel("T_WB (F) :")
	neuik.Element_Configure(label_T_WB, "FillAll", "HJustify=right")

	/*------------------------------------------------------------------------*/
	/* Create and configure the text entry boxes                              */
	/*------------------------------------------------------------------------*/
	te_T_DB, _ = neuik.NewTextEntry()
	neuik.Element_Configure(te_T_DB, "FillAll")

	te_P, _ = neuik.NewTextEntry()
	neuik.Element_Configure(te_P, "FillAll")

	te_rH, _ = neuik.NewTextEntry()
	neuik.Element_Configure(te_rH, "FillAll")

	te_T_WB, _ = neuik.NewTextEntry()
	neuik.Element_Configure(te_T_WB, "FillAll")

	teInp[0] = te_T_DB
	teInp[1] = te_P
	teInp[2] = te_rH

	// neuik.Element_SetCallback(te_T_DB, "OnActivated", cbFunc_TE_Act, teInp, te_T_WB)
	// neuik.Element_SetCallback(te_P, "OnActivated", cbFunc_TE_Act, teInp, te_T_WB)
	// neuik.Element_SetCallback(te_rH, "OnActivated", cbFunc_TE_Act, teInp, te_T_WB)

	neuik.Element_SetCallback(te_T_DB, "OnActivated", cbFunc_TE_Act, te_T_WB)
	neuik.Element_SetCallback(te_P, "OnActivated", cbFunc_TE_Act, te_T_WB)
	neuik.Element_SetCallback(te_rH, "OnActivated", cbFunc_TE_Act, te_T_WB)

	/*------------------------------------------------------------------------*/
	/* Create the HLine                                                       */
	/*------------------------------------------------------------------------*/
	hlineRes, _ = neuik.NewHLine()
	// hlineRes->thickness = 1;
	hlineRes.SetThickness(1)

	/*------------------------------------------------------------------------*/
	/* Create and load the items into horizontal groups                       */
	/*------------------------------------------------------------------------*/
	hg0, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg0, "FillAll", "PadLeft=5", "PadRight=5")
	neuik.Container_AddElements(hg0, label_T_DB, te_T_DB)
	hg0.SetHSpacing(3)

	hg1, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg1, "FillAll", "PadLeft=5", "PadRight=5")
	neuik.Container_AddElements(hg1, label_P, te_P)
	hg1.SetHSpacing(3)

	hg2, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg2, "FillAll", "PadLeft=5", "PadRight=5")
	neuik.Container_AddElements(hg2, label_rH, te_rH)
	hg2.SetHSpacing(3)

	hg3, _ = neuik.NewHGroup()
	neuik.Element_Configure(hg3, "FillAll", "PadLeft=5", "PadRight=5")
	neuik.Container_AddElements(hg3, label_T_WB, te_T_WB)
	hg3.SetHSpacing(3)

	/*------------------------------------------------------------------------*/
	/* Create and load the horizontal groups into the vertical group          */
	/*------------------------------------------------------------------------*/
	vg0, _ = neuik.NewVGroup()
	neuik.Element_Configure(vg0, "FillAll", "PadTop=5", "PadBottom=5")
	neuik.Container_AddElements(vg0, hg0, hg1, hg2, hlineRes, hg3)
	vg0.SetVSpacing(5)

	oFrame, _ = neuik.NewFrame()
	neuik.Element_Configure(oFrame, "HFill", "PadAll=10")
	neuik.Container_SetElement(oFrame, vg0)

	mw.SetElement(oFrame)

	mw.Create()
	neuik.EventLoop(true)
out:
	if neuik.HasErrors() {
		neuik.BacktraceErrors()
	}

	neuik.Quit()
}
