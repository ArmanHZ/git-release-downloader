package ui

import (
	"github.com/ArmanHZ/git-release-downloader/utils"

	"github.com/rivo/tview"
)

type App struct {
	app *tview.Application

	mainFocus *FocusManager
	// TODO: impl l8r
	// repoSelectFocus *FocusManager
	downloadModalFocus *FocusManager

	activeFocus *FocusManager

	downloadList map[utils.Asset][]string // FIXME: I mean, we're only using the key value.

	mainGrid *tview.Grid

	urlInput       *tview.InputField
	downloadButton *tview.Button
	releaseView    *tview.TreeView
}

func New() *App {
	a := &App{
		app: tview.NewApplication(),
	}

	a.buildUI()
	a.bindEvents()
	a.initInputCapture()

	a.app.SetFocus(a.urlInput)

	return a
}

func (a *App) Run() error {
	return a.app.Run()
}
