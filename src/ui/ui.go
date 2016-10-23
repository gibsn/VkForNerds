package ui

import (
	"fmt"
	"os"

	"github.com/gizak/termui"
)

const (
	dialogHeaderHeight = 7
)

type Ui struct {
	DialogsHeaders *termui.Grid
	Dialogs        []*termui.Grid
	Dialog         *termui.Grid
}

func NewUi() *Ui {
	ui := &Ui{}

	err := termui.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ui.initDialogsHeaders()
	ui.initDialog()
	ui.setCommandModeHandlers()

	return ui
}

func (*Ui) CloseUi() {
	termui.Close()
}

func (this *Ui) Start() {
	termui.Loop()
}

func (this *Ui) initDialogsHeaders() {
	this.DialogsHeaders = termui.NewGrid()
	this.DialogsHeaders.Width = termui.TermWidth()

	nDialogs := termui.TermHeight()/dialogHeaderHeight + 1
	for i := 0; i < nDialogs; i++ {
		dialogHeader := termui.NewPar("$$USERNAME$$\n\n$$TEXT$$")
		dialogHeader.Height = dialogHeaderHeight

		this.DialogsHeaders.AddRows(
			termui.NewRow(
				termui.NewCol(12, 0, dialogHeader),
			),
		)
	}
	this.DialogsHeaders.Align()
}

func (this *Ui) initDialog() {
	this.Dialog = termui.NewGrid()
	this.Dialog.Width = termui.TermWidth()

	dialog := termui.NewPar("$$MESSAGES$$")
	dialog.Height = int(float32(termui.TermHeight()) * 0.8)

	input := termui.NewPar("$$INPUT$$")
	input.Height = int(float32(termui.TermHeight()) * 0.2)

	this.Dialog.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, dialog),
		),
		termui.NewRow(
			termui.NewCol(12, 0, input),
		),
	)
	this.Dialog.Align()
}
