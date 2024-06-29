package applications

import (
	"fmt"

	"github.com/rivo/tview"
	"jediproj.io/tview-tests/window"
)

type windowManager interface {
	Quit()
	AddAndShowModal(string, tview.Primitive, bool)
}
type applicationWindowManager struct {
	defaultPage string
	window      *window.Window
}

func (a *applicationWindowManager) Quit() {
	a.window.Pages.SwitchToPage(a.defaultPage)
}
func (a *applicationWindowManager) AddAndShowModal(name string, modal tview.Primitive, visible bool) {
	a.window.Pages.AddAndSwitchToPage(name, modal, visible)
}

type application struct {
	Name string
	windowManager
}

func (a *application) GetName() string {
	return a.Name
}

func (a *application) DoAction() {
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}
	// TODO: forms should be dynamic
	versions := []string{}
	for i := 0; i < 1000; i++ {
		versions = append(versions, fmt.Sprintf("v-%d", i))
	}
	form := tview.NewForm().AddDropDown("Version", versions, 0, nil)
	form.AddButton("Quit", a.Quit)
	form.SetBorder(true).SetTitle("Change version")
	a.AddAndShowModal(fmt.Sprintf("%s-modal", a.Name), modal(form, 40, 10), true)
}
func NewManager(defaultPage string) *applicationWindowManager {
	return &applicationWindowManager{
		defaultPage: defaultPage,
	}
}
func (a *applicationWindowManager) SetWindow(window *window.Window) *applicationWindowManager {
	if a.defaultPage == "" {
		panic("default page not initalized")
	}
	a.window = window
	return a
}
func NewApplication(name string) *application {
	return &application{
		Name: name,
	}
}
func (a *application) SetWindowManager(mng windowManager) *application {
	a.windowManager = mng
	return a

}
