package ui

import (
	"github.com/rivo/tview"
)

type FocusablePrimitives int

// TODO FIXME: The order of these are kept manually for far. Need to research a good
// way of autumating this part. Specially if we add more components.
const (
	UrlInput FocusablePrimitives = iota
	DownloadButton
	ReleaseView
)

const (
	ModalWDInput FocusablePrimitives = iota
	ModalCloseButton
	ModalDownloadButton
)

type FocusManager struct {
	primitives []tview.Primitive
	index      int
}

func NewFocusManager(p ...tview.Primitive) *FocusManager {
	return &FocusManager{
		primitives: p,
	}
}

func (f *FocusManager) AddFocusable(p tview.Primitive) {
	f.primitives = append(f.primitives, p)
}

func (f *FocusManager) Next(app *tview.Application) {
	if len(f.primitives) == 0 {
		return
	}

	f.index = (f.index + 1) % len(f.primitives)
	app.SetFocus(f.primitives[f.index])
}

func (f *FocusManager) Prev(app *tview.Application) {
	if len(f.primitives) == 0 {
		return
	}

	f.index--
	if f.index < 0 {
		f.index = len(f.primitives) - 1
	}

	app.SetFocus(f.primitives[f.index])
}
