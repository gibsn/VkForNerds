package ui

import (
	"fmt"

	"github.com/gizak/termui"
)

func (*Ui) clearAllKbdHandlers() {
	//waiting for https://github.com/gizak/termui/pull/98
	termui.Handle("/sys/kbd/", func(_ termui.Event) {})
	termui.Handle("/sys/kbd/q", func(_ termui.Event) {})
	termui.Handle("/sys/kbd/m", func(_ termui.Event) {})
	termui.Handle("/sys/kbd/i", func(_ termui.Event) {})
	termui.Handle("/sys/kbd/<enter>", func(_ termui.Event) {})
	termui.Handle("/sys/kbd/<esc>", func(_ termui.Event) {})
}

func (this *Ui) setCommandModeHandlers() {
	this.clearAllKbdHandlers()

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/m", func(e termui.Event) {
		termui.Render(this.DialogsHeaders)
	})

	termui.Handle("/sys/kbd/i", func(e termui.Event) {
		this.setInsertModeHandlers()
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		// this.DrawDialog()
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		termui.Body.Width = termui.TermWidth()
		termui.Body.Align()
		termui.Render(termui.Body)
	})
}

func (this *Ui) setInsertModeHandlers() {
	this.clearAllKbdHandlers()

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		fmt.Println(e.Data)
	})

	termui.Handle("/sys/kbd/m", func(e termui.Event) {
		fmt.Println(e.Data)
	})

	termui.Handle("/sys/kbd/i", func(e termui.Event) {
		fmt.Println(e.Data)
	})
	termui.Handle("/sys/kbd/", func(e termui.Event) {
		fmt.Println(e.Data)
		//here i can print all of the chars
	})

	termui.Handle("/sys/kbd/<escape>", func(e termui.Event) {
		this.setCommandModeHandlers()
	})
}
