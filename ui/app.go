package ui

import (
	// "github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app *tview.Application

	mainGrid *tview.Grid

	urlInput    *tview.InputField
	releaseView *tview.TreeView
}

func New() *App {
	a := &App{
		app: tview.NewApplication(),
	}

	a.buildUI()
	a.bindEvents()
	// a.initInputCapture()

	a.app.EnableMouse(true)

	return a
}

func (a *App) Run() error {
	return a.app.Run()
}
