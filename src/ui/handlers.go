package ui

import (
	"fmt"

	"github.com/gizak/termui"
)

func (this *Ui) resizeCurrentGrid() {
	this.CurrGrid.Width = termui.TermWidth()
	this.CurrGrid.Align()
	termui.Clear()
	termui.Render(this.CurrGrid)
}

func (this *Ui) changeCurrentGrid(newGrid *termui.Grid) {
	this.CurrGrid = newGrid
	termui.Clear()
	termui.Render(this.CurrGrid)
}

func (this *Ui) setCommandModeHandlers() {
	termui.DefaultEvtStream.ResetHandlers()

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/m", func(e termui.Event) {
		this.changeCurrentGrid(this.DialogsHeaders)
	})

	termui.Handle("/sys/kbd/i", func(e termui.Event) {
		this.setInsertModeHandlers()
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		this.changeCurrentGrid(this.Dialog)
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		this.resizeCurrentGrid()
	})
}

func (this *Ui) setInsertModeHandlers() {
	termui.DefaultEvtStream.ResetHandlers()

	//here i can print all of the chars
	termui.Handle("/sys/kbd/", func(e termui.Event) {
		fmt.Println(e.Data)
	})

	termui.Handle("/sys/kbd/<escape>", func(e termui.Event) {
		this.setCommandModeHandlers()
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		//send message here
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		this.resizeCurrentGrid()
	})
}
