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

	Width  int
	Height int

	foreground termbox.Attribute
	background termbox.Attribute
}

var (
	id int = 1 // Current window id value
)

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  NewWindow returns a new Window with terminal's size

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func NewWindow() *Window {
	var win Window
	win.Width, win.Height = termbox.Size()

	win.id = id

	win.foreground = colors.Default
	win.background = colors.Default

	id++
	return &win
}

/**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /*

  SetSize changes windows's width & height value

*/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/ /**/
func (w *Window) SetSize(width, height int) *Window {
	w.Width, w.Height = width, height

	return w
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
		for y := 0; y < w.Height; y++ {
			termbox.SetCell(x, y, 'X', w.foreground, w.background)
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

	program.hide(w)
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
