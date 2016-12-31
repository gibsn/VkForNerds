package ui

import (
	"api"
	"fmt"

	"github.com/gizak/termui"
)

type Grid termui.Grid

type Ui struct {
	CurrGrid *termui.Grid
}

const (
	dialogHeaderHeight = 7
)

func NewUi() *Ui {
	ui := &Ui{}

	err := termui.Init()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// ui.initDialogsHeaders()
	// ui.initDialog()

	return ui
}

func (this *Ui) Routine() {
	defer termui.Close()

	termui.Loop()
}

func (this *Ui) ResizeCurrentGrid() {
	this.CurrGrid.Width = termui.TermWidth()
	this.CurrGrid.Align()
	termui.Clear()
	termui.Render(this.CurrGrid)
}

func (this *Ui) ChangeCurrentGrid(newGrid *termui.Grid) {
	this.CurrGrid = newGrid
	termui.Clear()
	termui.Render(this.CurrGrid)
}

func (this *Ui) initDialogsHeaders() {
}

func (this *Ui) CreateDialogGrid(dialog *api.Dialog) *termui.Grid {
	dialogGrid := termui.NewGrid()
	dialogGrid.Width = termui.TermWidth()

	messages := termui.NewPar("$$MESSAGES$$")
	messages.Height = int(float32(termui.TermHeight()) * 0.8)

	input := termui.NewPar("$$INPUT$$")
	input.Height = int(float32(termui.TermHeight()) * 0.2)

	dialogGrid.AddRows(
		termui.NewRow(
			termui.NewCol(12, 0, messages),
		),
		termui.NewRow(
			termui.NewCol(12, 0, input),
		),
	)

	dialogGrid.Align()

	return dialogGrid
}

func (this *Ui) CreateDialogsGrid(dialogs []api.Dialog) *termui.Grid {
	dialogsGrid := termui.NewGrid()
	dialogsGrid.Width = termui.TermWidth()

	nDialogs := termui.TermHeight()/dialogHeaderHeight + 1
	for i := 0; i < nDialogs && i < len(dialogs); i++ {
		var dialogTitle string
		if dialogs[i].Title != " ... " {
			dialogTitle = dialogs[i].Title
		} else {
			dialogTitle = fmt.Sprintf("%d", dialogs[i].Uid)
		}

		//TODO: mention files in message like [Photo] or smth
		str := fmt.Sprintf("%s\n\n%s", dialogTitle, dialogs[i].FirstMessage)
		messageHeader := termui.NewPar(str)
		messageHeader.Height = dialogHeaderHeight

		dialogsGrid.AddRows(
			termui.NewRow(
				termui.NewCol(12, 0, messageHeader),
			),
		)
	}

	dialogsGrid.Align()

	return dialogsGrid
}
