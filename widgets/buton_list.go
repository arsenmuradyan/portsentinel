package widgets

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ButtonListItem interface {
	DoAction()
	GetName() string
}

type ButtonList struct {
	*tview.Box
	items            []ButtonListItem
	currentItemIndex int
}

func (b *ButtonList) AddItem(item ButtonListItem) {
	b.items = append(b.items, item)
}

func NewButtonList() *ButtonList {
	return &ButtonList{
		Box: tview.NewBox(),
	}
}

func (r *ButtonList) Draw(screen tcell.Screen) {
	r.Box.DrawForSubclass(screen, r)
	x, y, width, height := r.GetInnerRect()
	for index, application := range r.items {
		if index >= height {
			break
		}
		button := ""
		if index == r.currentItemIndex {
			button = "[:blue]"
		}
		line := fmt.Sprintf(`%s%s%s`, button, application.GetName(), strings.Repeat(" ", width-len(application.GetName())))
		tview.Print(screen, line, x, y+index, width, tview.AlignLeft, tcell.ColorWhite)
	}
}

func (r *ButtonList) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return r.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyUp:
			r.currentItemIndex--
			if r.currentItemIndex < 0 {
				r.currentItemIndex = 0
			}
		case tcell.KeyDown:
			r.currentItemIndex++
			if r.currentItemIndex >= len(r.items) {
				r.currentItemIndex = len(r.items) - 1
			}
		case tcell.KeyEnter:
			r.items[r.currentItemIndex].DoAction()
		}
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (r *ButtonList) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return r.WrapMouseHandler(func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
		x, y := event.Position()
		_, rectY, _, _ := r.GetInnerRect()
		if !r.InRect(x, y) {
			return false, nil
		}

		if action == tview.MouseLeftClick {
			setFocus(r)
			index := y - rectY
			if index >= 0 && index < len(r.items) {
				r.currentItemIndex = index
				consumed = true
			}
		}

		return
	})
}
