/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  @author Axel Anceau - 2014
  Package termgow allows to creates console apps faster

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
package termgow

import (
	"errors"
	"github.com/nsf/termbox-go"
	"github.com/peekmo/termgo-widgets/colors"
)

/**
 * Window is the base of a view with termgow
 */
type Window struct {
	id       int
	priority int

	Width   int
	Height  int
	Spacing *Spacing

	foreground termbox.Attribute
	background termbox.Attribute
}

var (
	windowId int = 1 // Current window id value
)

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  NewWindow returns a new Window with terminal's size

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func NewWindow() *Window {
	var win Window
	win.Width, win.Height = termbox.Size()
	win.Spacing = &Spacing{0, 0, 0, 0}

	win.id = windowId
	win.priority = -1

	win.foreground = colors.Default
	win.background = colors.Default

	windowId++
	return &win
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetSize changes windows's width & height value

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) SetSize(width, height int) {
	w.Width, w.Height = width, height
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetSpacing sets the spacing (in terms of chars) for every bounds
  of the window

  Note : If an element has a specified x/y pos, the spacing will not
  be considered

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) SetSpacing(s *Spacing) {
	w.Spacing = s
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetBackground changes background's color

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) SetBackground(color termbox.Attribute) {
	w.background = color
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetForeground sets foreground's color (characters)

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) SetForeground(color termbox.Attribute) {
	w.foreground = color
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Draw the window

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) draw() {
	for x := 0; x < w.Width; x++ {
		if x >= w.Spacing.Left && x < w.Width-w.Spacing.Right {
			for y := 0; y < w.Height; y++ {
				if y >= w.Spacing.Top && y < w.Height-w.Spacing.Bottom {
					termbox.SetCell(x, y, 'x', w.foreground, w.background)
				}
			}
		}
	}

	termbox.Flush()
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Hide the window from the program

  If the program has not been created, an error will be returned

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) Hide() error {
	if program == nil {
		return errors.New("The program does not exists")
	}

	err := program.hide(w)
	if err != nil {
		return err
	}

	w.priority = -1
	return nil
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  Show the window in the program.
  If the window is already showed, it will be displayed in foreground

  If the program has not been created or the window is not in its list,
  an error will be returned

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) Show() error {
	if program == nil {
		return errors.New("The program does not exists")
	}

	return program.show(w)
}
