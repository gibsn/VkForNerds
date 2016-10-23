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
}

func NewUi() *Ui {
	return &Ui{}
}

func (this *Ui) Init() {
	err := termui.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer termui.Close()

	this.InitDialogsHeaders()
	this.setCommandModeHandlers()

	termui.Loop()
}

func (this *Ui) InitDialogsHeaders() {
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

func (*Ui) InitDialog() {
	dialog := termui.NewPar("$$MESSAGES$$")
	dialog.Height = int(float32(termui.TermHeight()) * 0.8)

	input := termui.NewPar("$$INPUT$$")
	input.Height = int(float32(termui.TermHeight()) * 0.2)

	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, dialog),
		),
		termui.NewRow(
			termui.NewCol(12, 0, input),
		),
	)
	termui.Body.Align()

	termui.Render(termui.Body)
}
