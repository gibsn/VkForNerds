package vkclient

import (
	"fmt"

	"github.com/gizak/termui"
)

func (this *VkClient) showDialogs() {
	//TODO: should load updates in background for user not to see the lags
	this.Dialogs = this.Api.RequestDialogsHeaders()
	for i, _ := range this.Dialogs {
		this.Dialogs[i].FullName = this.Api.ResolveNameByUid(this.Dialogs[i].Uid)
	}

	messagesGrid := this.Ui.CreateDialogsGrid(this.Dialogs)
	this.Ui.ChangeCurrentGrid(messagesGrid)
}

func (this *VkClient) showCurrDialog() {
	dialogGrid := this.Ui.CreateDialogGrid(this.CurrOpenedDialog)
	this.Ui.ChangeCurrentGrid(dialogGrid)
}

func (this *VkClient) resizeCurrentGrid() {
	switch this.State {
	case DialogsHeaders:
		this.showDialogs()
	case Dialog:
		this.showCurrDialog()
	}
}

func (this *VkClient) setCommandModeHandlers() {
	termui.DefaultEvtStream.ResetHandlers()

	termui.Handle("/sys/kbd/q", func(e termui.Event) {
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/m", func(e termui.Event) {
		this.showDialogs()
	})

	termui.Handle("/sys/kbd/i", func(e termui.Event) {
		this.setInsertModeHandlers()
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		//TODO: should be i instead of 0 here
		this.CurrOpenedDialog = &this.Dialogs[0]
		this.showCurrDialog()
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		this.resizeCurrentGrid()
	})
}

func (this *VkClient) setInsertModeHandlers() {
	termui.DefaultEvtStream.ResetHandlers()

	//here i can print all of the chars
	termui.Handle("/sys/kbd/", func(e termui.Event) {
		fmt.Print(e.Data.(termui.EvtKbd).KeyStr)
	})

	termui.Handle("/sys/kbd/<escape>", func(e termui.Event) {
		this.setCommandModeHandlers()
	})

	termui.Handle("/sys/kbd/<enter>", func(e termui.Event) {
		//TODO: send message here
	})

	termui.Handle("/sys/wnd/resize", func(e termui.Event) {
		this.resizeCurrentGrid()
	})
}
