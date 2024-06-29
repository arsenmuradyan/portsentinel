package window

import "github.com/rivo/tview"

type Window struct {
	App   *tview.Application
	Pages *tview.Pages
}

func (w *Window) InitWindow() {
	if w.App == nil {
		w.App = tview.NewApplication()
	}
	if w.Pages == nil {
		w.Pages = tview.NewPages()
	}
	w.App.SetRoot(w.Pages, true).EnableMouse(true)
}
