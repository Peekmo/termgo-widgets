package termgow

import (
	"github.com/peekmo/termbox-go"
)

type Spacing struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

type Element interface {
	draw()
	SetSize(width, height int)
	SetSpacing(s *Spacing)
}

type Layout interface {
	AddWidget(widget *Widget) error
	RemoveWidget(widget *Widget) error
}

type Widget interface {
	SetBackground(color termbox.Attribute)
	SetForeground(color termbox.Attribute)
	Hide() error
	Show() error
}
