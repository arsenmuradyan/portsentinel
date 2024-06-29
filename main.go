package main

import (
	"jediproj.io/tview-tests/applications"
	"jediproj.io/tview-tests/widgets"
	"jediproj.io/tview-tests/window"
)

const defaultPage = "applications"
const title = "Release Manager"

func main() {
	w := window.Window{}
	a := applications.NewManager(defaultPage).SetWindow(&w)
	w.InitWindow()
	list := widgets.NewButtonList()
	app1 := applications.NewApplication("visit").SetWindowManager(a)
	app2 := applications.NewApplication("test2").SetWindowManager(a)
	app3 := applications.NewApplication("test3").SetWindowManager(a)

	list.AddItem(app1)
	list.AddItem(app2)
	list.AddItem(app3)

	list.SetBorder(true).
		SetTitle(title)
	w.Pages.AddPage(defaultPage, list, true, true)
	w.App.Run()
}
