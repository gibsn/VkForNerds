package ui

import (
	"fmt"

	"github.com/gizak/termui"
)

func (this *Ui) setCommandModeHandlers() {
	termui.DefaultEvtStream.ResetHandlers()

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/m", func(e termui.Event) {
		termui.Clear()
		termui.Render(this.DialogsHeaders)
	})

	termui.Handle("/sys/kbd/i", func(e termui.Event) {
		this.setInsertModeHandlers()
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		termui.Clear()
		termui.Render(this.Dialog)
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
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
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Clear()
		termui.Render(termui.Body)
	})
}
